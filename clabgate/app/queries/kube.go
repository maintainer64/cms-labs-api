package queries

import (
	"context"
	"fmt"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/shared/connection"
	"gitlab.com/a10869/api-modules/shared/logs"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/dynamic"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesAdminQuery struct {
	clientset     *kubernetes.Clientset
	dynamicClient dynamic.Interface
	*KubernetesAdminConst
	*zerolog.Logger
}

type KubernetesAdminConst struct {
	DefaultNamespace string
	IssuerOIDC       string
	AdminBearerToken string
}

const (
	NamespaceListRole      = "namespace-lister"
	GitlabWebUrlDeploy     = "deploy-logs"
	NamespaceResourceQuota = "mem-cpu"
)

// NewKubernetesAdmin создает новый экземпляр администратора Kubernetes
func NewKubernetesAdmin(settings *connection.K8SConfig, zeroLogConf *logs.ZeroLoggerConf) (*KubernetesAdminQuery, error) {
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(settings.ConfigYaml))
	if err != nil {
		return nil, fmt.Errorf("failed to create config from KUBECONFIG: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes clientset: %v", err)
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes dynamicClient: %v", err)
	}
	return &KubernetesAdminQuery{
		clientset:     clientset,
		dynamicClient: dynamicClient,
		KubernetesAdminConst: &KubernetesAdminConst{
			DefaultNamespace: settings.Namespace,
			IssuerOIDC:       settings.IssuerOIDC,
			AdminBearerToken: config.BearerToken,
		},
		Logger: logs.NewZeroLogger(zeroLogConf),
	}, nil
}

func (k *KubernetesAdminQuery) NormalizeEntityName(entityName string) string {
	var result strings.Builder

	for _, r := range strings.ToLower(entityName) {
		switch {
		case r >= 'a' && r <= 'z':
			result.WriteRune(r)
		case r >= '0' && r <= '9':
			result.WriteRune(r)
		case r == '-':
			// Дефисы разрешены, но не в начале/конце и не подряд
			if result.Len() > 0 && result.String()[result.Len()-1] != '-' {
				result.WriteRune(r)
			}
		}
	}

	// Удаляем дефисы с конца
	normalized := strings.TrimRight(result.String(), "-")

	// Если после нормализации имя пустое, генерируем случайное
	if normalized == "" {
		normalized = fmt.Sprintf("user-%x", time.Now().UnixNano())
	}

	// Ограничиваем максимальную длину
	//(63 символа - ограничение Kubernetes)
	// (40 символов - чтобы ещё дописать можно было что-то)
	if len(normalized) > 63 {
		normalized = normalized[:63]
		normalized = strings.TrimRight(normalized, "-")
	}

	return normalized
}

// CreateUser создает ServiceAccount, если он еще не существует
func (k *KubernetesAdminQuery) CreateUser(ctx context.Context, username string) (*corev1.ServiceAccount, bool, error) {
	// Проверяем существование ServiceAccount
	sa, err := k.clientset.CoreV1().ServiceAccounts(k.KubernetesAdminConst.DefaultNamespace).Get(ctx, username, metav1.GetOptions{})
	if err == nil {
		k.Logger.Info().Msg(fmt.Sprintf("ServiceAccount %s already exists", username))
		return sa, false, nil
	}

	// Если не существует, создаем
	serviceAccount := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: username,
		},
	}

	sa, err = k.clientset.CoreV1().ServiceAccounts(k.KubernetesAdminConst.DefaultNamespace).Create(ctx, serviceAccount, metav1.CreateOptions{})
	if err != nil {
		return nil, false, fmt.Errorf("failed to create service account: %v", err)
	}
	k.Logger.Info().Msg(fmt.Sprintf("ServiceAccount %s created", username))
	return sa, true, nil
}

// CreateNamespace создает namespace, если он еще не существует
func (k *KubernetesAdminQuery) CreateNamespace(
	ctx context.Context,
	username string,
	namespace string,
) (*corev1.Namespace, bool, error) {
	// Проверяем существование namespace
	ns, err := k.clientset.CoreV1().Namespaces().Get(ctx, namespace, metav1.GetOptions{})
	if err == nil {
		k.Logger.Info().Msg(fmt.Sprintf("Namespace %s already exists", namespace))
		return ns, false, nil
	}

	// Если не существует, создаем
	newNs := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
			Labels: map[string]string{
				"owner": username,
			},
		},
	}

	createdNs, err := k.clientset.CoreV1().Namespaces().Create(ctx, newNs, metav1.CreateOptions{})
	if err != nil {
		return nil, false, fmt.Errorf("failed to create namespace: %v", err)
	}
	k.Logger.Info().Msg(fmt.Sprintf("Namespace %s created", namespace))
	// Накладываем органичения в ресурсах неймспейсу
	resourceQuota := &corev1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
			Name: NamespaceResourceQuota,
		},
		Spec: corev1.ResourceQuotaSpec{
			Hard: corev1.ResourceList{
				// 0.1 CPU
				"requests.cpu":    resource.MustParse("10m"),
				"requests.memory": resource.MustParse("10Mi"),
				// 0.05 CPU
				"limits.cpu":    resource.MustParse("50m"),
				"limits.memory": resource.MustParse("50Mi"),
			},
		},
	}

	_, err = k.clientset.CoreV1().ResourceQuotas(namespace).Create(ctx, resourceQuota, metav1.CreateOptions{})
	if err != nil {
		k.Logger.Error().Msg(fmt.Sprintf("Error creating ResourceQuota: %v", err))
	}
	k.Logger.Info().Msg(fmt.Sprintf("ResourceQuota %q created in namespace %q\n", resourceQuota.Name, namespace))
	return createdNs, true, nil
}

// GrantAccessNamespacesList - выдача доступов в список неймспейсов
func (k *KubernetesAdminQuery) GrantAccessNamespacesList(ctx context.Context, username []string) error {
	namespaceRole := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: NamespaceListRole,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"namespaces"},
				Verbs:     []string{"list", "get"},
			},
		},
	}

	// 1. Создаем или обновляем ClusterRole
	if _, err := k.clientset.RbacV1().ClusterRoles().Update(ctx, namespaceRole, metav1.UpdateOptions{}); err != nil {
		if errors.IsNotFound(err) {
			if _, err := k.clientset.RbacV1().ClusterRoles().Create(ctx, namespaceRole, metav1.CreateOptions{}); err != nil {
				k.Logger.Warn().Msg(fmt.Sprintf("ERROR: create namespace lister ClusterRole failed: %v", err))
			}
		} else {
			k.Logger.Warn().Msg(fmt.Sprintf("ERROR: update namespace lister ClusterRole failed: %v", err))
		}
	}

	for _, username := range username {
		// 2. Создаем ClusterRoleBinding для ролей
		binding := &rbacv1.ClusterRoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("%s-%s", username, NamespaceListRole),
			},
			Subjects: []rbacv1.Subject{
				{
					Kind:     "User",
					Name:     fmt.Sprintf("%s#%s", k.KubernetesAdminConst.IssuerOIDC, username),
					APIGroup: "rbac.authorization.k8s.io",
				},
			},
			RoleRef: rbacv1.RoleRef{
				Kind:     "ClusterRole",
				Name:     NamespaceListRole,
				APIGroup: "rbac.authorization.k8s.io",
			},
		}
		_, _ = k.clientset.RbacV1().ClusterRoleBindings().Create(ctx, binding, metav1.CreateOptions{})
	}
	return nil
}

// GetUserNamespacesOwner - получить все неймспейсы где пользователь является владельцем
func (k *KubernetesAdminQuery) GetUserNamespacesOwner(ctx context.Context, username string) ([]string, error) {
	// 1. Получаем список неймспейсов пользователя
	listOpts := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("owner=%s", username),
	}
	nsList, err := k.clientset.CoreV1().Namespaces().List(ctx, listOpts)
	if err != nil {
		return []string{}, fmt.Errorf("failed to list namespaces for %s: %v", username, err)
	}

	// Собираем имена неймспейсов пользователя
	var userNamespaces []string
	for _, ns := range nsList.Items {
		if strings.HasPrefix(ns.Name, username+"-") {
			userNamespaces = append(userNamespaces, ns.Name)
		}
	}
	return userNamespaces, nil
}

// GrantAccessUserNamespace выдача полных прав пользователям в неймспейс
func (k *KubernetesAdminQuery) GrantAccessUserNamespace(
	ctx context.Context,
	namespace string,
	usernames []string,
) error {
	// 1. Создаём роль полного доступа
	roleNameAccess := fmt.Sprintf("%s-full-access", namespace)
	role := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      roleNameAccess,
			Namespace: namespace,
		},
		Rules: []rbacv1.PolicyRule{{
			APIGroups: []string{"*"},
			Resources: []string{"*"},
			Verbs:     []string{"*"},
		}},
	}

	if _, err := k.clientset.RbacV1().Roles(namespace).Update(ctx, role, metav1.UpdateOptions{}); err != nil {
		if errors.IsNotFound(err) {
			if _, err := k.clientset.RbacV1().Roles(namespace).Create(ctx, role, metav1.CreateOptions{}); err != nil {
				k.Logger.Warn().Msg(fmt.Sprintf("ERROR: create Role %s in %s failed: %v", roleNameAccess, namespace, err))
			}
		} else {
			k.Logger.Warn().Msg(fmt.Sprintf("ERROR: update Role %s in %s failed: %v", roleNameAccess, namespace, err))
		}
	}

	for _, username := range usernames {
		rbName := fmt.Sprintf("%s-full-access-binding", username)
		rb := &rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      rbName,
				Namespace: namespace,
			},
			Subjects: []rbacv1.Subject{{
				Kind:      "User",
				Name:      fmt.Sprintf("%s#%s", k.KubernetesAdminConst.IssuerOIDC, username),
				Namespace: k.KubernetesAdminConst.DefaultNamespace,
			}},
			RoleRef: rbacv1.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "Role",
				Name:     roleNameAccess,
			},
		}
		k.Logger.Info().Msg(fmt.Sprintf("Granting access to username %s in role %s", username, rbName))
		_, _ = k.clientset.RbacV1().RoleBindings(namespace).Create(ctx, rb, metav1.CreateOptions{})
	}
	return nil
}

// GetNamespaceByName - получить определенный неймспейс
func (k *KubernetesAdminQuery) GetNamespaceByName(
	ctx context.Context,
	name string,
) (*corev1.Namespace, error) {
	// Проверяем существование namespace
	k.Logger.Info().Msg(fmt.Sprintf("Get namespace by name %s", name))
	ns, err := k.clientset.CoreV1().Namespaces().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		k.Logger.Info().Msg(fmt.Sprintf("Namespace %s is not defined", name))
		return ns, err
	}
	return ns, nil
}

// DeleteNamespaceByName - удалить определенный неймспейс
func (k *KubernetesAdminQuery) DeleteNamespaceByName(
	ctx context.Context,
	name string,
) error {
	// Проверяем существование namespace
	k.Logger.Info().Msg(fmt.Sprintf("Delete namespace by name %s", name))
	return k.clientset.CoreV1().Namespaces().Delete(ctx, name, metav1.DeleteOptions{})
}

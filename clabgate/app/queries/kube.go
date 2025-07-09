package queries

import (
	"context"
	"fmt"
	"gitlab.com/a10869/api-modules/shared/connection"
	"k8s.io/apimachinery/pkg/api/errors"
	"log"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesAdminQuery struct {
	clientset        *kubernetes.Clientset
	defaultNamespace string
}

// NewKubernetesAdmin создает новый экземпляр администратора Kubernetes
func NewKubernetesAdmin(settings *connection.K8SConfig) (*KubernetesAdminQuery, error) {
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(settings.ConfigYaml))
	if err != nil {
		return nil, fmt.Errorf("failed to create config from KUBECONFIG: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes clientset: %v", err)
	}
	return &KubernetesAdminQuery{clientset: clientset, defaultNamespace: settings.Namespace}, nil
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
func (k *KubernetesAdminQuery) CreateUser(ctx context.Context, username string) (*corev1.ServiceAccount, error) {
	// Проверяем существование ServiceAccount
	sa, err := k.clientset.CoreV1().ServiceAccounts(k.defaultNamespace).Get(ctx, username, metav1.GetOptions{})
	if err == nil {
		log.Printf("ServiceAccount %s already exists", username)
		return sa, nil
	}

	// Если не существует, создаем
	serviceAccount := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: username,
		},
	}

	sa, err = k.clientset.CoreV1().ServiceAccounts(k.defaultNamespace).Create(ctx, serviceAccount, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create service account: %v", err)
	}

	log.Printf("ServiceAccount %s created", username)
	return sa, nil
}

// CreateNamespace создает namespace, если он еще не существует
func (k *KubernetesAdminQuery) CreateNamespace(ctx context.Context, username string, namespace string) (*corev1.Namespace, error) {
	// Проверяем существование namespace
	ns, err := k.clientset.CoreV1().Namespaces().Get(ctx, namespace, metav1.GetOptions{})
	if err == nil {
		log.Printf("Namespace %s already exists", namespace)
		return ns, nil
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
		return nil, fmt.Errorf("failed to create namespace: %v", err)
	}

	log.Printf("Namespace %s created", namespace)
	return createdNs, nil
}

// GrantAccess предоставляет пользователю полные права только в своих неймспейсах (username-*)
func (k *KubernetesAdminQuery) GrantAccess(ctx context.Context, username string) error {
	user := k.NormalizeEntityName(username)

	// 1. Получаем список неймспейсов пользователя
	listOpts := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("owner=%s", user),
	}
	nsList, err := k.clientset.CoreV1().Namespaces().List(ctx, listOpts)
	if err != nil {
		return fmt.Errorf("failed to list namespaces for %s: %v", user, err)
	}

	// Собираем имена неймспейсов пользователя
	var userNamespaces []string
	for _, ns := range nsList.Items {
		if strings.HasPrefix(ns.Name, user+"-") {
			userNamespaces = append(userNamespaces, ns.Name)
		}
	}

	// 2. Создаем ClusterRole для просмотра списка неймспейсов (требуется для Dashboard)
	viewRole := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-namespace-lister", user),
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"namespaces"},
				Verbs:     []string{"list"},
			},
		},
	}

	// Создаем или обновляем ClusterRole
	if _, err := k.clientset.RbacV1().ClusterRoles().Update(ctx, viewRole, metav1.UpdateOptions{}); err != nil {
		if errors.IsNotFound(err) {
			if _, err := k.clientset.RbacV1().ClusterRoles().Create(ctx, viewRole, metav1.CreateOptions{}); err != nil {
				log.Printf("ERROR: create namespace lister ClusterRole failed: %v", err)
			}
		} else {
			log.Printf("ERROR: update namespace lister ClusterRole failed: %v", err)
		}
	}

	// 3. Создаем ClusterRole для просмотра только своих неймспейсов
	restrictedViewRole := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-namespace-viewer", user),
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups:     []string{""},
				Resources:     []string{"namespaces"},
				Verbs:         []string{"get"},
				ResourceNames: userNamespaces,
			},
		},
	}

	if _, err := k.clientset.RbacV1().ClusterRoles().Update(ctx, restrictedViewRole, metav1.UpdateOptions{}); err != nil {
		if errors.IsNotFound(err) {
			if _, err := k.clientset.RbacV1().ClusterRoles().Create(ctx, restrictedViewRole, metav1.CreateOptions{}); err != nil {
				log.Printf("ERROR: create restricted namespace viewer ClusterRole failed: %v", err)
			}
		} else {
			log.Printf("ERROR: update restricted namespace viewer ClusterRole failed: %v", err)
		}
	}

	// 4. Создаем ClusterRoleBinding для обоих ролей
	binding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-namespace-access", user),
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      user,
				Namespace: k.defaultNamespace,
			},
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "ClusterRole",
			Name:     fmt.Sprintf("%s-namespace-lister", user),
			APIGroup: "rbac.authorization.k8s.io",
		},
	}

	// Удаляем старый binding если существует
	_ = k.clientset.RbacV1().ClusterRoleBindings().Delete(ctx, binding.Name, metav1.DeleteOptions{})
	if _, err := k.clientset.RbacV1().ClusterRoleBindings().Create(ctx, binding, metav1.CreateOptions{}); err != nil {
		log.Printf("ERROR: create namespace access ClusterRoleBinding failed: %v", err)
	}

	// 5. Даем полные права в каждом неймспейсе пользователя
	roleName := fmt.Sprintf("%s-full-access", user)
	rbName := fmt.Sprintf("%s-full-access-binding", user)

	for _, ns := range nsList.Items {
		if !strings.HasPrefix(ns.Name, user+"-") {
			continue
		}

		role := &rbacv1.Role{
			ObjectMeta: metav1.ObjectMeta{
				Name:      roleName,
				Namespace: ns.Name,
			},
			Rules: []rbacv1.PolicyRule{{
				APIGroups: []string{"*"},
				Resources: []string{"*"},
				Verbs:     []string{"*"},
			}},
		}

		if _, err := k.clientset.RbacV1().Roles(ns.Name).Update(ctx, role, metav1.UpdateOptions{}); err != nil {
			if errors.IsNotFound(err) {
				if _, err := k.clientset.RbacV1().Roles(ns.Name).Create(ctx, role, metav1.CreateOptions{}); err != nil {
					log.Printf("ERROR: create Role %s in %s failed: %v", roleName, ns.Name, err)
				}
			} else {
				log.Printf("ERROR: update Role %s in %s failed: %v", roleName, ns.Name, err)
			}
		}

		rb := &rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      rbName,
				Namespace: ns.Name,
			},
			Subjects: []rbacv1.Subject{{
				Kind:      "ServiceAccount",
				Name:      user,
				Namespace: k.defaultNamespace,
			}},
			RoleRef: rbacv1.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "Role",
				Name:     roleName,
			},
		}

		_ = k.clientset.RbacV1().RoleBindings(ns.Name).Delete(ctx, rbName, metav1.DeleteOptions{})
		if _, err := k.clientset.RbacV1().RoleBindings(ns.Name).Create(ctx, rb, metav1.CreateOptions{}); err != nil {
			log.Printf("ERROR: create RoleBinding %s in %s failed: %v", rbName, ns.Name, err)
		}
	}

	return nil
}

// GetUserToken возвращает токен для ServiceAccount (улучшенная версия)
func (k *KubernetesAdminQuery) GetUserToken(ctx context.Context, username string) (string, error) {

	// 1. Проверяем существующий секрет
	secretName := fmt.Sprintf("%s-token", username)

	// Попробуем получить существующий секрет
	if secret, err := k.clientset.CoreV1().Secrets(k.defaultNamespace).Get(ctx, secretName, metav1.GetOptions{}); err == nil {
		if token, exists := secret.Data["token"]; exists {
			return string(token), nil
		}
	}

	// 2. Если секрета нет, создаем новый
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: k.defaultNamespace,
			Annotations: map[string]string{
				"kubernetes.io/service-account.name": username,
			},
		},
		Type: corev1.SecretTypeServiceAccountToken,
	}

	createdSecret, err := k.clientset.CoreV1().Secrets(k.defaultNamespace).Create(ctx, secret, metav1.CreateOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to create secret: %v", err)
	}

	// 3. Ждем генерации токена (с таймаутом)
	ctxTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	for {
		select {
		case <-ctxTimeout.Done():
			return "", fmt.Errorf("timeout waiting for token generation")
		default:
			updatedSecret, err := k.clientset.CoreV1().Secrets(k.defaultNamespace).Get(ctx, createdSecret.Name, metav1.GetOptions{})
			if err != nil {
				return "", fmt.Errorf("failed to get secret: %v", err)
			}

			if token, exists := updatedSecret.Data["token"]; exists && len(token) > 0 {
				return string(token), nil
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	admin, err := NewKubernetesAdmin(nil)
	if err != nil {
		log.Fatalf("Failed to create Kubernetes admin: %v", err)
	}

	ctx := context.Background()
	username := admin.NormalizeEntityName("kodolov-s")
	namespace := admin.NormalizeEntityName("kodolov-s-task-one")

	// 1. Создаем пользователя (ServiceAccount)
	_, err = admin.CreateUser(ctx, username)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	// 2. Создаем namespace
	_, err = admin.CreateNamespace(ctx, username, namespace)
	if err != nil {
		log.Fatalf("Failed to create namespace: %v", err)
	}

	// 3. Выдаем права на exec в namespace
	err = admin.GrantAccess(ctx, username)
	if err != nil {
		log.Fatalf("Failed to grant exec access: %v", err)
	}

	// 4. Получаем токен для пользователя
	token, err := admin.GetUserToken(ctx, username)
	if err != nil {
		log.Fatalf("Failed to get user token: %v", err)
	}
	log.Printf("Token for user %s: %s", username, token)
}

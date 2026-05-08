package queries

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type DeploymentInfo struct {
	Name          string `json:"name"`
	Status        string `json:"status"`
	Restarts      int32  `json:"restarts"`
	ReadyReplicas int32  `json:"ready_replicas"`
	Replicas      *int32 `json:"replicas"`
}

type ServiceInfo struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	ExternalIP []string `json:"external_ip,omitempty"`
	ClusterIP  string   `json:"cluster_ip"`
}

type KubernetesAdminQuery struct {
	clientset     *kubernetes.Clientset
	dynamicClient dynamic.Interface
	*zerolog.Logger
}

// NewKubernetesAdmin создает клиент Kubernetes из in-cluster serviceaccount
func NewKubernetesAdmin(zeroLogConf *logs.ZeroLoggerConf) (*KubernetesAdminQuery, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create in-cluster config: %v", err)
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
		Logger:        logs.NewZeroLogger(zeroLogConf),
	}, nil
}

// GetTopologyYAML возвращает YAML-представление конкретной топологии (CRD)
func (k *KubernetesAdminQuery) GetTopologyYAML(ctx context.Context, namespace string) ([]byte, error) {
	if namespace == "" {
		return []byte{}, fmt.Errorf("namespace is required")
	}

	gvr := schema.GroupVersionResource{
		Group:    "clabernetes.containerlab.dev",
		Version:  "v1alpha1",
		Resource: "topologies",
	}

	list, err := k.dynamicClient.Resource(gvr).Namespace(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return []byte{}, fmt.Errorf("failed to list Topologies in namespace %s: %v", namespace, err)
	}

	if len(list.Items) == 0 {
		return []byte{}, fmt.Errorf("no Topology found in namespace %s", namespace)
	}

	for _, unstructuredObj := range list.Items {
		yamlData, err := yaml.Marshal(unstructuredObj.Object)
		if err == nil {
			return yamlData, nil
		}
	}
	return []byte{}, fmt.Errorf("failed to marshal Topology to YAML")
}

// GetDeploymentsInfo возвращает информацию о deployments в namespace
func (k *KubernetesAdminQuery) GetDeploymentsInfo(ctx context.Context, namespace string) ([]DeploymentInfo, error) {
	deployments, err := k.clientset.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list deployments in namespace %s: %v", namespace, err)
	}

	var result []DeploymentInfo
	for _, deploy := range deployments.Items {
		status := "Ready"
		if deploy.Status.ReadyReplicas != deploy.Status.Replicas {
			status = "NotReady"
		}

		var restarts int32 = 0
		pods, err := k.clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
			LabelSelector: metav1.FormatLabelSelector(deploy.Spec.Selector),
		})
		if err == nil {
			for _, pod := range pods.Items {
				for _, cs := range pod.Status.ContainerStatuses {
					restarts += cs.RestartCount
				}
			}
		}

		result = append(result, DeploymentInfo{
			Name:          deploy.Name,
			Status:        status,
			Restarts:      restarts,
			ReadyReplicas: deploy.Status.ReadyReplicas,
			Replicas:      deploy.Spec.Replicas,
		})
	}
	return result, nil
}

// GetServicesInfo возвращает информацию о services типа LoadBalancer в namespace
func (k *KubernetesAdminQuery) GetServicesInfo(ctx context.Context, namespace string) ([]ServiceInfo, error) {
	services, err := k.clientset.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list services in namespace %s: %v", namespace, err)
	}

	var result []ServiceInfo
	for _, svc := range services.Items {
		if svc.Spec.Type != "LoadBalancer" {
			continue
		}

		externalIP := make([]string, 0)
		if svc.Status.LoadBalancer.Ingress != nil && len(svc.Status.LoadBalancer.Ingress) > 0 {
			for _, ingress := range svc.Status.LoadBalancer.Ingress {
				if ingress.IP != "" {
					externalIP = append(externalIP, ingress.IP)
				}
			}
		}

		result = append(result, ServiceInfo{
			Name:       svc.Name,
			Type:       string(svc.Spec.Type),
			ExternalIP: externalIP,
			ClusterIP:  svc.Spec.ClusterIP,
		})
	}
	return result, nil
}

package queries

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KubernetesContainerInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Pod  string `json:"pod"`
	// Status enum: running,waiting,terminated,unknown
	Status       string `json:"status"`
	Namespace    string `json:"namespace"`
	Ready        bool   `json:"ready"`
	RestartCount int32  `json:"restartCount"`
}

func getContainerState(state corev1.ContainerState) string {
	switch {
	case state.Running != nil:
		return "running"
	case state.Waiting != nil:
		return "waiting"
	case state.Terminated != nil:
		return "terminated"
	default:
		return "unknown"
	}
}

// GetPodByDeploymentName - получить определенный неймспейс
func (k *KubernetesAdminQuery) GetPodByDeploymentName(
	ctx context.Context,
	namespace string,
	deploymentName string,
) ([]KubernetesContainerInfo, error) {
	var containers []KubernetesContainerInfo
	k.Logger.Info().Msg(fmt.Sprintf("Get pod by deployment: %s into namespace: %s", deploymentName, namespace))
	pods, err := k.clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app.kubernetes.io/name=%s", deploymentName),
	})
	if err != nil {
		k.Logger.Info().Msg(fmt.Sprintf("Error listing pods: %v", err))
		return nil, err
	}
	if len(pods.Items) == 0 {
		k.Logger.Info().Msg(fmt.Sprintf("No pods found for deployment %s in namespace %s", deploymentName, namespace))
		return nil, fmt.Errorf("no pods found for deployment %s in namespace %s", deploymentName, namespace)
	}
	pod := pods.Items[0]

	// Get status for all containers in the pod
	statusMap := make(map[string]corev1.ContainerStatus)
	for _, status := range pod.Status.ContainerStatuses {
		statusMap[status.Name] = status
	}

	// Process regular containers
	for i, container := range pod.Spec.Containers {
		status, exists := statusMap[container.Name]
		if !exists {
			continue
		}

		containers = append(containers, KubernetesContainerInfo{
			ID:           fmt.Sprintf("%s-%d", pod.UID, i),
			Name:         container.Name,
			Pod:          pod.Name,
			Namespace:    namespace,
			Status:       getContainerState(status.State),
			Ready:        status.Ready,
			RestartCount: status.RestartCount,
		})
	}
	return containers, nil
}

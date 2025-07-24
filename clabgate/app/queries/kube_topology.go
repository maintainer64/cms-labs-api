package queries

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	yaml "sigs.k8s.io/yaml/goyaml.v3"
)

// GetTopologyYAML возвращает YAML-представление конкретной топологии (CRD)
func (k *KubernetesAdminQuery) GetTopologyYAML(ctx context.Context, namespace string) ([]byte, error) {
	if namespace == "" {
		return []byte{}, fmt.Errorf("namespace is required")
	}

	// GroupVersionResource (GVR) на основе анализа YAML
	gvr := schema.GroupVersionResource{
		Group:    "clabernetes.containerlab.dev", // Из apiVersion
		Version:  "v1alpha1",                     // Из apiVersion
		Resource: "topologies",                   // Plural от kind "Topology"
	}

	// Получаем список объектов Topology в указанном namespace
	list, err := k.dynamicClient.Resource(gvr).Namespace(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return []byte{}, fmt.Errorf("failed to list Topologies in namespace %s: %v", namespace, err)
	}

	// Проверяем, есть ли элементы в списке
	if len(list.Items) == 0 {
		return []byte{}, fmt.Errorf("no Topology found in namespace %s", namespace)
	}

	for _, unstructuredObj := range list.Items {
		// Конвертируем Unstructured в YAML
		yamlData, err := yaml.Marshal(unstructuredObj.Object)
		if err == nil {
			return yamlData, nil
		}
	}
	return []byte{}, fmt.Errorf("failed to marshal Topology to YAML: %v", err)
}

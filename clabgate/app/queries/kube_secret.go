package queries

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetSecretByName получить секрет по названию
func (k *KubernetesAdminQuery) GetSecretByName(ctx context.Context, namespace string, key string) (string, error) {
	if namespace == "" {
		return "", fmt.Errorf("namespace is required")
	}

	secret, err := k.clientset.CoreV1().Secrets(namespace).Get(ctx, key, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get secret: %v", err)
	}

	if value, exists := secret.Data["value"]; exists {
		return string(value), nil
	}

	return "", fmt.Errorf("secret '%s' in namespace '%s' doesn't contain 'value' field", key, namespace)
}

// SetSecretByName установить секрет по названию
func (k *KubernetesAdminQuery) SetSecretByName(ctx context.Context, namespace string, key string, value string) error {
	if namespace == "" {
		return fmt.Errorf("namespace is required")
	}

	secretInterface := k.clientset.CoreV1().Secrets(namespace)

	// Пробуем получить существующий секрет
	secret, err := secretInterface.Get(ctx, key, metav1.GetOptions{})
	if err != nil {
		// Если секрет не существует, создаем новый
		if errors.IsNotFound(err) {
			newSecret := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name: key,
				},
				Data: map[string][]byte{
					"value": []byte(value),
				},
			}

			_, err = secretInterface.Create(ctx, newSecret, metav1.CreateOptions{})
			return err
		}
		return fmt.Errorf("failed to check existing secret: %v", err)
	}

	// Обновляем существующий секрет
	secret.Data["value"] = []byte(value)
	_, err = secretInterface.Update(ctx, secret, metav1.UpdateOptions{})
	return err
}

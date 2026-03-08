package vault

import (
	"context"
)

// ClientInterface описывает публичные методы для работы с Vault.
type ClientInterface interface {
	// CreateServiceExtension сохраняет данные расширения для сервиса.
	CreateServiceExtension(ctx context.Context, serviceName, addonName string, data map[string]any) error

	// CreateUsernameExtension сохраняет данные расширения для пользователя.
	CreateUsernameExtension(ctx context.Context, username, addonName string, data map[string]any) error

	// IssueServiceToken выпускает сервисный токен с заданными опциями.
	IssueServiceToken(ctx context.Context, targetName string, opts ServiceTokenOptions) (*ServiceTokenInfo, error)

	// RevokeServiceToken отзывает сервисный токен.
	RevokeServiceToken(ctx context.Context, targetName string) error

	// RotateServiceToken перевыпускает сервисный токен (отзыв + выпуск нового).
	RotateServiceToken(ctx context.Context, targetName string, opts ServiceTokenOptions) (*ServiceTokenInfo, error)

	GetVaultAddr() string

	CreateKubernetesRole(ctx context.Context, targetName string) error
	RevokeKubernetesRole(ctx context.Context, targetName string) error

	UserBindAccessServices(ctx context.Context, binds []UsersAndServices) error
}

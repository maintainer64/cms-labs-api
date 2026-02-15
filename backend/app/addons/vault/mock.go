package vault

import "context"

// MockVaultClient — простая заглушка для VaultClientInterface.
// Позволяет задавать функции-обработчики для каждого метода.
type MockVaultClient struct {
	CreateServiceExtensionFunc  func(ctx context.Context, serviceName, addonName string, data map[string]any) error
	CreateUsernameExtensionFunc func(ctx context.Context, username, addonName string, data map[string]any) error
	IssueServiceTokenFunc       func(ctx context.Context, targetName string, opts ServiceTokenOptions) (*ServiceTokenInfo, error)
	RevokeServiceTokenFunc      func(ctx context.Context, targetName string) error
	RotateServiceTokenFunc      func(ctx context.Context, targetName string, opts ServiceTokenOptions) (*ServiceTokenInfo, error)
	GetVaultAddrFunc            func() string
}

// CreateServiceExtension вызывает заданную функцию-заглушку.
func (m *MockVaultClient) CreateServiceExtension(ctx context.Context, serviceName, addonName string, data map[string]any) error {
	if m.CreateServiceExtensionFunc != nil {
		return m.CreateServiceExtensionFunc(ctx, serviceName, addonName, data)
	}
	return nil
}

// CreateUsernameExtension вызывает заданную функцию-заглушку.
func (m *MockVaultClient) CreateUsernameExtension(ctx context.Context, username, addonName string, data map[string]any) error {
	if m.CreateUsernameExtensionFunc != nil {
		return m.CreateUsernameExtensionFunc(ctx, username, addonName, data)
	}
	return nil
}

// IssueServiceToken вызывает заданную функцию-заглушку.
func (m *MockVaultClient) IssueServiceToken(ctx context.Context, targetName string, opts ServiceTokenOptions) (*ServiceTokenInfo, error) {
	if m.IssueServiceTokenFunc != nil {
		return m.IssueServiceTokenFunc(ctx, targetName, opts)
	}
	return nil, nil
}

// RevokeServiceToken вызывает заданную функцию-заглушку.
func (m *MockVaultClient) RevokeServiceToken(ctx context.Context, targetName string) error {
	if m.RevokeServiceTokenFunc != nil {
		return m.RevokeServiceTokenFunc(ctx, targetName)
	}
	return nil
}

// RotateServiceToken вызывает заданную функцию-заглушку.
func (m *MockVaultClient) RotateServiceToken(ctx context.Context, targetName string, opts ServiceTokenOptions) (*ServiceTokenInfo, error) {
	if m.RotateServiceTokenFunc != nil {
		return m.RotateServiceTokenFunc(ctx, targetName, opts)
	}
	return nil, nil
}

func (m *MockVaultClient) GetVaultAddr() string {
	if m.GetVaultAddrFunc != nil {
		return m.GetVaultAddrFunc()
	}
	return ""
}

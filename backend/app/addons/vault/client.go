package vault

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/vault/api"
	"gitlab.com/a10869/api-modules/shared/connection"
)

func NewClient(cfg *connection.Vault) *Client {
	apiConfig := api.DefaultConfig()
	apiConfig.Address = cfg.VaultAddr
	c, _ := api.NewClient(apiConfig)
	c.SetToken(cfg.VaultToken)
	return &Client{
		Client:    c,
		VaultAddr: cfg.VaultAddr,
	}
}

type Client struct {
	Client    *api.Client
	VaultAddr string
}

const (
	Read  = "read"
	Write = "write"
	None  = "none"
)

type ServiceTokenOptions struct {
	TTL        string
	MaxTTL     string
	Renewable  *bool
	Policies   []string
	Metadata   map[string]string
	NumUses    int
	Permission string
}

type ServiceTokenInfo struct {
	Token       string
	Accessor    string
	ServiceName string
	Policies    []string
	TTL         int
	Renewable   bool
	CreatedAt   string
	ExpiresAt   string
}

func (v *Client) CreateServiceExtension(ctx context.Context, serviceName, addonName string, data map[string]any) error {
	p := fmt.Sprintf("secret/data/services/%s/.infra/%s", sanitize(serviceName), sanitize(addonName))

	payload := map[string]any{
		"data": map[string]any{
			"_metadata": map[string]string{
				"created": time.Now().Format(time.RFC3339),
				"updated": time.Now().Format(time.RFC3339),
				"service": serviceName,
				"addon":   addonName,
			},
		},
	}
	for k, val := range data {
		payload["data"].(map[string]any)[k] = val
	}

	_, err := v.Client.Logical().WriteWithContext(ctx, p, payload)
	return err
}

func (v *Client) CreateUsernameExtension(ctx context.Context, username, addonName string, data map[string]any) error {
	p := fmt.Sprintf("secret/data/users/%s/.infra/%s", sanitize(username), sanitize(addonName))

	payload := map[string]any{
		"data": map[string]any{
			"_metadata": map[string]string{
				"created":  time.Now().Format(time.RFC3339),
				"updated":  time.Now().Format(time.RFC3339),
				"username": username,
				"addon":    addonName,
			},
		},
	}
	for k, val := range data {
		payload["data"].(map[string]any)[k] = val
	}

	_, err := v.Client.Logical().WriteWithContext(ctx, p, payload)
	return err
}

// IssueServiceToken / RotateServiceToken / RevokeServiceToken — полностью аналог TS-версии
func (v *Client) IssueServiceToken(ctx context.Context, targetName string, opts ServiceTokenOptions) (*ServiceTokenInfo, error) {
	serviceName := "svc-" + sanitize(targetName)
	servicePath := "services/" + sanitize(targetName)

	perm := opts.Permission
	if perm == "" {
		perm = Write
	}
	policyName, err := v.createOrUpdatePolicy(ctx, servicePath, perm)
	if err != nil {
		return nil, err
	}

	policies := append(opts.Policies, policyName)

	ttl := opts.TTL
	if ttl == "" {
		ttl = "720h" // 30 дней по умолчанию
	}

	req := &api.TokenCreateRequest{
		DisplayName:    serviceName,
		Policies:       policies,
		TTL:            ttl,
		ExplicitMaxTTL: opts.MaxTTL,
		Renewable:      opts.Renewable,
		NumUses:        opts.NumUses,
		Metadata:       opts.Metadata,
		NoParent:       true,
	}

	secret, err := v.Client.Auth().Token().CreateWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	auth := secret.Auth

	info := &ServiceTokenInfo{
		Token:       auth.ClientToken,
		Accessor:    auth.Accessor,
		ServiceName: serviceName,
		Policies:    auth.Policies,
		TTL:         auth.LeaseDuration,
		Renewable:   auth.Renewable,
		CreatedAt:   time.Now().Format(time.RFC3339),
		ExpiresAt:   time.Now().Add(time.Duration(auth.LeaseDuration) * time.Second).Format(time.RFC3339),
	}

	// сохраняем accessor для revoke
	_ = v.CreateServiceExtension(ctx, ".service-tokens", serviceName, map[string]any{
		"accessor":    auth.Accessor,
		"serviceName": serviceName,
		"targetName":  targetName,
	})

	return info, nil
}

func (v *Client) RevokeServiceToken(ctx context.Context, targetName string) error {
	serviceName := "svc-" + sanitize(targetName)

	// читаем accessor
	secret, err := v.Client.Logical().ReadWithContext(ctx, "secret/data/.service-tokens/"+serviceName)
	if err != nil || secret == nil {
		return err
	}
	accessor, _ := secret.Data["data"].(map[string]any)["accessor"].(string)
	if accessor != "" {
		_ = v.Client.Sys().RevokeWithContext(ctx, accessor)
	}

	// удаляем политику и запись
	_ = v.Client.Sys().DeletePolicyWithContext(ctx, strings.ReplaceAll(sanitize("services/"+targetName), "/", "-")+"-write")
	_, _ = v.Client.Logical().DeleteWithContext(ctx, "secret/metadata/.service-tokens/"+serviceName)
	return nil
}

func (v *Client) RotateServiceToken(ctx context.Context, targetName string, opts ServiceTokenOptions) (*ServiceTokenInfo, error) {
	_ = v.RevokeServiceToken(ctx, targetName) // best-effort
	return v.IssueServiceToken(ctx, targetName, opts)
}

// ---------------------------------------------------------------------
// внутренние хелперы
// ---------------------------------------------------------------------

func (v *Client) createOrUpdatePolicy(ctx context.Context, servicePath string, perm string) (string, error) {
	name := strings.ReplaceAll(sanitize(servicePath), "/", "-") + "-" + perm
	hcl := generatePolicyHCL(servicePath, perm)
	return name, v.Client.Sys().PutPolicyWithContext(ctx, name, hcl)
}

func generatePolicyHCL(servicePath string, perm string) string {
	p := sanitize(servicePath)
	caps := ""
	switch perm {
	case Read:
		caps = `capabilities = ["read","list"]`
	case Write:
		caps = `capabilities = ["create","read","update","delete","list"]`
	case None:
		caps = `capabilities = []`
	}
	return fmt.Sprintf(`
path "secret/data/%s/*" { %s }
path "secret/metadata/%s/*" { %s }
`, p, caps, p, caps)
}

var sanitizeRe = regexp.MustCompile(`[^a-zA-Z0-9\-_\./]`)

func sanitize(s string) string {
	s = sanitizeRe.ReplaceAllString(s, "-")
	s = strings.ReplaceAll(s, "..", "")
	return strings.Trim(s, "/")
}

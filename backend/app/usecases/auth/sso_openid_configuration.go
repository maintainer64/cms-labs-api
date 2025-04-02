package auth

import (
	"net/url"
)

type SSOOpenidConfigurationInputDTO struct {
	BaseURL string `json:"base_url"`
}

type SSOOpenidConfigurationOutputDTO struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	UserInfoEndpoint                  string   `json:"userinfo_endpoint"`
	JwksUri                           string   `json:"jwks_uri"`
	IntrospectionEndpoint             string   `json:"introspection_endpoint"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IdTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	ScopesSupported                   []string `json:"scopes_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	CodeChallengeMethodsSupported     []string `json:"code_challenge_methods_supported"`
	GrantTypesSupported               []string `json:"grant_types_supported"`
}

type SSOOpenidConfiguration struct {
}

func (u *SSOOpenidConfiguration) Execute(inputDTO SSOOpenidConfigurationInputDTO) (
	*SSOOpenidConfigurationOutputDTO, error,
) {
	// Базовый URL вашего сервиса
	issuer, err := url.JoinPath(inputDTO.BaseURL, "/api/v1/sso")
	if err != nil {
		return nil, err
	}
	// Формируем ответ согласно спецификации OpenID Connect Discovery
	return &SSOOpenidConfigurationOutputDTO{
		Issuer:                            inputDTO.BaseURL,
		AuthorizationEndpoint:             issuer + "/authorize",
		TokenEndpoint:                     issuer + "/token",
		UserInfoEndpoint:                  issuer + "/userinfo",
		JwksUri:                           issuer + "/jwks",
		IntrospectionEndpoint:             issuer + "/introspect",
		ResponseTypesSupported:            []string{"code"},
		SubjectTypesSupported:             []string{"public"},
		IdTokenSigningAlgValuesSupported:  []string{"RS256"},
		ScopesSupported:                   []string{"default", "openid", "email", "profile"},
		TokenEndpointAuthMethodsSupported: []string{"client_secret_post", "client_secret_basic"},
		CodeChallengeMethodsSupported:     []string{"PS256"},
		GrantTypesSupported:               []string{"authorization_code", "implicit"},
	}, nil
}

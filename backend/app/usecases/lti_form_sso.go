package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

type AuthProviderListSSOUC struct {
	AuthProviderQueries *lti_query.AuthProviderQueries
}

type AuthProviderListSSORequest struct {
	JSONRPC string       `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string       `json:"method" default:"lti_form.sso_list" required:"true"`
	Params  *interface{} `json:"params,omitempty"`
	ID      string       `json:"id,omitempty" default:"1" required:"true"`
}

type AuthProviderListSSOOutputDTO struct {
	Model []models.AuthProviderListItem `json:"model" validate:"required"`
}

type AuthProviderListSSOResponse struct {
	JSONRPC string                       `json:"jsonrpc" default:"2.0" required:"true"`
	Result  AuthProviderListSSOOutputDTO `json:"result,omitempty"`
	Error   interface{}                  `json:"error,omitempty"`
	ID      string                       `json:"id,omitempty" default:"1" required:"true"`
}

func (u *AuthProviderListSSOUC) Execute() (AuthProviderListOutputDTO, error) {
	entities, err := u.AuthProviderQueries.SSOURLList()
	return AuthProviderListOutputDTO{
		Model: entities,
	}, err
}

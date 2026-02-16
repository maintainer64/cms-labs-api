package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

type AuthProviderListUC struct {
	AuthProviderQueries *lti_query.AuthProviderQueries
}

type AuthProviderListInputDTO struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type AuthProviderListRequest struct {
	JSONRPC string                   `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string                   `json:"method" default:"lti_form.list" required:"true"`
	Params  AuthProviderListInputDTO `json:"params,omitempty"`
	ID      string                   `json:"id,omitempty" default:"1" required:"true"`
}

type AuthProviderListOutputDTO struct {
	Model      []models.AuthProviderListItem `json:"model" validate:"required"`
	TotalCount int64                         `json:"total_count" validate:"required"`
}

type AuthProviderListResponse struct {
	JSONRPC string                    `json:"jsonrpc" default:"2.0" required:"true"`
	Result  AuthProviderListOutputDTO `json:"result,omitempty"`
	Error   interface{}               `json:"error,omitempty"`
	ID      string                    `json:"id,omitempty" default:"1" required:"true"`
}

func (u *AuthProviderListUC) Execute(dto AuthProviderListInputDTO) (AuthProviderListOutputDTO, error) {
	entities, count, err := u.AuthProviderQueries.List(dto.Search, dto.Limit, dto.Offset)
	return AuthProviderListOutputDTO{
		Model:      entities,
		TotalCount: count,
	}, err
}

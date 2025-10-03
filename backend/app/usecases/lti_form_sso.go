package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

type LTIFormListSSOUC struct {
	LTIFormQueries *lti_query.LTIFormQueries
}

type LTIFormListSSORequest struct {
	JSONRPC string       `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string       `json:"method" default:"lti_form.sso_list" required:"true"`
	Params  *interface{} `json:"params,omitempty"`
	ID      string       `json:"id,omitempty" default:"1" required:"true"`
}

type LTIFormListSSOOutputDTO struct {
	Model []models.LTIFormListItem `json:"model" validate:"required"`
}

type LTIFormListSSOResponse struct {
	JSONRPC string                  `json:"jsonrpc" default:"2.0" required:"true"`
	Result  LTIFormListSSOOutputDTO `json:"result,omitempty"`
	Error   interface{}             `json:"error,omitempty"`
	ID      string                  `json:"id,omitempty" default:"1" required:"true"`
}

func (u *LTIFormListSSOUC) Execute() (LTIFormListOutputDTO, error) {
	entities, err := u.LTIFormQueries.SSOURLList()
	return LTIFormListOutputDTO{
		Model: entities,
	}, err
}

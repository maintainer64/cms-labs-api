package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

type LTIFormEditUC struct {
	LTIFormQueries *lti_query.LTIFormQueries
}

type LTIFormEditInputDTO struct {
	ID            uint    `json:"id"`
	ClientID      string  `json:"client_id" validate:"required"`
	DeploymentID  string  `json:"deployment_id" validate:"required"`
	BaseURI       string  `json:"base_uri" validate:"required"`
	AuthTokenURI  string  `json:"auth_token_uri" validate:"required"`
	AuthLoginURI  string  `json:"auth_login_uri" validate:"required"`
	KeySetURI     string  `json:"key_set_uri" validate:"required"`
	TargetLinkURI string  `json:"target_link_uri" validate:"required"`
	Name          string  `json:"name" validate:"required"`
	SSOURL        *string `json:"sso_url"`
}

type LTIFormEditRequest struct {
	JSONRPC string              `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string              `json:"method" default:"lti_form.upsert" required:"true"`
	Params  LTIFormEditInputDTO `json:"params,omitempty"`
	ID      string              `json:"id,omitempty" default:"1" required:"true"`
}

type LTIFormEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIFormEditResponse struct {
	JSONRPC string               `json:"jsonrpc" default:"2.0" required:"true"`
	Result  LTIFormEditOutputDTO `json:"result,omitempty"`
	Error   interface{}          `json:"error,omitempty"`
	ID      string               `json:"id,omitempty" default:"1" required:"true"`
}

func (u *LTIFormEditUC) Execute(dto LTIFormEditInputDTO) (LTIFormEditOutputDTO, error) {
	entity := &models.LTIForm{}
	entity.ID = dto.ID
	entity.Name = dto.Name
	entity.LTIClientID = dto.ClientID
	entity.LTIDeploymentID = dto.DeploymentID
	entity.BaseURI = dto.BaseURI
	entity.LTIAuthTokenURI = dto.AuthTokenURI
	entity.LTIAuthLoginURI = dto.AuthLoginURI
	entity.KeySetURI = dto.KeySetURI
	entity.TargetLinkURI = dto.TargetLinkURI
	entity.SSOURL = dto.SSOURL
	err := u.LTIFormQueries.Upsert(entity)
	return LTIFormEditOutputDTO{ID: entity.ID}, err
}

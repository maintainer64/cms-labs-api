package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type LTIFormEditUC struct {
	LTIFormQueries *lti_query.LTIFormQueries
}

type LTIFormEditInputDTO struct {
	ID            uint   `json:"id"`
	ClientID      string `json:"client_id" validate:"required"`
	DeploymentID  string `json:"deployment_id" validate:"required"`
	BaseURI       string `json:"base_uri" validate:"required"`
	AuthTokenURI  string `json:"auth_token_uri" validate:"required"`
	AuthLoginURI  string `json:"auth_login_uri" validate:"required"`
	KeySetURI     string `json:"key_set_uri" validate:"required"`
	TargetLinkURI string `json:"target_link_uri" validate:"required"`
	Name          string `json:"name" validate:"required"`
}

type LTIFormEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIFormEditResponse = response.Response[LTIFormEditOutputDTO]

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
	err := u.LTIFormQueries.Upsert(entity)
	return LTIFormEditOutputDTO{ID: entity.ID}, err
}

package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type LTIFormListSSOUC struct {
	LTIFormQueries *lti_query.LTIFormQueries
}

type LTIFormListSSOOutputDTO struct {
	Model []models.LTIFormListItem `json:"model" validate:"required"`
}

type LTIFormListSSOResponse = response.Response[LTIFormListSSOOutputDTO]

func (u *LTIFormListSSOUC) Execute() (LTIFormListOutputDTO, error) {
	entities, err := u.LTIFormQueries.SSOURLList()
	return LTIFormListOutputDTO{
		Model: entities,
	}, err
}

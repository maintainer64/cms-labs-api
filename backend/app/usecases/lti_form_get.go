package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type LtiFormGetUC struct {
	LTIFromQuery *queries.LTIFromQueries
}

type LtiFormGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LtiFormGetOutputDTO struct {
	Model models.LTIForm `json:"model" required:"true"`
}

type LtiFormGetResponse = Response[LtiFormGetOutputDTO]

func (u *LtiFormGetUC) Execute(dto LtiFormGetInputDTO) (LtiFormGetOutputDTO, error) {
	form, err := u.LTIFromQuery.Get(dto.ID)
	return LtiFormGetOutputDTO{
		Model: form,
	}, err
}

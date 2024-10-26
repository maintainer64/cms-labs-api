package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

type LTIFormGetUC struct {
	LTIFormQueries *lti_query.LTIFormQueries
}

type LTIFormGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIFormGetOutputDTO struct {
	Model models.LTIForm `json:"model" required:"true"`
}

type LTIFormGetResponse = Response[LTIFormGetOutputDTO]

func (u *LTIFormGetUC) Execute(dto LTIFormGetInputDTO) (LTIFormGetOutputDTO, error) {
	form, err := u.LTIFormQueries.Get(dto.ID)
	return LTIFormGetOutputDTO{
		Model: form,
	}, err
}

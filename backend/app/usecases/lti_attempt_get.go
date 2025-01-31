package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type LTIAttemptGetUC struct {
	LTIAttemptQueries *queries.LTIAttemptQueries
}

type LTIAttemptGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIAttemptGetOutputDTO struct {
	Model models.LTIAttempt `json:"model" required:"true"`
}

type LTIAttemptGetResponse = response.Response[LTIAttemptGetOutputDTO]

func (u *LTIAttemptGetUC) Execute(dto LTIAttemptGetInputDTO) (LTIAttemptGetOutputDTO, error) {
	form, err := u.LTIAttemptQueries.Get(dto.ID)
	return LTIAttemptGetOutputDTO{
		Model: form,
	}, err
}

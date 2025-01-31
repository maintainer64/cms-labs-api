package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type PNETServerGetUC struct {
	PNETServerQueries *queries.PNETServerQueries
}

type PNETServerGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type PNETServerGetOutputDTO struct {
	Model models.PNETServer `json:"model" required:"true"`
}

type PNETServerGetResponse = response.Response[PNETServerGetOutputDTO]

func (u *PNETServerGetUC) Execute(dto PNETServerGetInputDTO) (PNETServerGetOutputDTO, error) {
	form, err := u.PNETServerQueries.Get(dto.ID)
	return PNETServerGetOutputDTO{
		Model: form,
	}, err
}

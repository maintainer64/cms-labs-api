package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type PNETServerDeleteUC struct {
	PNETServerQueries *queries.PNETServerQueries
}

type PNETServerDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type PNETServerDeleteResponse = response.Response[PNETServerDeleteInputDTO]

func (u *PNETServerDeleteUC) Execute(dto PNETServerDeleteInputDTO) (PNETServerDeleteInputDTO, error) {
	err := u.PNETServerQueries.Delete(dto.ID)
	return dto, err
}

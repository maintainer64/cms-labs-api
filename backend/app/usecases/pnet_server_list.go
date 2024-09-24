package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type PNETServerListUC struct {
	PNETServerQueries *queries.PNETServerQueries
}

type PNETServerListInputDTO struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type PNETServerListOutputDTO struct {
	Model []models.PNETServerListItem `json:"model" validate:"required"`
}

type PNETServerListResponse = Response[LtiFormListOutputDTO]

func (u *PNETServerListUC) Execute(dto PNETServerListInputDTO) (PNETServerListOutputDTO, error) {
	entities, err := u.PNETServerQueries.List(dto.Limit, dto.Offset)
	return PNETServerListOutputDTO{
		Model: entities,
	}, err
}

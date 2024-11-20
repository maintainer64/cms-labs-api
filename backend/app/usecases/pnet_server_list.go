package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type PNETServerListUC struct {
	PNETServerQueries *queries.PNETServerQueries
}

type PNETServerListInputDTO struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type PNETServerListOutputDTO struct {
	Model      []models.PNETServerListItem `json:"model" validate:"required"`
	TotalCount int64                       `json:"total_count" validate:"required"`
}

type PNETServerListResponse = Response[PNETServerListOutputDTO]

func (u *PNETServerListUC) Execute(dto PNETServerListInputDTO) (PNETServerListOutputDTO, error) {
	entities, count, err := u.PNETServerQueries.List(dto.Search, dto.Limit, dto.Offset)
	return PNETServerListOutputDTO{
		Model:      entities,
		TotalCount: count,
	}, err
}

package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type LTIRoutingListUC struct {
	LTIRoutingQueries *queries.LTIRoutingQueries
}

type LTIRoutingListInputDTO struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type LTIRoutingListOutputDTO struct {
	Model      []models.LTIRoutingListItem `json:"model" validate:"required"`
	TotalCount int64                       `json:"total_count" validate:"required"`
}

type LTIRoutingListResponse = Response[LTIRoutingListOutputDTO]

func (u *LTIRoutingListUC) Execute(dto LTIRoutingListInputDTO) (LTIRoutingListOutputDTO, error) {
	entities, count, err := u.LTIRoutingQueries.List(dto.Search, dto.Limit, dto.Offset)
	return LTIRoutingListOutputDTO{
		Model:      entities,
		TotalCount: count,
	}, err
}

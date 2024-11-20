package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

type LTIFormListUC struct {
	LTIFormQueries *lti_query.LTIFormQueries
}

type LTIFormListInputDTO struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type LTIFormListOutputDTO struct {
	Model      []models.LTIFormListItem `json:"model" validate:"required"`
	TotalCount int64                    `json:"total_count" validate:"required"`
}

type LTIFormListResponse = Response[LTIFormListOutputDTO]

func (u *LTIFormListUC) Execute(dto LTIFormListInputDTO) (LTIFormListOutputDTO, error) {
	entities, count, err := u.LTIFormQueries.List(dto.Search, dto.Limit, dto.Offset)
	return LTIFormListOutputDTO{
		Model:      entities,
		TotalCount: count,
	}, err
}

package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type LTIAttemptListUC struct {
	LTIAttemptQueries *queries.LTIAttemptQueries
}

type LTIAttemptListInputDTO struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type LTIAttemptListOutputDTO struct {
	Model      []models.LTIAttemptListItem `json:"model" validate:"required"`
	TotalCount int64                       `json:"total_count" validate:"required"`
}

type LTIAttemptListResponse = Response[LTIAttemptListOutputDTO]

func (u *LTIAttemptListUC) Execute(dto LTIAttemptListInputDTO) (LTIAttemptListOutputDTO, error) {
	entities, count, err := u.LTIAttemptQueries.List(dto.Search, dto.Limit, dto.Offset)
	return LTIAttemptListOutputDTO{
		Model:      entities,
		TotalCount: count,
	}, err
}

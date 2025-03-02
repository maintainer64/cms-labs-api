package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type LTIAttemptListUC struct {
	LTIAttemptQueries *queries.LTIAttemptQueries
}

type LTIAttemptListInputDTO struct {
	UserIds []uint `json:"user_ids"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
}

type LTIAttemptListOutputDTO struct {
	Model []models.LTIAttemptListItem `json:"model" validate:"required"`
}

type LTIAttemptListResponse = response.Response[LTIAttemptListOutputDTO]

func (u *LTIAttemptListUC) Execute(dto LTIAttemptListInputDTO) (LTIAttemptListOutputDTO, error) {
	entities, err := u.LTIAttemptQueries.List(dto.UserIds, dto.Limit, dto.Offset)
	return LTIAttemptListOutputDTO{
		Model: entities,
	}, err
}

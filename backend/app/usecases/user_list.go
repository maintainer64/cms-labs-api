package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type UserListUC struct {
	UserQueries *queries.UserQueries
}

type UserListInputDTO struct {
	Search  string `json:"search"`
	UserIds []uint `json:"user_ids"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
}

type UserListOutputDTO struct {
	Model      []models.UserListItem `json:"model" validate:"required"`
	TotalCount int64                 `json:"total_count" validate:"required"`
}

type UserListResponse = response.Response[UserListOutputDTO]

func (u *UserListUC) Execute(dto UserListInputDTO) (UserListOutputDTO, error) {
	entities, count, err := u.UserQueries.List(dto.Search, dto.UserIds, dto.Limit, dto.Offset)
	return UserListOutputDTO{
		Model:      entities,
		TotalCount: count,
	}, err
}

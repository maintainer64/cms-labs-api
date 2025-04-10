package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type UserListUC struct {
	UserQueries *queries.UserQueries
	RoleQueries *queries.RoleQueries
}

type UserListInputDTO struct {
	Search  string `json:"search"`
	UserIds []uint `json:"user_ids"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
}

type UserListModel struct {
	Model models.UserListItem `json:"model" validate:"required"`
	Roles []uint              `json:"roles"`
}
type UserListOutputDTO struct {
	Model      []UserListModel `json:"model" validate:"required"`
	TotalCount int64           `json:"total_count" validate:"required"`
}

type UserListResponse = response.Response[UserListOutputDTO]

func (u *UserListUC) Execute(dto UserListInputDTO) (UserListOutputDTO, error) {
	entities, count, err := u.UserQueries.List(dto.Search, dto.UserIds, dto.Limit, dto.Offset)
	if err != nil {
		return UserListOutputDTO{}, err
	}
	usersIds := make([]uint, 0)
	for _, entity := range entities {
		usersIds = append(usersIds, entity.ID)
	}
	hmap, err := u.RoleQueries.GetByRelationUsersIds(usersIds)
	if err != nil {
		return UserListOutputDTO{}, err
	}
	result := UserListOutputDTO{TotalCount: count}
	for _, entity := range entities {
		item := UserListModel{
			Model: entity,
		}
		if val, ok := hmap[entity.ID]; ok {
			item.Roles = val
		}
		result.Model = append(result.Model, item)
	}
	return result, err
}

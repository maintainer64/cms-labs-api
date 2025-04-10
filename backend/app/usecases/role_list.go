package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type RoleListUC struct {
	RoleQueries *queries.RoleQueries
}

type RoleListInputDTO struct {
}

type RoleListOutputDTO struct {
	Model      []models.Role `json:"model" validate:"required"`
	TotalCount int           `json:"total_count" validate:"required"`
}

type RoleListResponse = response.Response[RoleListOutputDTO]

func (u *RoleListUC) Execute(dto RoleListInputDTO) (RoleListOutputDTO, error) {
	entities, err := u.RoleQueries.List()
	return RoleListOutputDTO{
		Model:      entities,
		TotalCount: len(entities),
	}, err
}

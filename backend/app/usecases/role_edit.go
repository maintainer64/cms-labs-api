package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type RoleEditUC struct {
	RoleQueries *queries.RoleQueries
}

type RoleEditInputDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type RoleEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type RoleEditResponse = response.Response[UserEditOutputDTO]

func (u *RoleEditUC) Execute(dto RoleEditInputDTO) (RoleEditOutputDTO, error) {
	entity := &models.Role{}
	entity.ID = dto.ID
	entity.Name = dto.Name
	entity.Code = dto.Code
	err := u.RoleQueries.Upsert(entity)
	return RoleEditOutputDTO{ID: entity.ID}, err
}

package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type RoleDeleteUC struct {
	RoleQueries *queries.RoleQueries
}

type RoleDeleteInputDTO struct {
	ID uint `json:"id" validate:"required"`
}

type RoleDeleteResponse = response.Response[RoleDeleteInputDTO]

func (u *RoleDeleteUC) Execute(dto RoleDeleteInputDTO) (RoleDeleteInputDTO, error) {
	err := u.RoleQueries.Delete(dto.ID)
	return RoleDeleteInputDTO{ID: dto.ID}, err
}

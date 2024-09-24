package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type LtiFormEditUC struct {
	LTIFromQuery *queries.LTIFromQueries
}

type LtiFormEditInputDTO struct {
	ID      int    `json:"id"`
	Name    string `json:"name" validate:"required"`
	Version string `json:"version" validate:"required"`
}

type LtiFormEditOutputDTO struct {
	ID int `json:"id" validate:"required"`
}

type LtiFormEditResponse = Response[LtiFormEditOutputDTO]

func (u *LtiFormEditUC) Execute(dto LtiFormEditInputDTO) (LtiFormEditOutputDTO, error) {
	entity := &models.LTIForm{}
	entity.Name = dto.Name
	entity.Version = dto.Version
	entity.ID = uint(dto.ID)
	err := u.LTIFromQuery.Upsert(entity)
	return LtiFormEditOutputDTO{ID: int(entity.ID)}, err
}

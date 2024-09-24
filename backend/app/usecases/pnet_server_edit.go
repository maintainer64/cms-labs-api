package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type PNETServerEditUC struct {
	PNETServerQueries *queries.PNETServerQueries
}

type PNETServerEditInputDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Url      string `json:"url" validate:"required"`
	IsActive bool   `json:"is_active"`
	UnitRate uint   `json:"unit_rate"`
	Token    string `json:"token" validate:"required"`
}

type PNETServerEditOutputDTO struct {
	ID uint `json:"id" validate:"required"`
}

type PNETServerEditResponse = Response[PNETServerEditOutputDTO]

func (u *PNETServerEditUC) Execute(dto PNETServerEditInputDTO) (PNETServerEditOutputDTO, error) {
	entity := &models.PNETServer{}
	entity.ID = dto.ID
	entity.Name = dto.Name
	entity.Url = dto.Url
	entity.IsActive = dto.IsActive
	entity.UnitRate = dto.UnitRate
	entity.Token = dto.Token
	err := u.PNETServerQueries.Upsert(entity)
	return PNETServerEditOutputDTO{ID: entity.ID}, err
}

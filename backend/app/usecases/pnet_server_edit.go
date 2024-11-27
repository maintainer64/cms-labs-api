package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type PNETServerEditUC struct {
	PNETServerQueries *queries.PNETServerQueries
}

type PNETServerEditInputDTO struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name" validate:"required"`
	Url                  string `json:"url" validate:"required"`
	IsActive             bool   `json:"is_active"`
	MinutesForDisconnect int    `json:"minutes_for_disconnect"`
	MaxCountUsersLimit   int    `json:"max_count_users_limit"`
	UnitRate             int    `json:"unit_rate"`
}

type PNETServerEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type PNETServerEditResponse = Response[PNETServerEditOutputDTO]

func (u *PNETServerEditUC) Execute(dto PNETServerEditInputDTO) (PNETServerEditOutputDTO, error) {
	entity := &models.PNETServer{}
	entity.ID = dto.ID
	entity.Name = dto.Name
	entity.Url = dto.Url
	entity.IsActive = dto.IsActive
	entity.UnitRate = dto.UnitRate
	entity.MinutesForDisconnect = dto.MinutesForDisconnect
	entity.MaxCountUsersLimit = dto.MaxCountUsersLimit
	err := u.PNETServerQueries.Upsert(entity)
	return PNETServerEditOutputDTO{ID: entity.ID}, err
}

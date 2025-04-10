package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type PNETServerEditUC struct {
	PNETServerQueries *queries.PNETServerQueries
	RoleQueries       *queries.RoleQueries
}

type PNETServerEditInputDTO struct {
	ID                   uint   `json:"id"`
	ClientID             string `json:"client_id"`
	Type                 string `json:"type" validate:"required"`
	Name                 string `json:"name" validate:"required"`
	Url                  string `json:"url" validate:"required"`
	IsActive             bool   `json:"is_active"`
	MinutesForDisconnect int    `json:"minutes_for_disconnect"`
	MaxCountUsersLimit   int    `json:"max_count_users_limit"`
	UnitRate             int    `json:"unit_rate"`
	Roles                []uint `json:"roles"`
}

type PNETServerEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type PNETServerEditResponse = response.Response[PNETServerEditOutputDTO]

func (u *PNETServerEditUC) Execute(dto PNETServerEditInputDTO) (PNETServerEditOutputDTO, error) {
	entity := &models.PNETServer{}
	entity.ID = dto.ID
	entity.ClientID = dto.ClientID
	entity.Type = dto.Type
	entity.Name = dto.Name
	entity.Url = dto.Url
	entity.IsActive = dto.IsActive
	entity.UnitRate = dto.UnitRate
	entity.MinutesForDisconnect = dto.MinutesForDisconnect
	entity.MaxCountUsersLimit = dto.MaxCountUsersLimit
	err := u.PNETServerQueries.Upsert(entity)
	if err != nil {
		return PNETServerEditOutputDTO{}, err
	}
	err = u.RoleQueries.SetByServerId(entity.ID, dto.Roles)
	return PNETServerEditOutputDTO{ID: entity.ID}, err
}

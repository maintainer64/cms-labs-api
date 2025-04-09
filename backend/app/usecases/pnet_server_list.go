package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type PNETServerListUC struct {
	PNETServerQueries *queries.PNETServerQueries
	RoleQueries       *queries.RoleQueries
}

type PNETServerListInputDTO = queries.PNETServerQueriesListDTO

type PNETServerListModel struct {
	Model models.PNETServerListItem `json:"model" validate:"required"`
	Roles []uint                    `json:"roles"`
}

type PNETServerListOutputDTO struct {
	Model      []PNETServerListModel `json:"model" validate:"required"`
	TotalCount int64                 `json:"total_count" validate:"required"`
}

type PNETServerListResponse = response.Response[PNETServerListOutputDTO]

func (u *PNETServerListUC) Execute(dto PNETServerListInputDTO) (PNETServerListOutputDTO, error) {
	entities, count, err := u.PNETServerQueries.List(dto)
	serverIDS := make([]uint, 0)
	for _, entity := range entities {
		serverIDS = append(serverIDS, entity.ID)
	}
	hmap, err := u.RoleQueries.GetByRelationServerIds(serverIDS)
	if err != nil {
		return PNETServerListOutputDTO{}, err
	}
	result := PNETServerListOutputDTO{TotalCount: count}
	for _, entity := range entities {
		item := PNETServerListModel{
			Model: entity,
		}
		if val, ok := hmap[entity.ID]; ok {
			item.Roles = val
		}
		result.Model = append(result.Model, item)
	}
	return result, err
}

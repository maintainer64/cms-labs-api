package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type PNETServerGetUC struct {
	PNETServerQueries *queries.PNETServerQueries
	RoleQueries       *queries.RoleQueries
}

type PNETServerGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type PNETServerGetOutputDTO struct {
	Model models.PNETServer `json:"model" required:"true"`
	Roles []uint            `json:"roles"`
}

type PNETServerGetResponse = response.Response[PNETServerGetOutputDTO]

func (u *PNETServerGetUC) Execute(dto PNETServerGetInputDTO) (PNETServerGetOutputDTO, error) {
	form, err := u.PNETServerQueries.Get(dto.ID)
	if err != nil {
		return PNETServerGetOutputDTO{}, err
	}
	result := PNETServerGetOutputDTO{
		Model: form,
	}
	roles, err := u.RoleQueries.GetByRelationServerIds([]uint{dto.ID})
	if err != nil {
		return result, err
	}
	if val, ok := roles[dto.ID]; ok {
		result.Roles = val
	}
	return result, err
}

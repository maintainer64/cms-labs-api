package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type PNETServerGetUC struct {
	PNETServerQueries *queries.PNETServerQueries
	RoleQueries       *queries.RoleQueries
}

type PNETServerGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type PNETServerGetRequest struct {
	JSONRPC string                `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string                `json:"method" default:"server.get" required:"true"`
	Params  PNETServerGetInputDTO `json:"params,omitempty"`
	ID      string                `json:"id,omitempty" default:"1" required:"true"`
}

type PNETServerGetOutputDTO struct {
	Model models.PNETServer `json:"model" required:"true"`
	Roles []uint            `json:"roles"`
}

type PNETServerGetResponse struct {
	JSONRPC string                 `json:"jsonrpc" default:"2.0" required:"true"`
	Result  PNETServerGetOutputDTO `json:"result,omitempty"`
	Error   interface{}            `json:"error,omitempty"`
	ID      string                 `json:"id,omitempty" default:"1" required:"true"`
}

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

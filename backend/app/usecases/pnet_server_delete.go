package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type PNETServerDeleteUC struct {
	PNETServerQueries *queries.PNETServerQueries
}

type PNETServerDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type PNETServerDeleteRequest struct {
	JSONRPC string                   `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string                   `json:"method" default:"server.list" required:"true"`
	Params  PNETServerDeleteInputDTO `json:"params,omitempty"`
	ID      string                   `json:"id,omitempty" default:"1" required:"true"`
}

type PNETServerDeleteResponse struct {
	JSONRPC string                   `json:"jsonrpc" default:"2.0" required:"true"`
	Result  PNETServerDeleteInputDTO `json:"result,omitempty"`
	Error   interface{}              `json:"error,omitempty"`
	ID      string                   `json:"id,omitempty" default:"1" required:"true"`
}

func (u *PNETServerDeleteUC) Execute(dto PNETServerDeleteInputDTO) (PNETServerDeleteInputDTO, error) {
	err := u.PNETServerQueries.Delete(dto.ID)
	return dto, err
}

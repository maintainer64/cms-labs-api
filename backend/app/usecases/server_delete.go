package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type ServerDeleteUC struct {
	ServerQueries *queries.ServerQueries
}

type ServerDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type ServerDeleteRequest struct {
	JSONRPC string               `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string               `json:"method" default:"server.list" required:"true"`
	Params  ServerDeleteInputDTO `json:"params,omitempty"`
	ID      string               `json:"id,omitempty" default:"1" required:"true"`
}

type ServerDeleteResponse struct {
	JSONRPC string               `json:"jsonrpc" default:"2.0" required:"true"`
	Result  ServerDeleteInputDTO `json:"result,omitempty"`
	Error   interface{}          `json:"error,omitempty"`
	ID      string               `json:"id,omitempty" default:"1" required:"true"`
}

func (u *ServerDeleteUC) Execute(dto ServerDeleteInputDTO) (ServerDeleteInputDTO, error) {
	err := u.ServerQueries.Delete(dto.ID)
	return dto, err
}

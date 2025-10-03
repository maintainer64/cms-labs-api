package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type CurlRequestDeleteUC struct {
	CurlRequestQueries *queries.CurlRequestQueries
}

type CurlRequestDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type CurlRequestDeleteRequest struct {
	JSONRPC string                  `json:"jsonrpc" default:"2.0" validate:"required"`
	Method  string                  `json:"method" default:"curl_request.delete" validate:"required"`
	Params  CurlRequestEditInputDTO `json:"params,omitempty"`
	ID      string                  `json:"id,omitempty" default:"1" validate:"required"`
}

type CurlRequestDeleteResponse struct {
	JSONRPC string                  `json:"jsonrpc" default:"2.0" validate:"required"`
	Result  CurlRequestEditInputDTO `json:"result,omitempty"`
	Error   interface{}             `json:"error,omitempty"`
	ID      string                  `json:"id,omitempty" default:"1" validate:"required"`
}

func (u *CurlRequestDeleteUC) Execute(dto CurlRequestDeleteInputDTO) (CurlRequestDeleteInputDTO, error) {
	err := u.CurlRequestQueries.Delete(dto.ID)
	return dto, err
}

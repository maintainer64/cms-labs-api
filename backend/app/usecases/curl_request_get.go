package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type CurlRequestGetUC struct {
	CurlRequestQueries *queries.CurlRequestQueries
}

type CurlRequestGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type CurlRequestGetRequest struct {
	JSONRPC string                 `json:"jsonrpc" default:"2.0" validate:"required"`
	Method  string                 `json:"method" default:"curl_request.get" validate:"required"`
	Params  CurlRequestGetInputDTO `json:"params,omitempty"`
	ID      string                 `json:"id,omitempty" default:"1" validate:"required"`
}

type CurlRequestGetOutputDTO struct {
	Model models.CurlRequest `json:"model" required:"true"`
}

type CurlRequestGetResponse struct {
	JSONRPC string                  `json:"jsonrpc" default:"2.0" validate:"required"`
	Result  CurlRequestGetOutputDTO `json:"result,omitempty"`
	Error   interface{}             `json:"error,omitempty"`
	ID      string                  `json:"id,omitempty" default:"1" validate:"required"`
}

func (u *CurlRequestGetUC) Execute(dto CurlRequestGetInputDTO) (CurlRequestGetOutputDTO, error) {
	form, err := u.CurlRequestQueries.Get(dto.ID)
	if err != nil {
		return CurlRequestGetOutputDTO{}, err
	}
	result := CurlRequestGetOutputDTO{
		Model: form,
	}
	return result, err
}

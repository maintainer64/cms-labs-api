package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type CurlRequestListUC struct {
	CurlRequestQueries *queries.CurlRequestQueries
}

type CurlRequestListRequest struct {
	JSONRPC string                            `json:"jsonrpc" default:"2.0" validate:"required"`
	Method  string                            `json:"method" default:"curl_request.list" validate:"required"`
	Params  queries.CurlRequestQueriesListDTO `json:"params,omitempty"`
	ID      string                            `json:"id,omitempty" default:"1" validate:"required"`
}

type CurlRequestListOutputDTO struct {
	Model      []models.CurlRequestListItem `json:"model" validate:"required"`
	TotalCount int64                        `json:"total_count" validate:"required"`
}

type CurlRequestListResponse struct {
	JSONRPC string                   `json:"jsonrpc" default:"2.0" validate:"required"`
	Result  CurlRequestListOutputDTO `json:"result,omitempty"`
	Error   interface{}              `json:"error,omitempty"`
	ID      string                   `json:"id,omitempty" default:"1" validate:"required"`
}

func (u *CurlRequestListUC) Execute(dto queries.CurlRequestQueriesListDTO) (CurlRequestListOutputDTO, error) {
	entities, count, err := u.CurlRequestQueries.List(dto)
	if err != nil {
		return CurlRequestListOutputDTO{}, err
	}
	result := CurlRequestListOutputDTO{TotalCount: count, Model: entities}
	return result, err
}

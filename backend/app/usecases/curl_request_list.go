package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type CurlRequestListUC struct {
	CurlRequestQueries *queries.CurlRequestQueries
}

type CurlRequestListInputDTO = queries.CurlRequestQueriesListDTO

type CurlRequestListOutputDTO struct {
	Model      []models.CurlRequestListItem `json:"model" validate:"required"`
	TotalCount int64                        `json:"total_count" validate:"required"`
}

type CurlRequestListResponse = response.Response[CurlRequestListOutputDTO]

func (u *CurlRequestListUC) Execute(dto CurlRequestListInputDTO) (CurlRequestListOutputDTO, error) {
	entities, count, err := u.CurlRequestQueries.List(dto)
	if err != nil {
		return CurlRequestListOutputDTO{}, err
	}
	result := CurlRequestListOutputDTO{TotalCount: count, Model: entities}
	return result, err
}

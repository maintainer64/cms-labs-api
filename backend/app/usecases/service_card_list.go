package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type ServiceCardListUC struct {
	ServiceCardQueries *queries.ServiceCardQueries
}

type ServiceCardListInputDTO struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type ServiceCardListOutputDTO struct {
	Model      []models.ServiceCardListItem `json:"model" validate:"required"`
	TotalCount int64                        `json:"total_count" validate:"required"`
}

type ServiceCardListResponse = response.Response[ServiceCardListOutputDTO]

func (u *ServiceCardListUC) Execute(dto ServiceCardListInputDTO) (ServiceCardListOutputDTO, error) {
	entities, count, err := u.ServiceCardQueries.List(dto.Search, dto.Limit, dto.Offset)
	return ServiceCardListOutputDTO{
		Model:      entities,
		TotalCount: count,
	}, err
}

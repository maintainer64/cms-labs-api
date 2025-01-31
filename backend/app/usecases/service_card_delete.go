package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type ServiceCardDeleteUC struct {
	ServiceCardQueries *queries.ServiceCardQueries
}

type ServiceCardDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type ServiceCardDeleteResponse = response.Response[ServiceCardDeleteInputDTO]

func (u *ServiceCardDeleteUC) Execute(dto ServiceCardDeleteInputDTO) (ServiceCardDeleteInputDTO, error) {
	err := u.ServiceCardQueries.Delete(dto.ID)
	return dto, err
}

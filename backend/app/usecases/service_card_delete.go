package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type ServiceCardDeleteUC struct {
	ServiceCardQueries *queries.ServiceCardQueries
}

type ServiceCardDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type ServiceCardDeleteResponse = Response[ServiceCardDeleteInputDTO]

func (u *ServiceCardDeleteUC) Execute(dto ServiceCardDeleteInputDTO) (ServiceCardDeleteInputDTO, error) {
	err := u.ServiceCardQueries.Delete(dto.ID)
	return dto, err
}

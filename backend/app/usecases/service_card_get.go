package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type ServiceCardGetUC struct {
	ServiceCardQueries *queries.ServiceCardQueries
}

type ServiceCardGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type ServiceCardGetOutputDTO struct {
	Model models.ServiceCard `json:"model" required:"true"`
}

type ServiceCardGetResponse = Response[ServiceCardGetOutputDTO]

func (u *ServiceCardGetUC) Execute(dto ServiceCardGetInputDTO) (ServiceCardGetOutputDTO, error) {
	form, err := u.ServiceCardQueries.Get(dto.ID)
	return ServiceCardGetOutputDTO{
		Model: form,
	}, err
}

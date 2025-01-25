package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type ServiceCardEditUC struct {
	ServiceCardQueries *queries.ServiceCardQueries
}

type ServiceCardEditInputDTO struct {
	ID          uint   `json:"id"`
	ImageUrl    string `json:"image_url"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       uint   `json:"order"`
	IsActive    bool   `json:"is_active"`
}

type ServiceCardEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type ServiceCardEditResponse = Response[ServiceCardEditOutputDTO]

func (u *ServiceCardEditUC) Execute(dto ServiceCardEditInputDTO) (ServiceCardEditOutputDTO, error) {
	entity := &models.ServiceCard{}
	entity.ID = dto.ID
	entity.ImageUrl = dto.ImageUrl
	entity.Url = dto.Url
	entity.Name = dto.Name
	entity.Description = dto.Description
	entity.Order = dto.Order
	entity.IsActive = dto.IsActive
	err := u.ServiceCardQueries.Upsert(entity)
	return ServiceCardEditOutputDTO{ID: entity.ID}, err
}

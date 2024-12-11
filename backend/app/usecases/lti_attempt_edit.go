package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type LTIAttemptEditUC struct {
	LTIAttemptQueries *queries.LTIAttemptQueries
}

type LTIAttemptEditInputDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Url      string `json:"url" validate:"required"`
	IsActive bool   `json:"is_active"`
	UnitRate uint   `json:"unit_rate"`
	Token    string `json:"token" validate:"required"`
}

type LTIAttemptEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIAttemptEditResponse = Response[LTIAttemptEditOutputDTO]

func (u *LTIAttemptEditUC) Execute(dto LTIAttemptEditInputDTO) (LTIAttemptEditOutputDTO, error) {
	entity := &models.LTIAttempt{}
	entity.ID = dto.ID
	/*
	   TODO: Add attributes set to upsert LTIAttempt
	*/
	err := u.LTIAttemptQueries.Upsert(entity)
	return LTIAttemptEditOutputDTO{ID: entity.ID}, err
}

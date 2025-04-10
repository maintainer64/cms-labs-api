package usecases

import (
	"time"

	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type LTIAttemptEditUC struct {
	LTIAttemptQueries *queries.LTIAttemptQueries
}

type LTIAttemptEditInputDTO struct {
	ID           uint      `json:"id"`
	PNETServerID uint      `json:"pnet_server_id"`
	ExpiredAt    time.Time `json:"expired_at" validate:"required"`
}

type LTIAttemptEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIAttemptEditResponse = response.Response[LTIAttemptEditOutputDTO]

func (u *LTIAttemptEditUC) Execute(dto LTIAttemptEditInputDTO) (LTIAttemptEditOutputDTO, error) {
	entity, err := u.LTIAttemptQueries.Get(dto.ID)
	if err != nil {
		return LTIAttemptEditOutputDTO{}, err
	}
	entity.ID = dto.ID
	entity.PNETServerID = &dto.PNETServerID
	entity.ExpiredAt = dto.ExpiredAt
	err = u.LTIAttemptQueries.Upsert(&entity)
	return LTIAttemptEditOutputDTO{ID: entity.ID}, err
}

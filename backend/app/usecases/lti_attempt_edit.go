package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gorm.io/datatypes"
)

type LTIAttemptEditUC struct {
	LTIAttemptQueries *queries.LTIAttemptQueries
}

type LTIAttemptEditInputDTO struct {
	ID           uint                                         `json:"id"`
	PNETServerID uint                                         `json:"pnet_server_id"`
	Status       string                                       `json:"status" validate:"required"`
	Result       *datatypes.JSONType[models.LTIAttemptResult] `json:"result" swaggertype:"object"`
}

type LTIAttemptEditRequest struct {
	JSONRPC string                 `json:"jsonrpc" default:"2.0" validate:"required"`
	Method  string                 `json:"method" default:"lti_attempt.update" validate:"required"`
	Params  LTIAttemptEditInputDTO `json:"params,omitempty"`
	ID      string                 `json:"id,omitempty" default:"1" validate:"required"`
}

type LTIAttemptEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIAttemptEditResponse struct {
	JSONRPC string                  `json:"jsonrpc" default:"2.0" validate:"required"`
	Result  LTIAttemptEditOutputDTO `json:"result,omitempty"`
	Error   interface{}             `json:"error,omitempty"`
	ID      string                  `json:"id,omitempty" default:"1" validate:"required"`
}

func (u *LTIAttemptEditUC) Execute(dto LTIAttemptEditInputDTO) (LTIAttemptEditOutputDTO, error) {
	entity, err := u.LTIAttemptQueries.Get(dto.ID)
	if err != nil {
		return LTIAttemptEditOutputDTO{}, err
	}
	entity.ID = dto.ID
	entity.PNETServerID = &dto.PNETServerID
	entity.Status = dto.Status
	entity.Result = dto.Result
	err = u.LTIAttemptQueries.Upsert(&entity)
	return LTIAttemptEditOutputDTO{ID: entity.ID}, err
}

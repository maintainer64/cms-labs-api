package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type LTIAttemptDeleteUC struct {
	LTIAttemptQueries *queries.LTIAttemptQueries
}

type LTIAttemptDeleteRequest struct {
	JSONRPC string                   `json:"jsonrpc" default:"2.0" validate:"required"`
	Method  string                   `json:"method" default:"lti_attempt.delete" validate:"required"`
	Params  LTIAttemptCreateInputDTO `json:"params,omitempty"`
	ID      string                   `json:"id,omitempty" default:"1" validate:"required"`
}

type LTIAttemptDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIAttemptDeleteResponse struct {
	JSONRPC string                   `json:"jsonrpc" default:"2.0" validate:"required"`
	Result  LTIAttemptDeleteInputDTO `json:"result,omitempty"`
	Error   interface{}              `json:"error,omitempty"`
	ID      string                   `json:"id,omitempty" default:"1" validate:"required"`
}

func (u *LTIAttemptDeleteUC) Execute(dto LTIAttemptDeleteInputDTO) (LTIAttemptDeleteInputDTO, error) {
	err := u.LTIAttemptQueries.Delete(dto.ID)
	return dto, err
}

package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

type LTIFormDeleteUC struct {
	LTIFormQueries *lti_query.LTIFormQueries
}

type LTIFormDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIFormDeleteRequest struct {
	JSONRPC string                `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string                `json:"method" default:"lti_form.delete" required:"true"`
	Params  LTIAttemptGetInputDTO `json:"params,omitempty"`
	ID      string                `json:"id,omitempty" default:"1" required:"true"`
}

type LTIFormDeleteResponse struct {
	JSONRPC string                `json:"jsonrpc" default:"2.0" required:"true"`
	Result  LTIFormDeleteInputDTO `json:"result,omitempty"`
	Error   interface{}           `json:"error,omitempty"`
	ID      string                `json:"id,omitempty" default:"1" required:"true"`
}

func (u *LTIFormDeleteUC) Execute(dto LTIFormDeleteInputDTO) (LTIFormDeleteInputDTO, error) {
	err := u.LTIFormQueries.Delete(dto.ID)
	return dto, err
}

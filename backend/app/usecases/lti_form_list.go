package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

type LTIFormListUC struct {
	LTIFormQueries *lti_query.LTIFormQueries
}

type LTIFormListInputDTO struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type LTIFormListRequest struct {
	JSONRPC string              `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string              `json:"method" default:"lti_form.list" required:"true"`
	Params  LTIFormListInputDTO `json:"params,omitempty"`
	ID      string              `json:"id,omitempty" default:"1" required:"true"`
}

type LTIFormListOutputDTO struct {
	Model      []models.LTIFormListItem `json:"model" validate:"required"`
	TotalCount int64                    `json:"total_count" validate:"required"`
}

type LTIFormListResponse struct {
	JSONRPC string              `json:"jsonrpc" default:"2.0" required:"true"`
	Result  LTIFormGetOutputDTO `json:"result,omitempty"`
	Error   interface{}         `json:"error,omitempty"`
	ID      string              `json:"id,omitempty" default:"1" required:"true"`
}

func (u *LTIFormListUC) Execute(dto LTIFormListInputDTO) (LTIFormListOutputDTO, error) {
	entities, count, err := u.LTIFormQueries.List(dto.Search, dto.Limit, dto.Offset)
	return LTIFormListOutputDTO{
		Model:      entities,
		TotalCount: count,
	}, err
}

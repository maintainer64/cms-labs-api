package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

// TargetGetUC – получение цели по ID (без проверки прав)
type TargetGetUC struct {
	TargetQueries *queries.TargetQueries
}

// TargetGetInputDTO – входные данные
type TargetGetInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type TargetGetRequest struct {
	JSONRPC string            `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string            `json:"method" default:"service_card.upsert" required:"true"`
	Params  TargetGetInputDTO `json:"params,omitempty"`
	ID      string            `json:"id,omitempty" default:"1" required:"true"`
}

type TargetGetResponse struct {
	JSONRPC string        `json:"jsonrpc" default:"2.0" required:"true"`
	Result  models.Target `json:"result,omitempty"`
	Error   interface{}   `json:"error,omitempty"`
	ID      string        `json:"id,omitempty" default:"1" required:"true"`
}

// Execute – возвращает полную модель Target
func (uc *TargetGetUC) Execute(dto TargetGetInputDTO) (models.Target, error) {
	return uc.TargetQueries.Get(dto.ID)
}

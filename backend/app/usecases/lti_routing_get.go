package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type LTIRoutingGetUC struct {
	LTIRoutingQueries *queries.LTIRoutingQueries
}

type LTIRoutingGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIRoutingGetOutputDTO struct {
	Model models.LTIRouting `json:"model" required:"true"`
}

type LTIRoutingGetResponse = response.Response[LTIRoutingGetOutputDTO]

func (u *LTIRoutingGetUC) Execute(dto LTIRoutingGetInputDTO) (LTIRoutingGetOutputDTO, error) {
	form, err := u.LTIRoutingQueries.Get(dto.ID)
	return LTIRoutingGetOutputDTO{
		Model: form,
	}, err
}

package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type LTIRoutingDeleteUC struct {
	LTIRoutingQueries *queries.LTIRoutingQueries
}

type LTIRoutingDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIRoutingDeleteResponse = response.Response[LTIRoutingDeleteInputDTO]

func (u *LTIRoutingDeleteUC) Execute(dto LTIRoutingDeleteInputDTO) (LTIRoutingDeleteInputDTO, error) {
	err := u.LTIRoutingQueries.Delete(dto.ID)
	return dto, err
}

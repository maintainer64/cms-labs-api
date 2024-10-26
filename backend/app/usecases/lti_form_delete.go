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

type LTIFormDeleteResponse = Response[LTIFormDeleteInputDTO]

func (u *LTIFormDeleteUC) Execute(dto LTIFormDeleteInputDTO) (LTIFormDeleteInputDTO, error) {
	err := u.LTIFormQueries.Delete(dto.ID)
	return dto, err
}

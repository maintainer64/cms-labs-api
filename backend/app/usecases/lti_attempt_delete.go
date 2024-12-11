package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type LTIAttemptDeleteUC struct {
	LTIAttemptQueries *queries.LTIAttemptQueries
}

type LTIAttemptDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIAttemptDeleteResponse = Response[LTIAttemptDeleteInputDTO]

func (u *LTIAttemptDeleteUC) Execute(dto LTIAttemptDeleteInputDTO) (LTIAttemptDeleteInputDTO, error) {
	err := u.LTIAttemptQueries.Delete(dto.ID)
	return dto, err
}

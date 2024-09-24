package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type LtiFormDeleteUC struct {
	LTIFromQuery *queries.LTIFromQueries
}

type LtiFormDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LtiFormDeleteResponse = Response[LtiFormDeleteInputDTO]

func (u *LtiFormDeleteUC) Execute(dto LtiFormDeleteInputDTO) (LtiFormDeleteInputDTO, error) {
	err := u.LTIFromQuery.Delete(dto.ID)
	return dto, err
}

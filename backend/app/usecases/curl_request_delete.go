package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type CurlRequestDeleteUC struct {
	CurlRequestQueries *queries.CurlRequestQueries
}

type CurlRequestDeleteInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type CurlRequestDeleteResponse = response.Response[CurlRequestDeleteInputDTO]

func (u *CurlRequestDeleteUC) Execute(dto CurlRequestDeleteInputDTO) (CurlRequestDeleteInputDTO, error) {
	err := u.CurlRequestQueries.Delete(dto.ID)
	return dto, err
}

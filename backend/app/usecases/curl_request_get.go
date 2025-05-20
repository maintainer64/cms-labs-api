package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type CurlRequestGetUC struct {
	CurlRequestQueries *queries.CurlRequestQueries
}

type CurlRequestGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type CurlRequestGetOutputDTO struct {
	Model models.CurlRequest `json:"model" required:"true"`
}

type CurlRequestGetResponse = response.Response[CurlRequestGetOutputDTO]

func (u *CurlRequestGetUC) Execute(dto CurlRequestGetInputDTO) (CurlRequestGetOutputDTO, error) {
	form, err := u.CurlRequestQueries.Get(dto.ID)
	if err != nil {
		return CurlRequestGetOutputDTO{}, err
	}
	result := CurlRequestGetOutputDTO{
		Model: form,
	}
	return result, err
}

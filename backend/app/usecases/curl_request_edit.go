package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type CurlRequestEditUC struct {
	CurlRequestQueries *queries.CurlRequestQueries
}

type CurlRequestEditInputDTO struct {
	ID         uint              `json:"id"`
	Name       string            `json:"name" validate:"required"`
	Method     string            `json:"method" validate:"required"`
	URL        string            `json:"url" validate:"required"`
	Headers    map[string]string `json:"headers" validate:"required"`
	Body       string            `json:"body"`
	Timeout    int64             `json:"timeout"`
	RawRequest string            `json:"raw_request"`
}

type CurlRequestEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type CurlRequestEditResponse = response.Response[CurlRequestEditOutputDTO]

func (u *CurlRequestEditUC) Execute(dto CurlRequestEditInputDTO) (CurlRequestEditOutputDTO, error) {
	entity := &models.CurlRequest{}
	entity.ID = dto.ID
	entity.Name = dto.Name
	entity.Method = dto.Method
	entity.URL = dto.URL
	entity.Headers = dto.Headers
	entity.Body = dto.Body
	entity.Raw = dto.RawRequest
	entity.Timeout = dto.Timeout
	err := u.CurlRequestQueries.Upsert(entity)
	if err != nil {
		return CurlRequestEditOutputDTO{}, err
	}
	return CurlRequestEditOutputDTO{ID: entity.ID}, err
}

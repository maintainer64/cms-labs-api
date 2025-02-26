package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type UNLFileGetUC struct {
	UNLFileQueries *queries.UNLFileQueries
}

type UNLFileGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type UNLFileGetOutputDTO struct {
	Model models.UNLFile `json:"model" required:"true"`
}

type UNLFileGetResponse = response.Response[UNLFileGetOutputDTO]

func (u *UNLFileGetUC) Execute(dto UNLFileGetInputDTO) (UNLFileGetOutputDTO, error) {
	form, err := u.UNLFileQueries.Get(dto.ID)
	return UNLFileGetOutputDTO{
		Model: form,
	}, err
}

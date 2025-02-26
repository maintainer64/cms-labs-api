package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type UNLFileListUC struct {
	UNLFileQueries *queries.UNLFileQueries
}

type UNLFileListInputDTO struct {
	Search string   `json:"search"`
	Type   []string `json:"type"`
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
}

type UNLFileListOutputDTO struct {
	Model []models.UNLFileListItem `json:"model" validate:"required"`
}

type UNLFileListResponse = response.Response[UNLFileListOutputDTO]

func (u *UNLFileListUC) Execute(dto UNLFileListInputDTO) (UNLFileListOutputDTO, error) {
	entities, err := u.UNLFileQueries.List(dto.Search, dto.Type, dto.Limit, dto.Offset)
	return UNLFileListOutputDTO{
		Model: entities,
	}, err
}

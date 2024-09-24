package usecases

import (
	"time"

	"github.com/thoas/go-funk"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type LtiFormListUC struct {
	LTIFromQuery *queries.LTIFromQueries
}

type LtiFormListInputDTO struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type LtiFormListItem struct {
	ID        uint      `json:"id" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Version   string    `json:"version" validate:"required"`
}

type LtiFormListOutputDTO struct {
	Model []LtiFormListItem `json:"model" validate:"required"`
}

type LtiFormListResponse = Response[LtiFormListOutputDTO]

func (u *LtiFormListUC) Execute(dto LtiFormListInputDTO) (LtiFormListOutputDTO, error) {
	entities, err := u.LTIFromQuery.List(dto.Limit, dto.Offset)
	forms := funk.Map(entities, func(form models.LTIForm) LtiFormListItem {
		return LtiFormListItem{
			ID:        form.ID,
			CreatedAt: form.CreatedAt,
			UpdatedAt: form.UpdatedAt,
			Name:      form.Name,
			Version:   form.Version,
		}
	}).([]LtiFormListItem)
	return LtiFormListOutputDTO{
		Model: forms,
	}, err
}

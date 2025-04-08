package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
)

type UserGetUC struct {
	UserQueries *queries.UserQueries
	RoleQueries *queries.RoleQueries
}

type UserGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type UserGetOutputDTO struct {
	Model models.User `json:"model" required:"true"`
	Roles []uint      `json:"roles"`
}

type UserGetResponse = response.Response[UserGetOutputDTO]

func (u *UserGetUC) Execute(dto UserGetInputDTO) (UserGetOutputDTO, error) {
	form, err := u.UserQueries.Get(dto.ID)
	if err != nil && err.Error() == queries.UserNotActive.Error() {
		err = nil
	}
	if err != nil {
		return UserGetOutputDTO{}, err
	}
	result := UserGetOutputDTO{
		Model: form,
	}
	roles, err := u.RoleQueries.GetByRelationUsersIds([]uint{dto.ID})
	if err != nil {
		return result, err
	}
	if val, ok := roles[dto.ID]; ok {
		result.Roles = val
	}
	return result, err
}

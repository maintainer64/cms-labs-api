package usecases

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type UserGetUC struct {
	UserQueries *queries.UserQueries
}

type UserGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type UserGetOutputDTO struct {
	Model models.User `json:"model" required:"true"`
}

type UserGetResponse = Response[UserGetOutputDTO]

func (u *UserGetUC) Execute(dto UserGetInputDTO) (UserGetOutputDTO, error) {
	form, err := u.UserQueries.Get(dto.ID)
	if err != nil && err.Error() == queries.UserNotActive.Error() {
		err = nil
	}
	return UserGetOutputDTO{
		Model: form,
	}, err
}

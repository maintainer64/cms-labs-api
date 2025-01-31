package auth

import (
	"errors"

	"gitlab.com/a10869/api-modules/backend/app/usecases/response"

	"gitlab.com/a10869/api-modules/backend/app/queries"
	"golang.org/x/crypto/bcrypt"
)

type UserPasswordRecoverUC struct {
	UserPasswordQueries *queries.UserPasswordQueries
	User                *SSOTokenPublicData
}

type UserPasswordChangeInputDTO struct {
	OldPassword   string `json:"old_password"`
	NewPassword   string `json:"new_password"`
	AgainPassword string `json:"again_password"`
}

type UserPasswordChangeOutputDTO struct {
	Id uint `json:"id"`
}

type UserPasswordRecoverResponse = response.Response[UserPasswordChangeOutputDTO]

func (u *UserPasswordRecoverUC) SetContext(user *SSOTokenPublicData) *UserPasswordRecoverUC {
	u.User = user
	return u
}

func (u *UserPasswordRecoverUC) Execute(dto UserPasswordChangeInputDTO) (UserPasswordChangeOutputDTO, error) {
	wrongPassword := errors.New("wrong password")
	response := UserPasswordChangeOutputDTO{
		Id: u.User.Id,
	}
	if dto.NewPassword != dto.AgainPassword {
		return response, wrongPassword
	}
	creds, err := u.UserPasswordQueries.Get(u.User.Id)
	if err != nil {
		return response, wrongPassword
	}
	err = bcrypt.CompareHashAndPassword([]byte(creds.HashPassword), []byte(dto.OldPassword))
	if err != nil {
		return response, wrongPassword
	}
	newHashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(dto.NewPassword),
		14,
	)
	if err != nil {
		return response, wrongPassword
	}
	creds.HashPassword = string(newHashPassword)
	_ = u.UserPasswordQueries.Upsert(&creds)
	response.Id = u.User.Id
	return response, err
}

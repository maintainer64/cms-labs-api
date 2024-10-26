package usecases

import (
	"github.com/ory/go-convenience/stringsx"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"time"
)

type UserEditUC struct {
	UserQueries *queries.UserQueries
}

type UserEditInputDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	LTIUserID string `json:"lti_user_id"`
	UserRole  string `json:"user_role"`
	GroupName string `json:"group_name"`
	IsActive  bool   `json:"is_active"`
}

type UserEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type UserEditResponse = Response[UserEditOutputDTO]

func (u *UserEditUC) Execute(dto UserEditInputDTO) (UserEditOutputDTO, error) {
	now := time.Now().UTC()
	userFromDB, err := u.UserQueries.Get(dto.ID)
	entity := &models.User{}
	entity.ID = dto.ID
	entity.Email = stringsx.Coalesce(dto.Email, userFromDB.Email)
	entity.Name = stringsx.Coalesce(dto.Name, userFromDB.Name)
	entity.LTIUserID = stringsx.Coalesce(dto.LTIUserID, userFromDB.LTIUserID)
	entity.UserRole = stringsx.Coalesce(
		models.UsersRoleValidate(dto.UserRole),
		models.UsersRoleValidate(userFromDB.UserRole),
	)
	entity.GroupName = stringsx.Coalesce(dto.GroupName, userFromDB.GroupName)
	entity.DeletedAt = userFromDB.DeletedAt
	entity.LastLaunchID = userFromDB.LastLaunchID
	if dto.IsActive {
		entity.DeletedAt = nil
	}
	if !dto.IsActive && entity.IsActive() {
		entity.DeletedAt = &now
	}
	err = u.UserQueries.Upsert(entity)
	return UserEditOutputDTO{ID: entity.ID}, err
}

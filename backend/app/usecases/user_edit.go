package usecases

import (
	"time"

	"gitlab.com/a10869/api-modules/backend/app/usecases/response"

	"github.com/ory/go-convenience/stringsx"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type UserEditUC struct {
	UserQueries *queries.UserQueries
	RoleQueries *queries.RoleQueries
}

type UserEditInputDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	LTIUserID string `json:"lti_user_id"`
	Roles     []uint `json:"roles"`
	GroupName string `json:"group_name"`
	IsActive  bool   `json:"is_active"`
}

type UserEditOutputDTO struct {
	ID uint `json:"id" required:"true"`
}

type UserEditResponse = response.Response[UserEditOutputDTO]

func (u *UserEditUC) Execute(dto UserEditInputDTO) (UserEditOutputDTO, error) {
	now := time.Now().UTC()
	userFromDB, _ := u.UserQueries.Get(dto.ID)
	entity := &models.User{}
	entity.ID = dto.ID
	entity.Email = stringsx.Coalesce(dto.Email, userFromDB.Email)
	entity.Name = stringsx.Coalesce(dto.Name, userFromDB.Name)
	entity.LTIUserID = stringsx.Coalesce(dto.LTIUserID, userFromDB.LTIUserID)
	entity.GroupName = stringsx.Coalesce(dto.GroupName, userFromDB.GroupName)
	entity.DeletedAt = userFromDB.DeletedAt
	entity.LastLaunchID = userFromDB.LastLaunchID
	if dto.IsActive {
		entity.DeletedAt = nil
	}
	if !dto.IsActive && entity.IsActive() {
		entity.DeletedAt = &now
	}
	err := u.UserQueries.Upsert(entity)
	if err != nil {
		return UserEditOutputDTO{}, err
	}
	err = u.RoleQueries.SetByUserId(entity.ID, dto.Roles)
	return UserEditOutputDTO{ID: entity.ID}, err
}

package tasks

import (
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/shared/logs"
	"golang.org/x/crypto/bcrypt"
)

var (
	log = logs.NewZeroLogger("tasks")
)

type StartupFiberUC struct {
	UserQueries         *queries.UserQueries
	UserPasswordQueries *queries.UserPasswordQueries
}

const UserDefaultEmail = "admin@admin.com"
const UserDefaultPassword = "admin"
const UserDefaultName = "admin"

func (u *StartupFiberUC) userDefaultCreate() error {
	entity, _ := u.UserQueries.GetByEmail(UserDefaultEmail)
	if entity.Email != UserDefaultEmail {
		entity.Name = UserDefaultName
		entity.Email = UserDefaultEmail
		entity.UserRole = models.UsersRoleAdmin
		_ = u.UserQueries.Upsert(&entity)
	}
	creds, _ := u.UserPasswordQueries.Get(entity.ID)
	if creds.UserID == entity.ID {
		return nil

	}
	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(UserDefaultPassword),
		14,
	)
	if err != nil {
		return err
	}
	err = u.UserPasswordQueries.Upsert(&models.UserPassword{
		UserID:       entity.ID,
		HashPassword: string(hashPassword),
	})
	log.Info().Msg("User default create")
	return err
}

func (u *StartupFiberUC) Startup() error {
	_ = u.userDefaultCreate()
	return nil
}

package tasks

import (
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"golang.org/x/crypto/bcrypt"
)

type StartupFiberUC struct {
	UserQueries         *queries.UserQueries
	UserPasswordQueries *queries.UserPasswordQueries
	RoleQueries         *queries.RoleQueries
	*zerolog.Logger
}

const UserDefaultEmail = "admin@admin.com"
const UserDefaultPassword = "admin"
const UserDefaultName = "admin"

func (u *StartupFiberUC) userDefaultCreate() (uint, error) {
	entity, _ := u.UserQueries.GetByEmail(UserDefaultEmail)
	if entity.Email != UserDefaultEmail {
		entity.Name = UserDefaultName
		entity.Email = UserDefaultEmail
		_ = u.UserQueries.Upsert(&entity)
	}
	creds, _ := u.UserPasswordQueries.Get(entity.ID)
	if creds.UserID == entity.ID {
		return entity.ID, nil
	}
	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(UserDefaultPassword),
		14,
	)
	if err != nil {
		return entity.ID, err
	}
	err = u.UserPasswordQueries.Upsert(&models.UserPassword{
		UserID:       entity.ID,
		HashPassword: string(hashPassword),
	})
	u.Logger.Info().Msg("User default created")
	return entity.ID, err
}

func (u *StartupFiberUC) createRoles() []uint {
	student := &models.Role{}
	student.Code = models.UsersRoleStudent
	student.Name = "Student"

	_ = u.RoleQueries.Upsert(student)

	instructor := &models.Role{}
	instructor.Code = models.UsersRoleInstructor
	instructor.Name = "Instructor"

	_ = u.RoleQueries.Upsert(instructor)

	admin := &models.Role{}
	admin.Code = models.UsersRoleAdmin
	admin.Name = "Admin"

	_ = u.RoleQueries.Upsert(admin)

	u.Logger.Info().Msg("Roles default created")
	return []uint{student.ID, instructor.ID, admin.ID}
}

func (u *StartupFiberUC) setUserRoles(userId uint, rolesIds []uint) error {
	err := u.RoleQueries.SetByUserId(userId, rolesIds)
	u.Logger.Info().Msg("Roles admin user set")
	return err
}

func (u *StartupFiberUC) Startup() error {
	userId, _ := u.userDefaultCreate()
	rolesIds := u.createRoles()
	_ = u.setUserRoles(userId, rolesIds)
	return nil
}

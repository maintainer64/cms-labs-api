package queries

import (
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/models"
	"gorm.io/gorm"
)

const (
	UserRoleStudent            = "student"
	UserRoleStudentDescription = "LTI Role integration"
)

type UserRoleQueries struct {
	*gorm.DB
}

func (q *UserRoleQueries) GetOrCreateDefault() (models.UserRole, error) {
	entityDB := models.UserRole{}
	q.Where("name = ?", UserRoleStudent).Find(&entityDB)
	if entityDB.UserRoleID != 0 {
		return entityDB, nil
	}
	entityDB.UserRoleID = 0
	entityDB.UserRoleName = UserRoleStudent
	entityDB.UserRoleWorkspace = "/"
	entityDB.UserRoleNote = UserRoleStudentDescription
	entityDB.UserRoleRAM = 0
	entityDB.UserRoleCPU = 0
	entityDB.UserRoleHDD = 0
	result := q.Create(entityDB)
	return entityDB, result.Error
}

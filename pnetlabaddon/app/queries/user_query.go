package queries

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/models"
	"gitlab.com/a10869/api-modules/shared/utils"
	"gorm.io/gorm"
)

var (
	UserNotFoundError = errors.New("User not found")
)

type UserQueries struct {
	DB *gorm.DB
}

func (q *UserQueries) Get(id int) (models.User, error) {
	entityDB := models.User{}
	q.DB.Where("pod = ?", id).Find(&entityDB)
	if entityDB.Pod != 0 {
		return entityDB, nil
	}
	return entityDB, utils.FiberValidationException{
		Status:    fiber.StatusNotFound,
		Exception: UserNotFoundError,
	}
}

func (q *UserQueries) GetOrCreateByEmail(user models.User) (models.User, error) {
	entityDB := models.User{}
	q.DB.Where("email = ?", user.Email).Find(&entityDB)
	if entityDB.Pod != 0 {
		return entityDB, nil
	}
	result := q.DB.Create(&user)
	if result.Error != nil {
		return user, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: result.Error,
		}
	}
	return user, nil
}

package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/ory/go-convenience/stringsx"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/shared/utils"
	"gorm.io/gorm"
)

type UserPasswordQueries struct {
	DB     *gorm.DB
	Logger *zerolog.Logger
}

var IncorrectPassword = errors.New("Incorrect password")

func (q *UserPasswordQueries) Get(userID uint) (models.UserPassword, error) {
	var entity models.UserPassword
	q.DB.Where("`user_id` = ?", userID).Find(&entity)
	if entity.UserID != userID {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: IncorrectPassword,
		}
	}
	return entity, nil
}

func (q *UserPasswordQueries) Upsert(entity *models.UserPassword) error {
	if entity == nil {
		return nil
	}
	entityDB := models.UserPassword{}
	q.DB.Where("`user_id` = ?", entity.UserID).Find(&entityDB)
	if entityDB.UserID == entity.UserID {
		// Update
		q.Logger.Info().Msg(fmt.Sprintf("UserPasswordQueries: entity update user_id=%+v", entityDB.UserID))
		entity.HashPassword = stringsx.Coalesce(entity.HashPassword, entityDB.HashPassword)
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		result := q.DB.Save(&entity)
		return result.Error
	} else {
		// Create
		q.Logger.Info().Msg(fmt.Sprintf("UserPasswordQueries: entity create user_id=%+v", entity.UserID))
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.DB.Create(entity)
		return result.Error
	}
}

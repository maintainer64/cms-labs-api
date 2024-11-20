package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ory/go-convenience/stringsx"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/gorm"
)

type UserTokenQueries struct {
	*gorm.DB
}

var RefreshTokenNotFoundError = errors.New("Refresh token not found")

func (q *UserTokenQueries) Get(userID uint) (models.UserToken, error) {
	var entity models.UserToken
	q.Where("`user_id` = ?", userID).Find(&entity)
	if entity.UserID != userID {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: RefreshTokenNotFoundError,
		}
	}
	return entity, nil
}

func (q *UserTokenQueries) GetByRefreshToken(refreshToken string) (models.UserToken, error) {
	var entity models.UserToken
	q.Where("`refresh_token` = ?", refreshToken).Find(&entity)
	if entity.RefreshToken != refreshToken {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: RefreshTokenNotFoundError,
		}
	}
	return entity, nil
}

func (q *UserTokenQueries) Upsert(entity *models.UserToken) error {
	if entity == nil {
		return nil
	}
	entityDB := models.UserToken{}
	q.Where("`user_id` = ?", entity.UserID).Find(&entityDB)
	if entityDB.UserID == entity.UserID {
		// Update
		log.Info().Msg(fmt.Sprintf("UserTokenQueries: entity update user_id=%+v", entityDB.UserID))
		entity.HashPassword = stringsx.Coalesce(entity.HashPassword, entityDB.HashPassword)
		entity.RefreshToken = stringsx.Coalesce(entity.RefreshToken, entityDB.RefreshToken)
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		log.Info().Msg(fmt.Sprintf("UserTokenQueries: entity create user_id=%+v", entity.UserID))
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.Create(entity)
		return result.Error
	}
}

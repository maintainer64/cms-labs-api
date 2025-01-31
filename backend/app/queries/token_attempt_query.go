package queries

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/gorm"
)

type TokenAttemptQueries struct {
	*gorm.DB
}

func (q *TokenAttemptQueries) GetByToken(token string) (models.TokenAttempt, error) {
	entity := models.TokenAttempt{}
	q.Where("token = ?", token).Limit(1).Find(&entity)
	exception := utils.FiberValidationException{
		Status:    fiber.StatusNotFound,
		Exception: errors.New("TokenAttempt has not found"),
	}
	if entity.ID == 0 {
		return entity, exception
	}
	return entity, nil
}

func (q *TokenAttemptQueries) GetByAuthCode(code string) (models.TokenAttempt, error) {
	entity := models.TokenAttempt{}
	q.Where("authorization_code = ?", code).Limit(1).Find(&entity)
	exception := utils.FiberValidationException{
		Status:    fiber.StatusNotFound,
		Exception: errors.New("TokenAttempt code not found"),
	}
	if entity.ID == 0 {
		return entity, exception
	}
	return entity, nil
}

func (q *TokenAttemptQueries) Upsert(entity *models.TokenAttempt) error {
	err := q.DeleteByParams(entity.UserID, entity.ServerID)
	if err != nil {
		return nil
	}
	result := q.Create(entity)
	log.Debug().Msg(fmt.Sprintf("TokenAttemptQueries: entity create: %+v", entity))
	log.Info().Msg(fmt.Sprintf("TokenAttemptQueries: entity create user_id=%+v", entity.UserID))
	return result.Error
}

func (q *TokenAttemptQueries) GetByParams(userID uint, serverID uint) (models.TokenAttempt, error) {
	entity := models.TokenAttempt{}
	q.Where("user_id = ?", userID).Where("server_id = ?", serverID).
		Limit(1).Find(&entity)
	exception := utils.FiberValidationException{
		Status:    fiber.StatusNotFound,
		Exception: errors.New("TokenAttempt server or user not found"),
	}
	if entity.ID == 0 {
		return entity, exception
	}
	return entity, nil
}

func (q *TokenAttemptQueries) DeleteByParams(userID uint, serverID uint) error {
	result := q.Where("user_id = ?", userID).Where("server_id = ?", serverID).Delete(
		&models.TokenAttempt{},
	)
	if result.Error != nil {
		return result.Error
	}
	log.Info().Msg(
		fmt.Sprintf("TokenAttemptQueries: delete token by user_id=%+v server_id=%+v", userID, serverID),
	)
	return nil
}

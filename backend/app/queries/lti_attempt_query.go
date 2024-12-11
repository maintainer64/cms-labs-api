package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/gorm"
)

type LTIAttemptQueries struct {
	*gorm.DB
}

func (q *LTIAttemptQueries) Get(id uint) (models.LTIAttempt, error) {
	var entity models.LTIAttempt
	result := q.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("LTIAttempt not found"),
		}
	}
	return entity, result.Error
}

func (q *LTIAttemptQueries) Upsert(entity *models.LTIAttempt) error {
	if entity == nil {
		return nil
	}
	entityDB := models.LTIAttempt{}
	q.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		log.Debug().Msg(fmt.Sprintf("LTIAttemptQueries: entity update: %+v", entityDB))
		log.Info().Msg(fmt.Sprintf("LTIAttemptQueries: entity update user_id=%+v", entityDB.UserID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		entity.UserID = entityDB.UserID
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		log.Debug().Msg(fmt.Sprintf("LTIAttemptQueries: entity create: %+v", entity))
		log.Info().Msg(fmt.Sprintf("LTIAttemptQueries: entity create user_id=%+v", entity.UserID))
		entity.ID = 0
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.Create(entity)
		return result.Error
	}
}

func (q *LTIAttemptQueries) GetActiveByUserId(userId uint) (models.LTIAttempt, error) {
	var entity models.LTIAttempt
	result := q.Model(&entity).Where(
		"user_id = ?", userId,
	).Where("UTC_TIMESTAMP() < expired_at").Limit(
		1,
	).Offset(0).Order(`created_at desc`).Find(&entity)
	return entity, result.Error
}

func (q *LTIAttemptQueries) GetByRoomNumber(roomNumber *int) ([]models.LTIAttemptListItem, error) {
	var entities []models.LTIAttemptListItem
	if roomNumber == nil {
		return entities, nil
	}
	result := q.Model(&entities).Where(
		"room_number = ?",
		*roomNumber,
	).Where(
		"UTC_TIMESTAMP() < expired_at",
	).Limit(MaxLimitCount).Offset(0).Find(&entities)
	count := result.RowsAffected
	log.Info().Msg(fmt.Sprintf("LTIAttemptQueries GetByRoomNumber: count %+v", count))
	return entities, result.Error
}

func (q *LTIAttemptQueries) List(
	search string,
	limit int,
	offset int,
) ([]models.LTIAttemptListItem, int64, error) {
	var entities []models.LTIAttemptListItem
	result := q.listFilter(
		search,
		q.Limit(MaxLimitCount).Offset(0),
	).Find(&entities)
	count := result.RowsAffected
	log.Debug().Msg(fmt.Sprintf("LTIAttemptQueries list: count %+v", count))
	result = q.listFilter(
		search,
		q.Limit(limit).Offset(offset),
	).Find(&entities)
	log.Debug().Msg(fmt.Sprintf("LTIAttemptQueries list: entities %+v", entities))
	return entities, count, result.Error
}

func (q *LTIAttemptQueries) listFilter(search string, tx *gorm.DB) *gorm.DB {
	tx = tx.Model(&models.LTIAttempt{})
	tx = tx.Order(`created_at desc`)
	if search == "" {
		return tx
	}
	tx = tx.Or("id = ?", search)
	return tx
}

func (q *LTIAttemptQueries) Delete(id uint) error {
	_ = q.Where("id = ?", id).Delete(&models.LTIAttempt{})
	log.Debug().Msg(fmt.Sprintf("LTIAttemptQueries: delete entity by id: %+v", id))
	return nil
}

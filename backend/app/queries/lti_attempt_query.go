package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/shared/utils"
	"gorm.io/gorm"
)

type LTIAttemptQueries struct {
	*gorm.DB
	*zerolog.Logger
}

const MaxLimitCount = 5000

func (q *LTIAttemptQueries) tableName(object interface{}) string {
	stmt := &gorm.Statement{DB: q.DB}
	_ = stmt.Parse(object)
	return stmt.Schema.Table
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
		q.Logger.Debug().Msg(fmt.Sprintf("LTIAttemptQueries: entity update: %+v", entityDB))
		q.Logger.Info().Msg(fmt.Sprintf("LTIAttemptQueries: entity update user_id=%+v", entityDB.UserID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		entity.UserID = entityDB.UserID
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		q.Logger.Debug().Msg(fmt.Sprintf("LTIAttemptQueries: entity create: %+v", entity))
		q.Logger.Info().Msg(fmt.Sprintf("LTIAttemptQueries: entity create user_id=%+v", entity.UserID))
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

func (q *LTIAttemptQueries) GetByRoomNumber(roomNumber *int64) ([]models.LTIAttempt, error) {
	var entities []models.LTIAttempt
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
	q.Logger.Info().Msg(fmt.Sprintf("LTIAttemptQueries GetByRoomNumber: count %+v", count))
	return entities, result.Error
}

func (q *LTIAttemptQueries) List(
	userIds []uint,
	limit int,
	offset int,
) ([]models.LTIAttemptListItem, error) {
	var entities []models.LTIAttemptListItem
	result := q.listFilter(
		userIds,
		q.Limit(limit).Offset(offset),
	).Find(&entities)
	return entities, result.Error
}

func (q *LTIAttemptQueries) listFilter(userIds []uint, tx *gorm.DB) *gorm.DB {
	tx = tx.Table(
		q.tableName(&models.LTIAttempt{}) + " AS lti_attempts",
	).Select(
		"lti_attempts.id, lti_attempts.user_id, lti_attempts.pnet_server_id, lti_attempts.lti_routing_id, lti_attempts.expired_at, " +
			"users.email as user_email, users.name as user_name, pnet_servers.name as pnet_server_name, lti_routings.name as lti_routing_name",
	).Joins(
		"join " + q.tableName(&models.User{}) + " users on users.id = lti_attempts.user_id",
	).Joins(
		"join " + q.tableName(&models.PNETServer{}) + " pnet_servers on pnet_servers.id = lti_attempts.pnet_server_id",
	).Joins(
		"join " + q.tableName(&models.LTIRouting{}) + " lti_routings on lti_routings.id = lti_attempts.lti_routing_id",
	)
	tx = tx.Order(`lti_attempts.created_at desc`)
	if len(userIds) > 0 {
		tx = tx.Where("lti_attempts.user_id IN (?)", userIds)
	}
	return tx
}

func (q *LTIAttemptQueries) Delete(id uint) error {
	_ = q.Where("id = ?", id).Delete(&models.LTIAttempt{})
	q.Logger.Debug().Msg(fmt.Sprintf("LTIAttemptQueries: delete entity by id: %+v", id))
	return nil
}

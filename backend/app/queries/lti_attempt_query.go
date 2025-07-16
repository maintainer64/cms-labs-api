package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ory/go-convenience/stringsx"

	"github.com/rs/zerolog"

	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/shared/utils"
	"gorm.io/gorm"
)

type LTIAttemptQueries struct {
	DB     *gorm.DB
	Logger *zerolog.Logger
}

const MaxLimitCount = 5000

func (q *LTIAttemptQueries) tableName(object interface{}) string {
	stmt := &gorm.Statement{DB: q.DB}
	_ = stmt.Parse(object)
	return stmt.Schema.Table
}

func (q *LTIAttemptQueries) Get(id uint) (models.LTIAttempt, error) {
	var entity models.LTIAttempt
	result := q.DB.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("LTIAttempt not found"),
		}
	}
	return entity, result.Error
}

func (q *LTIAttemptQueries) GetByAttemptID(attemptID string) (models.LTIAttempt, error) {
	var entity models.LTIAttempt
	result := q.DB.Where("attempt_id = ?", attemptID).Find(&entity)
	if entity.ID == 0 {
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
	q.DB.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		q.Logger.Debug().Msg(fmt.Sprintf("LTIAttemptQueries: entity update: %+v", entityDB))
		q.Logger.Info().Msg(fmt.Sprintf("LTIAttemptQueries: entity update user_id=%+v", entityDB.UserID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.AttemptID = stringsx.Coalesce(
			entityDB.AttemptID,
			uuid.New().String(),
		)
		entity.UpdatedAt = time.Now().UTC()
		entity.UserID = entityDB.UserID
		result := q.DB.Save(&entity)
		return result.Error
	} else {
		// Create
		q.Logger.Debug().Msg(fmt.Sprintf("LTIAttemptQueries: entity create: %+v", entity))
		q.Logger.Info().Msg(fmt.Sprintf("LTIAttemptQueries: entity create user_id=%+v", entity.UserID))
		entity.ID = 0
		entity.AttemptID = stringsx.Coalesce(
			entity.AttemptID,
			uuid.New().String(),
		)
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.DB.Create(entity)
		return result.Error
	}
}

func (q *LTIAttemptQueries) GetActiveByUserId(userId uint, routeId uint) (models.LTIAttempt, error) {
	var entity models.LTIAttempt
	result := q.DB.Model(&entity).Where(
		"user_id = ? AND lti_routing_id = ?",
		userId,
		routeId,
	).Where("UTC_TIMESTAMP() < expired_at").Limit(
		1,
	).Offset(0).Order(`created_at desc`).Find(&entity)
	return entity, result.Error
}

func (q *LTIAttemptQueries) GetByRoomID(roomID uint) ([]models.LTIAttempt, error) {
	var entities []models.LTIAttempt
	if roomID == 0 {
		return entities, nil
	}
	result := q.DB.Model(&entities).Where(
		"room_id = ?",
		roomID,
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
		q.DB.Limit(limit).Offset(offset),
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

// AllocatedServerLock получает advisory lock по roomID или attemptID и возвращает функцию для освобождения
func (q *LTIAttemptQueries) AllocatedServerLock(attemptID uint, roomID *uint) (func(), error) {
	// Определяем ID для блокировки
	var lockID uint64
	if roomID != nil && *roomID != 0 {
		lockID = uint64(*roomID)
	} else {
		lockID = uint64(attemptID)
	}

	// Пытаемся получить блокировку с таймаутом 10 секунд
	var result int
	err := q.DB.Raw("SELECT GET_LOCK(?, 10)", lockID).Scan(&result).Error
	if err != nil {
		q.Logger.Info().Msg(fmt.Sprintf("LTIAttemptQueries: failed to acquire advisory lock: %v", err))
		return nil, fmt.Errorf("failed to acquire advisory lock: %v", err)
	}
	if result != 1 {
		q.Logger.Info().Msg(fmt.Sprintf("LTIAttemptQueries: could not acquire advisory lock for ID %d", lockID))
		return nil, fmt.Errorf("could not acquire advisory lock for ID %d", lockID)
	}

	// Функция для освобождения блокировки
	unlockFn := func() {
		q.DB.Exec("SELECT RELEASE_LOCK(?)", lockID)
	}

	return unlockFn, nil
}

// AllocatedServer закрепляет сервер за attempt(ами)
// Если roomID != nil и != 0 - обновляет все attempts с этим roomID
// Иначе обновляет только указанный attempt
func (q *LTIAttemptQueries) AllocatedServer(attemptID uint, roomID *uint, pnetServerID uint) error {
	q.Logger.Info().Msg(
		fmt.Sprintf(
			"LTIAttemptQueries: set attemptID: %d by pnetServerID: %d",
			attemptID,
			pnetServerID,
		),
	)
	updateData := map[string]interface{}{
		"pnet_server_id": pnetServerID,
		"updated_at":     time.Now().UTC(),
	}

	if roomID != nil && *roomID != 0 {
		// Обновляем все attempts с этим roomID
		err := q.DB.Model(&models.LTIAttempt{}).
			Where("room_id = ?", *roomID).
			Updates(updateData).Error
		if err != nil {
			q.Logger.Info().Msg(fmt.Sprintf("failed to update attempts by roomID %d: %v", *roomID, err))
			return fmt.Errorf("failed to update attempts by roomID %d: %v", *roomID, err)
		}
		return nil
	}

	// Обновляем только указанный attempt
	err := q.DB.Model(&models.LTIAttempt{}).
		Where("id = ?", attemptID).
		Updates(updateData).Error
	if err != nil {
		q.Logger.Info().Msg(fmt.Sprintf("failed to update attempt %d: %v", attemptID, err))
		return fmt.Errorf("failed to update attempt %d: %v", attemptID, err)
	}
	return nil
}

func (q *LTIAttemptQueries) Delete(id uint) error {
	err := q.DB.Where("id = ?", id).Delete(&models.LTIAttempt{}).Error
	q.Logger.Debug().Msg(fmt.Sprintf("LTIAttemptQueries: delete entity by id: %+v", id))
	return err
}

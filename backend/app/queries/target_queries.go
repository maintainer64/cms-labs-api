package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
)

type TargetQueries struct {
	DB     *gorm.DB
	Logger *zerolog.Logger
}

var TargetNotFoundError = jsonrpc.NewRpcError("target_not_found", "Target not found")

// Get возвращает полную запись Target по ID
func (q *TargetQueries) Get(id string) (models.Target, error) {
	var entity models.Target
	err := q.DB.First(&entity, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity, TargetNotFoundError
	}
	return entity, err
}

// GetByName возвращает полную запись Target по name
func (q *TargetQueries) GetByName(name string) (models.Target, error) {
	var entity models.Target
	err := q.DB.First(&entity, "name = ?", name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity, TargetNotFoundError
	}
	return entity, err
}

// Upsert создаёт или обновляет Target
func (q *TargetQueries) Upsert(entity *models.Target) error {
	if entity == nil {
		return nil
	}
	var existing models.Target
	err := q.DB.Where("id = ?", entity.ID).First(&existing).Error
	if err == nil {
		// Update
		q.Logger.Debug().Msg(fmt.Sprintf("TargetQueries: update %+v", existing))
		q.Logger.Info().Msg(fmt.Sprintf("TargetQueries: update id=%s", existing.ID))
		entity.CreatedAt = existing.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		return q.DB.Save(entity).Error
	} else if err == gorm.ErrRecordNotFound {
		// Create
		q.Logger.Debug().Msg(fmt.Sprintf("TargetQueries: create %+v", entity))
		q.Logger.Info().Msg(fmt.Sprintf("TargetQueries: create id=%s", entity.ID))
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		return q.DB.Create(entity).Error
	}
	return err
}

// Delete удаляет Target по ID
func (q *TargetQueries) Delete(id string) error {
	q.Logger.Debug().Msg(fmt.Sprintf("TargetQueries: delete id=%s", id))
	return q.DB.Where("id = ?", id).Delete(&models.Target{}).Error
}

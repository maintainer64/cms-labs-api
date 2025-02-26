package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/shared/utils"
	"gorm.io/gorm"
)

type UNLFileQueries struct {
	*gorm.DB
	*zerolog.Logger
}

func (q *UNLFileQueries) Get(id uint) (models.UNLFile, error) {
	var entity models.UNLFile
	result := q.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("UNLFile not found"),
		}
	}
	return entity, result.Error
}

func (q *UNLFileQueries) Upsert(entity *models.UNLFile) error {
	if entity == nil {
		return nil
	}
	entityDB := models.UNLFile{}
	q.Where("path = ?", entity.Path).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		q.Logger.Debug().Msg(fmt.Sprintf("UNLFileQueries: entity update: %+v", entityDB))
		q.Logger.Info().Msg(fmt.Sprintf("UNLFileQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		result := q.Save(&entity)
		return result.Error
	}
	// Create
	q.Logger.Debug().Msg(fmt.Sprintf("UNLFileQueries: entity create: %+v", entity))
	q.Logger.Info().Msg(fmt.Sprintf("UNLFileQueries: entity create path=%+v", entity.Path))
	entity.ID = 0
	entity.CreatedAt = time.Now().UTC()
	entity.UpdatedAt = time.Now().UTC()
	result := q.Create(entity)
	return result.Error
}

func (q *UNLFileQueries) List(
	search string,
	typeFile []string,
	limit int,
	offset int,
) ([]models.UNLFileListItem, error) {
	var entities []models.UNLFileListItem
	result := q.listFilter(
		search,
		typeFile,
		q.Limit(limit).Offset(offset),
	).Find(&entities)
	q.Logger.Debug().Msg(fmt.Sprintf("UNLFileQueries list: entities %+v", entities))
	return entities, result.Error
}

func (q *UNLFileQueries) listFilter(search string, typeFile []string, tx *gorm.DB) *gorm.DB {
	tx = unlFileNoBlobSelect(tx)
	tx = tx.Order(`created_at desc`)
	// Фильтр по typeFile, если он не пустой
	if len(typeFile) > 0 {
		tx = tx.Where("type IN ?", typeFile)
	}

	// Фильтр по deleted_at, чтобы исключить удаленные записи
	tx = tx.Where("deleted_at IS NULL")

	// Если search не пустой, добавляем условия для поиска по id или path
	if search != "" {
		tx = tx.Where("id = ? OR path LIKE ?", search, "%"+search+"%")
	}

	return tx
}

func (q *UNLFileQueries) SoftDeleteNotSyncedId(syncedId string) error {
	entityDB := models.UNLFile{}
	result := q.Where("synced_id != ?", syncedId).Delete(&entityDB)
	return result.Error
}

func unlFileNoBlobSelect(tx *gorm.DB) *gorm.DB {
	tx = tx.Model(&models.UNLFile{}).Select(
		"id",
		"created_at",
		"updated_at",
		"path",
		"type",
		"deleted_at",
	)
	return tx
}

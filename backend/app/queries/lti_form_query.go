package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/gorm"
)

type LTIFromQueries struct {
	*gorm.DB
}

func (q *LTIFromQueries) Get(id uint) (models.LTIForm, error) {
	var entity models.LTIForm
	result := q.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("LTIForm not found"),
		}
	}
	return entity, result.Error
}

func (q *LTIFromQueries) Upsert(entity *models.LTIForm) error {
	if entity == nil {
		return nil
	}
	entityDB := models.LTIForm{}
	q.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		log.Debug(fmt.Sprintf("LTIFromQueries: entity update: %+v", entityDB))
		log.Info(fmt.Sprintf("LTIFromQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		entity.Version = entityDB.Version
		entity.LTISecret = entityDB.LTISecret
		entity.LTIKey = entityDB.LTIKey
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		log.Debug(fmt.Sprintf("LTIFromQueries: entity create: %+v", entityDB))
		log.Info(fmt.Sprintf("LTIFromQueries: entity create name=%+v", entityDB.Name))
		entity.ID = 0
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		entity.LTIKey = uuid.New().String()
		entity.LTISecret = fmt.Sprintf("%s%s", uuid.New().String(), uuid.New().String())
		result := q.Create(entity)
		return result.Error
	}
}

func (q *LTIFromQueries) List(limit int, offset int) ([]models.LTIFormListItem, error) {
	var entities []models.LTIFormListItem
	result := q.Limit(limit).Offset(offset).Order(`created_at desc`).Find(&entities)
	log.Debug(fmt.Sprintf("LTIFromQueries: entities %+v", entities))
	return entities, result.Error
}

func (q *LTIFromQueries) Delete(id uint) error {
	_ = q.Where("id = ?", id).Delete(&models.LTIForm{})
	log.Debug(fmt.Sprintf("LTIFromQueries: delete entity by id: %+v", id))
	return nil
}

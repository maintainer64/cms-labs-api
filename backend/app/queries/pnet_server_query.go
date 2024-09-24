package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/gorm"
)

type PNETServerQueries struct {
	*gorm.DB
}

func (q *PNETServerQueries) Get(id uint) (models.PNETServer, error) {
	var entity models.PNETServer
	result := q.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("PNETServer not found"),
		}
	}
	return entity, result.Error
}

func (q *PNETServerQueries) Upsert(entity *models.PNETServer) error {
	if entity == nil {
		return nil
	}
	entityDB := models.PNETServer{}
	q.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		log.Debug(fmt.Sprintf("PNETServerQueries: entity update: %+v", entityDB))
		log.Info(fmt.Sprintf("PNETServerQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		entity.LastOnlineStatus = entityDB.LastOnlineStatus
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		log.Debug(fmt.Sprintf("PNETServerQueries: entity create: %+v", entityDB))
		log.Info(fmt.Sprintf("PNETServerQueries: entity create name=%+v", entityDB.Name))
		entity.ID = 0
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.Create(entity)
		return result.Error
	}
}

func (q *PNETServerQueries) List(limit int, offset int) ([]models.PNETServerListItem, error) {
	var entities []models.PNETServerListItem
	result := q.Limit(limit).Offset(offset).Order(`created_at desc`).Find(&entities)
	log.Debug(fmt.Sprintf("PNETServerQueries: entities %+v", entities))
	return entities, result.Error
}

func (q *PNETServerQueries) Delete(id uint) error {
	_ = q.Where("id = ?", id).Delete(&models.PNETServer{})
	log.Debug(fmt.Sprintf("PNETServerQueries: delete entity by id: %+v", id))
	return nil
}

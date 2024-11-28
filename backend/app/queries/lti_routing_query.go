package queries

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type LTIRoutingQueries struct {
	*gorm.DB
}

func (q *LTIRoutingQueries) Get(id uint) (models.LTIRouting, error) {
	var entity models.LTIRouting
	result := q.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("LTIRouting not found"),
		}
	}
	return entity, result.Error
}

func (q *LTIRoutingQueries) Upsert(entity *models.LTIRouting) error {
	if entity == nil {
		return nil
	}
	entityDB := models.LTIRouting{}
	q.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		log.Debug().Msg(fmt.Sprintf("LTIRoutingQueries: entity update: %+v", entityDB))
		log.Info().Msg(fmt.Sprintf("LTIRoutingQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.Name = entityDB.Name
		entity.LTITitle = entityDB.LTITitle
		entity.LTIDescription = entityDB.LTIDescription
		entity.LTITaskID = entityDB.LTITaskID
		entity.LTIParamsTask = entityDB.LTIParamsTask
		entity.Collaboration = entityDB.Collaboration
		entity.PNETLabsPath = entityDB.PNETLabsPath
		entity.PNETTestPath = entityDB.PNETTestPath
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		log.Debug().Msg(fmt.Sprintf("LTIRoutingQueries: entity create: %+v", entity))
		log.Info().Msg(fmt.Sprintf("LTIRoutingQueries: entity create name=%+v", entity.Name))
		entity.ID = 0
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.Create(entity)
		return result.Error
	}
}

func (q *LTIRoutingQueries) List(
	search string,
	limit int,
	offset int,
) ([]models.LTIRoutingListItem, int64, error) {
	var entities []models.LTIRoutingListItem
	result := q.listFilter(
		search,
		q.Limit(MaxLimitCount).Offset(0),
	).Find(&entities)
	count := result.RowsAffected
	log.Debug().Msg(fmt.Sprintf("LTIRoutingQueries list: count %+v", count))
	result = q.listFilter(
		search,
		q.Limit(limit).Offset(offset),
	).Find(&entities)
	log.Debug().Msg(fmt.Sprintf("LTIRoutingQueries list: entities %+v", entities))
	return entities, count, result.Error
}

func (q *LTIRoutingQueries) listFilter(search string, tx *gorm.DB) *gorm.DB {
	tx = tx.Order(`created_at desc`)
	if search == "" {
		return tx
	}
	tx = tx.Or("id = ?", search)
	return tx
}

func (q *LTIRoutingQueries) Delete(id uint) error {
	_ = q.Where("id = ?", id).Delete(&models.LTIRouting{})
	log.Info().Msg(fmt.Sprintf("LTIRoutingQueries: delete entity by id: %+v", id))
	return nil
}

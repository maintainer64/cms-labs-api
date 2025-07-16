package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"

	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/shared/utils"
	"gorm.io/gorm"
)

type ServiceCardQueries struct {
	DB     *gorm.DB
	Logger *zerolog.Logger
}

func (q *ServiceCardQueries) Get(id uint) (models.ServiceCard, error) {
	var entity models.ServiceCard
	result := q.DB.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("ServiceCard not found"),
		}
	}
	return entity, result.Error
}

func (q *ServiceCardQueries) Upsert(entity *models.ServiceCard) error {
	if entity == nil {
		return nil
	}
	entityDB := models.ServiceCard{}
	q.DB.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		q.Logger.Debug().Msg(fmt.Sprintf("ServiceCardQueries: entity update: %+v", entityDB))
		q.Logger.Info().Msg(fmt.Sprintf("ServiceCardQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		result := q.DB.Save(&entity)
		return result.Error
	} else {
		// Create
		q.Logger.Debug().Msg(fmt.Sprintf("ServiceCardQueries: entity create: %+v", entity))
		q.Logger.Info().Msg(fmt.Sprintf("ServiceCardQueries: entity create name=%+v", entity.Name))
		entity.ID = 0
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.DB.Create(entity)
		return result.Error
	}
}

func (q *ServiceCardQueries) List(
	search string,
	limit int,
	offset int,
) ([]models.ServiceCardListItem, int64, error) {
	var entities []models.ServiceCardListItem
	result := q.listFilter(
		search,
		q.DB.Limit(MaxLimitCount).Offset(0),
	).Find(&entities)
	count := result.RowsAffected
	q.Logger.Debug().Msg(fmt.Sprintf("ServiceCardQueries list: count %+v", count))
	result = q.listFilter(
		search,
		q.DB.Limit(limit).Offset(offset),
	).Find(&entities)
	q.Logger.Debug().Msg(fmt.Sprintf("ServiceCardQueries list: entities %+v", entities))
	return entities, count, result.Error
}

func (q *ServiceCardQueries) listFilter(search string, tx *gorm.DB) *gorm.DB {
	tx = tx.Model(&models.ServiceCard{})
	tx = tx.Order("`order` desc").Order(`created_at desc`)
	if search == "" {
		return tx
	}
	tx = tx.Or("id = ?", search)
	return tx
}

func (q *ServiceCardQueries) Delete(id uint) error {
	err := q.DB.Where("id = ?", id).Delete(&models.ServiceCard{}).Error
	q.Logger.Debug().Msg(fmt.Sprintf("ServiceCardQueries: delete entity by id: %+v", id))
	return err
}

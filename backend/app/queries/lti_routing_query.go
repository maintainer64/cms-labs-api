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

type LTIRoutingQueries struct {
	*gorm.DB
	*zerolog.Logger
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
		q.Logger.Debug().Msg(fmt.Sprintf("LTIRoutingQueries: entity update: %+v", entityDB))
		q.Logger.Info().Msg(fmt.Sprintf("LTIRoutingQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		q.Logger.Debug().Msg(fmt.Sprintf("LTIRoutingQueries: entity create: %+v", entity))
		q.Logger.Info().Msg(fmt.Sprintf("LTIRoutingQueries: entity create name=%+v", entity.Name))
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
	q.Logger.Debug().Msg(fmt.Sprintf("LTIRoutingQueries list: count %+v", count))
	result = q.listFilter(
		search,
		q.Limit(limit).Offset(offset),
	).Find(&entities)
	q.Logger.Debug().Msg(fmt.Sprintf("LTIRoutingQueries list: entities %+v", entities))
	return entities, count, result.Error
}

func (q *LTIRoutingQueries) listFilter(search string, tx *gorm.DB) *gorm.DB {
	tx = tx.Model(&models.LTIRouting{})
	tx = tx.Order(`created_at desc`)
	if search == "" {
		return tx
	}
	tx = tx.Or("id = ?", search)
	return tx
}

func (q *LTIRoutingQueries) GetRelevantRouting(
	title string,
	description string,
	customParams []string,
) (models.LTIRouting, error) {
	var entity models.LTIRouting
	tx := q.Model(&models.LTIRouting{})
	tx = tx.Order(`created_at desc`).Limit(1).Offset(0)
	execute := false
	q.Logger.Debug().Msg(
		fmt.Sprintf(
			"LTIRoutingQueries: GetRelevantRouting execute by params title=%+v description=%+v customParams=%+v",
			title,
			description,
			customParams,
		),
	)
	if title != "" {
		tx = tx.Or("lti_title = ?", title)
		execute = true
	}
	if description != "" {
		tx = tx.Or("lti_description = ?", title)
		execute = true
	}
	if len(customParams) > 0 {
		tx = tx.Or("lti_params_task IN ?", customParams)
		execute = true
	}
	if !execute {
		q.Logger.Info().Msg("LTIRoutingQueries: GetRelevantRouting not execute null params")
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("LTIRouting not found"),
		}
	}
	result := tx.Scan(&entity)
	q.Logger.Info().Msg(
		fmt.Sprintf(
			"LTIRoutingQueries: GetRelevantRouting execute by params result: routingId=%+v routingName=%+v",
			entity.ID,
			entity.Name,
		),
	)
	if entity.ID == 0 {
		q.Logger.Info().Msg("LTIRoutingQueries: GetRelevantRouting default params")
		q.Model(&models.LTIRouting{}).Where("is_default = ?", true).Scan(&entity).Order(`created_at desc`).Limit(1).Offset(0)
		q.Logger.Info().Msg(
			fmt.Sprintf(
				"LTIRoutingQueries: GetRelevantRouting fetch default result: routingId=%+v routingName=%+v",
				entity.ID,
				entity.Name,
			),
		)
	}
	return entity, result.Error
}

func (q *LTIRoutingQueries) Delete(id uint) error {
	err := q.Where("id = ?", id).Delete(&models.LTIRouting{}).Error
	q.Logger.Info().Msg(fmt.Sprintf("LTIRoutingQueries: delete entity by id: %+v", id))
	return err
}

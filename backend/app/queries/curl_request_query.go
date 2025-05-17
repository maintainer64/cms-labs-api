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

type CurlRequestQueries struct {
	*gorm.DB
	*zerolog.Logger
}

type CurlRequestQueriesListDTO struct {
	Search string `json:"search"`
	Ids    []uint `json:"ids"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

func (q *CurlRequestQueries) tableName(object interface{}) string {
	stmt := &gorm.Statement{DB: q.DB}
	_ = stmt.Parse(object)
	return stmt.Schema.Table
}

func (q *CurlRequestQueries) Get(id uint) (models.CurlRequest, error) {
	var entity models.CurlRequest
	result := q.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("CurlRequest not found"),
		}
	}
	return entity, result.Error
}

func (q *CurlRequestQueries) Upsert(entity *models.CurlRequest) error {
	if entity == nil {
		return nil
	}
	entityDB := models.CurlRequest{}
	q.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		q.Logger.Debug().Msg(fmt.Sprintf("CurlRequestQueries: entity update: %+v", entityDB))
		q.Logger.Info().Msg(fmt.Sprintf("CurlRequestQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		q.Logger.Debug().Msg(fmt.Sprintf("CurlRequestQueries: entity create: %+v", entity))
		q.Logger.Info().Msg(fmt.Sprintf("CurlRequestQueries: entity create name=%+v", entity.Name))
		entity.ID = 0
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.Create(entity)
		return result.Error
	}
}

func (q *CurlRequestQueries) List(filter CurlRequestQueriesListDTO) ([]models.CurlRequestListItem, int64, error) {
	var entities []models.CurlRequestListItem
	result := q.listFilter(
		filter,
		q.Limit(MaxLimitCount).Offset(0),
	).Find(&entities)
	count := result.RowsAffected
	q.Logger.Debug().Msg(fmt.Sprintf("CurlRequestQueries list: count %+v", count))
	result = q.listFilter(
		filter,
		q.Limit(filter.Limit).Offset(filter.Offset),
	).Find(&entities)
	q.Logger.Debug().Msg(fmt.Sprintf("CurlRequestQueries list: entities %+v", entities))
	return entities, count, result.Error
}

func (q *CurlRequestQueries) listFilter(filter CurlRequestQueriesListDTO, tx *gorm.DB) *gorm.DB {
	tx = tx.Table(q.tableName(&models.CurlRequest{}) + " AS curl_requests")
	tx = tx.Order(`created_at desc`)
	if len(filter.Ids) > 0 {
		tx = tx.Where("id in (?)", filter.Ids)
	} else if filter.Search != "" {
		searchPattern := "%" + filter.Search + "%"
		tx = tx.Where("name LIKE ? OR url LIKE ?", searchPattern, searchPattern)
	}
	return tx
}

func (q *CurlRequestQueries) Delete(id uint) error {
	tx := q.Where("id = ?", id).Delete(&models.CurlRequest{})
	q.Logger.Debug().Msg(fmt.Sprintf("CurlRequestQueries: delete entity by id: %+v", id))
	return tx.Error
}

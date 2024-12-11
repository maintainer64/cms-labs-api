package main

const AppQueriesTemplate = `package queries

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"{{.Root}}/app/models"
	"{{.Root}}/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type {{.Name}}Queries struct {
	*gorm.DB
}

func (q *{{.Name}}Queries) Get(id uint) (models.{{.Name}}, error) {
	var entity models.{{.Name}}
	result := q.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("{{.Name}} not found"),
		}
	}
	return entity, result.Error
}

func (q *{{.Name}}Queries) Upsert(entity *models.{{.Name}}) error {
	if entity == nil {
		return nil
	}
	entityDB := models.{{.Name}}{}
	q.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		log.Debug().Msg(fmt.Sprintf("{{.Name}}Queries: entity update: %+v", entityDB))
		log.Info().Msg(fmt.Sprintf("{{.Name}}Queries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
        /*
            TODO: Complete code with {{.Name}} entity
        */
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		log.Debug().Msg(fmt.Sprintf("{{.Name}}Queries: entity create: %+v", entity))
		log.Info().Msg(fmt.Sprintf("{{.Name}}Queries: entity create name=%+v", entity.Name))
		entity.ID = 0
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		/*
		    Complete code with {{.Name}} entity
		*/
		result := q.Create(entity)
		return result.Error
	}
}

func (q *{{.Name}}Queries) List(
	search string,
	limit int,
	offset int,
) ([]models.{{.Name}}ListItem, int64, error) {
	var entities []models.{{.Name}}ListItem
	result := q.listFilter(
		search,
		q.Limit(MaxLimitCount).Offset(0),
	).Find(&entities)
	count := result.RowsAffected
	log.Debug().Msg(fmt.Sprintf("{{.Name}}Queries list: count %+v", count))
	result = q.listFilter(
		search,
		q.Limit(limit).Offset(offset),
	).Find(&entities)
	log.Debug().Msg(fmt.Sprintf("{{.Name}}Queries list: entities %+v", entities))
	return entities, count, result.Error
}

func (q *{{.Name}}Queries) listFilter(search string, tx *gorm.DB) *gorm.DB {
	tx = tx.Model(&models.{{.Name}}{})
	tx = tx.Order(` + "`created_at desc`" + `)
	if search == "" {
		return tx
	}
	tx = tx.Or("id = ?", search)
	return tx
}

func (q *{{.Name}}Queries) Delete(id uint) error {
	_ = q.Where("id = ?", id).Delete(&models.{{.Name}}{})
	log.Debug().Msg(fmt.Sprintf("{{.Name}}Queries: delete entity by id: %+v", id))
	return nil
}

`

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
		log.Debug().Msg(fmt.Sprintf("{{.Name}}Queries: entity create: %+v", entityDB))
		log.Info().Msg(fmt.Sprintf("{{.Name}}Queries: entity create name=%+v", entityDB.Name))
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

func (q *{{.Name}}Queries) List(limit int, offset int) ([]models.{{.Name}}ListItem, error) {
	var entities []models.{{.Name}}ListItem
	result := q.Limit(limit).Offset(offset).Order(` + "`created_at desc`" + `).Find(&entities)
	log.Debug().Msg(fmt.Sprintf("{{.Name}}Queries: entities %+v", entities))
	return entities, result.Error
}

func (q *{{.Name}}Queries) Delete(id uint) error {
	_ = q.Where("id = ?", id).Delete(&models.{{.Name}}{})
	log.Debug().Msg(fmt.Sprintf("{{.Name}}Queries: delete entity by id: %+v", id))
	return nil
}

`

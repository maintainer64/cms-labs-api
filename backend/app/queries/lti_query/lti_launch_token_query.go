package lti_query

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/shared/utils"

	"gitlab.com/a10869/api-modules/backend/app/models"
	"gorm.io/gorm"
)

type LTILaunchDataQueries struct {
	*gorm.DB
	*zerolog.Logger
}

func (q *LTILaunchDataQueries) Get(id string) (models.LTILaunchData, error) {
	var entity models.LTILaunchData
	result := q.First(&entity, "id = ?", id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("LTILaunchData not found"),
		}
	}
	return entity, result.Error
}

func (q *LTILaunchDataQueries) Upsert(entity *models.LTILaunchData) error {
	if entity == nil {
		return nil
	}
	entityDB := models.LTILaunchData{}
	if entityDB.ID != "" {
		q.Where("id = ?", entity.ID).Find(&entityDB)
	}
	if entityDB.ID != "" {
		// Update
		q.Logger.Debug().Msg(fmt.Sprintf("LTILaunchDataQueries: entity update: %+v", entityDB))
		q.Logger.Info().Msg(fmt.Sprintf("LTILaunchDataQueries: entity update launch_id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		q.Logger.Debug().Msg(fmt.Sprintf("LTILaunchDataQueries: entity create: %+v", entity))
		q.Logger.Info().Msg(fmt.Sprintf("LTILaunchDataQueries: entity create launch_id=%+v", entity.ID))
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.Create(entity)
		return result.Error
	}
}

// StoreLaunchData stores the JSON launch data associated with the supplied launch ID.
func (q *LTILaunchDataQueries) StoreLaunchData(launchID string, launchData json.RawMessage, regUid uint) error {
	q.Logger.Info().Msg(fmt.Sprintf("LTILaunchDataQueries: StoreLaunchData by launch_id=%+v", launchID))
	entity := models.LTILaunchData{}
	entity.ID = launchID
	entity.LaunchData = string(launchData)
	entity.LTIFormID = regUid
	return q.Upsert(&entity)
}

// FindLaunchData retrieves previously-stored launch data using the `launchID'. If the launch data cannot be
// found, it returns ErrLaunchDataNotFound.
func (q *LTILaunchDataQueries) FindLaunchData(launchID string) (json.RawMessage, error) {
	q.Logger.Info().Msg(fmt.Sprintf("LTILaunchDataQueries: FindLaunchData by launch_id: %+v", launchID))
	entity, err := q.Get(launchID)
	if err != nil {
		return []byte{}, err
	}
	return []byte(entity.LaunchData), nil
}

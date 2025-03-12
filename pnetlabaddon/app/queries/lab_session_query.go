package queries

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/models"
	"gitlab.com/a10869/api-modules/shared/utils"
	"gorm.io/gorm"
)

type LabSessionQuery struct {
	*gorm.DB
}

var (
	LabSessionNotFoundError = errors.New("Lab Session not found")
)

func (q *LabSessionQuery) GetByAttemptId(attemptId string) (models.LabSession, error) {
	entityDB := models.LabSession{}
	q.Where("lab_session_lid = ?", attemptId).Find(&entityDB)
	if entityDB.LabSessionID != 0 {
		return entityDB, nil
	}
	return entityDB, utils.FiberValidationException{
		Status:    fiber.StatusNotFound,
		Exception: LabSessionNotFoundError,
	}
}

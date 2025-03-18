package queries

import (
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/rs/zerolog"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/shared/utils"
	"gorm.io/gorm"
)

type LTIRoomQueries struct {
	*gorm.DB
	*zerolog.Logger
}

func (q *LTIRoomQueries) Get(id uint) (models.LTIRoom, error) {
	var entity models.LTIRoom
	result := q.Where("id = ?", id).Find(&entity)
	if entity.ID == 0 {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("LTIAttempt not found"),
		}
	}
	return entity, result.Error
}

func (q *LTIRoomQueries) GetByRoomNumber(roomNumber int64) (models.LTIRoom, error) {
	var entity models.LTIRoom
	result := q.Where("room_number = ?", roomNumber).Order(
		`id desc`,
	).Limit(1).Offset(0).Find(&entity)
	if entity.ID == 0 {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("LTIAttempt not found"),
		}
	}
	return entity, result.Error
}

func (q *LTIRoomQueries) Create() (*models.LTIRoom, error) {
	var minRand int64 = 10000
	var maxRand int64 = 99999
	roomNumberFrom, _ := rand.Int(rand.Reader, big.NewInt(maxRand-minRand+1))
	roomNumber := roomNumberFrom.Int64() + minRand
	entity := models.LTIRoom{}
	entity.RoomNumber = roomNumber
	err := q.Model(&models.LTIRoom{}).Where("room_number = ?", roomNumber).Update("room_number", nil).Error
	if err != nil {
		return &entity, err
	}
	err = q.Save(&entity).Error
	return &entity, err
}

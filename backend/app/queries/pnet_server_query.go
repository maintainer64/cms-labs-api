package queries

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ory/go-convenience/stringsx"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/gorm"
)

type PNETServerQueries struct {
	*gorm.DB
}

type PNETServerQueriesListDTO struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	// The type of status
	// enum: all,active
	Status string `json:"status"`
	// The type of orderBy
	// enum: createdAt,unitRate,lastCountUsers
	OrderBy string `json:"order_by"`
}

const (
	PNETServerListInputDTOStatusAll    = "all"
	PNETServerListInputDTOStatusActive = "active"
)

const (
	PNETServerListInputDTOOrderByCreatedAt      = "createdAt"
	PNETServerListInputDTOOrderByUnitRate       = "unitRate"
	PNETServerListInputDTOOrderByLastCountUsers = "lastCountUsers"
)

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

func (q *PNETServerQueries) securityTokenGenerate() string {
	uid := uuid.New().String()
	hash, err := bcrypt.GenerateFromPassword([]byte(uid), bcrypt.DefaultCost)
	if err != nil {
		log.Warn().Msg(fmt.Sprintf("PNETServerQueries: securityTokenGenerate error: %+v", err))
		return ""
	}
	hasher := sha256.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (q *PNETServerQueries) Upsert(entity *models.PNETServer) error {
	if entity == nil {
		return nil
	}
	entityDB := models.PNETServer{}
	q.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		log.Debug().Msg(fmt.Sprintf("PNETServerQueries: entity update: %+v", entityDB))
		log.Info().Msg(fmt.Sprintf("PNETServerQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		entity.LastOnlineStatus = entityDB.LastOnlineStatus
		entity.Token = stringsx.Coalesce(entityDB.Token, q.securityTokenGenerate())
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		log.Debug().Msg(fmt.Sprintf("PNETServerQueries: entity create: %+v", entity))
		log.Info().Msg(fmt.Sprintf("PNETServerQueries: entity create name=%+v", entity.Name))
		entity.ID = 0
		entity.Token = q.securityTokenGenerate()
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.Create(entity)
		return result.Error
	}
}

func (q *PNETServerQueries) List(filter PNETServerQueriesListDTO) ([]models.PNETServerListItem, int64, error) {
	var entities []models.PNETServerListItem
	result := q.listFilter(
		filter,
		q.Limit(MaxLimitCount).Offset(0),
	).Find(&entities)
	count := result.RowsAffected
	log.Debug().Msg(fmt.Sprintf("PNETServerQueries list: count %+v", count))
	result = q.listFilter(
		filter,
		q.Limit(filter.Limit).Offset(filter.Offset),
	).Find(&entities)
	log.Debug().Msg(fmt.Sprintf("PNETServerQueries list: entities %+v", entities))
	return entities, count, result.Error
}

func (q *PNETServerQueries) listFilter(filter PNETServerQueriesListDTO, tx *gorm.DB) *gorm.DB {
	if filter.OrderBy == PNETServerListInputDTOOrderByLastCountUsers {
		tx = tx.Order(`last_count_users desc`)
	}
	if filter.OrderBy == PNETServerListInputDTOOrderByUnitRate {
		tx = tx.Order(`unit_rate desc`)
	}
	if filter.OrderBy == PNETServerListInputDTOOrderByCreatedAt || filter.OrderBy == "" {
		tx = tx.Order(`created_at desc`)
	}

	if filter.Search != "" {
		tx = tx.Where("name LIKE ?", fmt.Sprintf("%%%s%%", filter.Search))
		tx = tx.Or("url LIKE ?", fmt.Sprintf("%%%s%%", filter.Search))
		tx = tx.Or("id = ?", filter.Search)
	}

	if filter.Status == PNETServerListInputDTOStatusActive {
		tx = models.PNETServeIsRealActive(tx)
	}

	return tx
}

func (q *PNETServerQueries) Delete(id uint) error {
	tx := q.Where("id = ?", id).Delete(&models.PNETServer{})
	log.Debug().Msg(fmt.Sprintf("PNETServerQueries: delete entity by id: %+v", id))
	return tx.Error
}

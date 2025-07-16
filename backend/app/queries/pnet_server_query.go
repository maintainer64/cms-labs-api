package queries

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"

	"github.com/google/uuid"
	"github.com/ory/go-convenience/stringsx"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/shared/utils"
	"gorm.io/gorm"
)

type PNETServerQueries struct {
	DB     *gorm.DB
	Logger *zerolog.Logger
}

type PNETServerQueriesListDTO struct {
	Search string   `json:"search"`
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
	Types  []string `json:"types"`
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

func (q *PNETServerQueries) tableName(object interface{}) string {
	stmt := &gorm.Statement{DB: q.DB}
	_ = stmt.Parse(object)
	return stmt.Schema.Table
}

func (q *PNETServerQueries) Get(id uint) (models.PNETServer, error) {
	var entity models.PNETServer
	result := q.DB.First(&entity, id)
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
		q.Logger.Warn().Msg(fmt.Sprintf("PNETServerQueries: securityTokenGenerate error: %+v", err))
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
	q.DB.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		q.Logger.Debug().Msg(fmt.Sprintf("PNETServerQueries: entity update: %+v", entityDB))
		q.Logger.Info().Msg(fmt.Sprintf("PNETServerQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
		entity.LastOnlineStatus = entityDB.LastOnlineStatus
		entity.Token = stringsx.Coalesce(entityDB.Token, q.securityTokenGenerate())
		result := q.DB.Save(&entity)
		return result.Error
	} else {
		// Create
		q.Logger.Debug().Msg(fmt.Sprintf("PNETServerQueries: entity create: %+v", entity))
		q.Logger.Info().Msg(fmt.Sprintf("PNETServerQueries: entity create name=%+v", entity.Name))
		entity.ID = 0
		entity.Token = q.securityTokenGenerate()
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		result := q.DB.Create(entity)
		return result.Error
	}
}

func (q *PNETServerQueries) List(filter PNETServerQueriesListDTO) ([]models.PNETServerListItem, int64, error) {
	var entities []models.PNETServerListItem
	result := q.listFilter(
		filter,
		q.DB.Limit(MaxLimitCount).Offset(0),
	).Find(&entities)
	count := result.RowsAffected
	q.Logger.Debug().Msg(fmt.Sprintf("PNETServerQueries list: count %+v", count))
	result = q.listFilter(
		filter,
		q.DB.Limit(filter.Limit).Offset(filter.Offset),
	).Find(&entities)
	q.Logger.Debug().Msg(fmt.Sprintf("PNETServerQueries list: entities %+v", entities))
	return entities, count, result.Error
}

func (q *PNETServerQueries) listFilter(filter PNETServerQueriesListDTO, tx *gorm.DB) *gorm.DB {
	tx = tx.Table(q.tableName(&models.PNETServer{}) + " AS pnet_servers")

	switch filter.OrderBy {
	case PNETServerListInputDTOOrderByLastCountUsers:
		tx = tx.Order("last_count_users DESC")
	case PNETServerListInputDTOOrderByUnitRate:
		tx = tx.Order("unit_rate DESC")
	default:
		tx = tx.Order("created_at DESC")
	}

	if filter.Status == PNETServerListInputDTOStatusActive {
		tx = models.PNETServeIsRealActive(tx)
	}

	if len(filter.Types) > 0 {
		tx = tx.Where("type IN (?)", filter.Types)
	}

	if filter.Search != "" {
		searchPattern := "%" + filter.Search + "%"
		tx = tx.Where("name LIKE ? OR url LIKE ? OR id = ? OR client_id = ?", searchPattern, searchPattern, filter.Search, filter.Search)
	}

	return tx
}

func (q *PNETServerQueries) Delete(id uint) error {
	tx := q.DB.Where("id = ?", id).Delete(&models.PNETServer{})
	q.Logger.Debug().Msg(fmt.Sprintf("PNETServerQueries: delete entity by id: %+v", id))
	return tx.Error
}

func (q *PNETServerQueries) GetByClientId(clientID string) (models.PNETServer, error) {
	var entity models.PNETServer
	result := q.DB.Where("client_id = ?", clientID).Limit(1).Find(&entity)
	if entity.ID == 0 {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("PNETServer not found"),
		}
	}
	return entity, result.Error
}

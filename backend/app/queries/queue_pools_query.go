package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/thoas/go-funk"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/gorm"
)

type RoundQueuePoolQueries struct {
	*gorm.DB
}

type RoundQueuePoolPnetListItem struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Type        string    `gorm:"type:varchar(255)" json:"type" validate:"required"`
	ConnectedAt time.Time `gorm:"type:datetime(3)" json:"connected_at" validate:"required"`
	LastUsed    bool      `json:"last_used" validate:"required"`
}

func (q *RoundQueuePoolQueries) Get(id uint) (models.RoundQueuePool, error) {
	var entity models.RoundQueuePool
	result := q.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, utils.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("RoundQueuePool not found"),
		}
	}
	return entity, result.Error
}

func (q *RoundQueuePoolQueries) UpsertQueueByPnetServerIds(ids []uint) error {
	log.Info().Msg(fmt.Sprintf("RoundQueuePoolQueries upsert queue by pnet-server-ids: count %+v", len(ids)))
	entities := funk.Map(ids, func(pnetServerId uint) models.RoundQueuePool {
		entity := models.RoundQueuePool{}
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		entity.ConnectedAt = time.Now().UTC()
		entity.LastUsed = false
		entity.IsActive = true
		entity.PNETServerID = pnetServerId
		entity.Type = models.RoundQueuePoolTypePNET
		return entity
	}).([]models.RoundQueuePool)
	err := q.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("type = ?", models.RoundQueuePoolTypePNET).Delete(&models.RoundQueuePool{}).Error; err != nil {
				return err
			}
			if err := tx.CreateInBatches(entities, 10).Error; err != nil {
				return err
			}
			return nil
		},
	)
	return err
}

func (q *RoundQueuePoolQueries) List() ([]RoundQueuePoolPnetListItem, error) {
	var entities []RoundQueuePoolPnetListItem
	query := q.Model(&models.RoundQueuePool{}).Select(
		"pnet_servers.id, pnet_servers.name, round_queue_pools.type, round_queue_pools.last_used, round_queue_pools.connected_at",
	).Joins(
		"join pnet_servers on pnet_servers.id = round_queue_pools.pnet_server_id",
	).Where(
		"round_queue_pools.type = ?",
		models.RoundQueuePoolTypePNET,
	)
	query = models.PNETServeIsRealActive(query)
	query = query.Limit(MaxLimitCount).Offset(0).Order("round_queue_pools.connected_at asc")
	if err := query.Scan(&entities).Error; err != nil {
		return entities, err
	}
	return entities, nil
}

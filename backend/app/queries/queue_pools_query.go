package queries

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/gorm"
)

type RoundQueuePoolQueries struct {
	*gorm.DB
}

func (q *RoundQueuePoolQueries) tableName(object interface{}) string {
	stmt := &gorm.Statement{DB: q.DB}
	stmt.Parse(object)
	return stmt.Schema.Table
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
	var entities []models.RoundQueuePool
	for index, pnetServerId := range ids {
		entity := models.RoundQueuePool{}
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		entity.ConnectedAt = time.Now().UTC().Add(time.Duration(index-len(ids)) * time.Minute)
		entity.LastUsed = false
		entity.IsActive = true
		entity.PNETServerID = pnetServerId
		entity.Type = models.RoundQueuePoolTypePNET
		entities = append(entities, entity)
	}
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
	query := q.Table(
		q.tableName(&models.RoundQueuePool{})+" AS round_queue_pools",
	).Select(
		"pnet_servers.id, pnet_servers.name, round_queue_pools.type, round_queue_pools.last_used, round_queue_pools.connected_at",
	).Joins(
		"join "+q.tableName(&models.PNETServer{})+" pnet_servers on pnet_servers.id = round_queue_pools.pnet_server_id",
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

func (q *RoundQueuePoolQueries) GetNextByType(poolType string) (models.RoundQueuePool, error) {
	var entityID uint
	err := q.Transaction(
		func(tx *gorm.DB) error {
			entity, err := q.nextPoolItemByType(tx, poolType)
			if err != nil {
				return err
			}
			entityID = entity.ID
			return q.finishDistribution(tx, entity.Type, entity.ID)
		},
	)
	if err != nil {
		return models.RoundQueuePool{}, err
	}
	return q.Get(entityID)
}

func (q *RoundQueuePoolQueries) nextPoolItemByType(tx *gorm.DB, poolType string) (models.RoundQueuePool, error) {
	var entities []models.RoundQueuePool
	log.Info().Msg(fmt.Sprintf("RoundQueuePoolQueries get next pool item from queue by type: %+v", poolType))
	// Получаем последнюю сущность last_used = true
	selectCurrentDistribution := tx.Table(
		q.tableName(&models.RoundQueuePool{})+" AS round_queue_pools",
	).Select(
		"round_queue_pools.connected_at",
	).Where(
		"round_queue_pools.type = ?",
		poolType,
	).Where(
		"round_queue_pools.last_used = ?",
		true,
	).Limit(1)
	// Получаем следующую сущность для распределения относительно времени
	selectNextDistribution := tx.Table(
		q.tableName(&models.RoundQueuePool{})+" AS round_queue_pools",
	).Joins(
		"join "+q.tableName(&models.PNETServer{})+" pnet_servers on pnet_servers.id = round_queue_pools.pnet_server_id",
		q.tableName(&models.PNETServer{}),
	).Where(
		"round_queue_pools.type = ?",
		poolType,
	).Where(
		"round_queue_pools.is_active = ?",
		true,
	).Where(
		"round_queue_pools.last_used = ?",
		false,
	).Where(
		"round_queue_pools.connected_at > (?)",
		selectCurrentDistribution,
	)
	selectNextDistribution = models.PNETServeIsRealActive(selectNextDistribution)
	selectNextDistribution = selectNextDistribution.Limit(1).Offset(0).Order(
		"round_queue_pools.connected_at asc",
	)
	if err := selectNextDistribution.Scan(&entities).Error; err != nil {
		log.Info().Msg(
			fmt.Sprintf(
				"RoundQueuePoolQueries exception for get next pool item from queue by type: %+v",
				poolType,
			),
		)
		return models.RoundQueuePool{}, err
	}
	if len(entities) > 0 {
		entity := entities[0]
		log.Info().Msg(
			fmt.Sprintf(
				"RoundQueuePoolQueries success get id: %+v, pnetServerId: %+v, pool item from queue by type: %+v",
				entity.ID,
				entity.PNETServerID,
				entity.Type,
			),
		)
		return entity, nil
	}
	// Get first entity on distribution
	selectFirstDistribution := tx.Table(q.tableName(&models.RoundQueuePool{})+" AS round_queue_pools").Select(
		"pnet_servers.id, pnet_servers.name, round_queue_pools.type, round_queue_pools.last_used, round_queue_pools.connected_at",
	).Joins(
		"join "+q.tableName(&models.PNETServer{})+" pnet_servers on pnet_servers.id = round_queue_pools.pnet_server_id",
	).Where(
		"round_queue_pools.type = ?",
		poolType,
	).Where(
		"round_queue_pools.is_active = ?",
		true,
	)
	selectFirstDistribution = models.PNETServeIsRealActive(selectFirstDistribution)
	selectFirstDistribution = selectFirstDistribution.Limit(1).Offset(0).Order(
		"round_queue_pools.connected_at asc",
	)
	if err := selectFirstDistribution.Scan(&entities).Error; err != nil {
		log.Info().Msg(
			fmt.Sprintf(
				"RoundQueuePoolQueries exception for get first pool item from queue by type: %+v",
				poolType,
			),
		)
		return models.RoundQueuePool{}, err
	}
	if len(entities) > 0 {
		entity := entities[0]
		log.Info().Msg(
			fmt.Sprintf(
				"RoundQueuePoolQueries success get first entity id: %+v, pnetServerId: %+v, pool item from queue by type: %+v",
				entity.ID,
				entity.PNETServerID,
				entity.Type,
			),
		)
		return entity, nil
	}
	log.Info().Msg(
		fmt.Sprintf(
			"RoundQueuePoolQueries not success get entity from queue by type: %+v",
			poolType,
		),
	)
	return models.RoundQueuePool{}, utils.FiberValidationException{
		Status:    fiber.StatusNotFound,
		Exception: errors.New("no pool item found"),
	}
}

func (q *RoundQueuePoolQueries) finishDistribution(tx *gorm.DB, poolType string, id uint) error {
	log.Info().Msg(
		fmt.Sprintf(
			"RoundQueuePoolQueries: finish distribution item from queue by type: %+v, id: %+v",
			poolType,
			id,
		),
	)
	resetLastDistribution := tx.Table(
		q.tableName(&models.RoundQueuePool{})+" AS round_queue_pools",
	).Where("round_queue_pools.type = ?", poolType).Updates(
		map[string]interface{}{
			"updated_at": time.Now().UTC(),
			"last_used":  false,
		},
	)
	if resetLastDistribution.Error != nil {
		log.Info().Msg(
			fmt.Sprintf(
				"RoundQueuePoolQueries: error on change reset last distribution by type: %+v, id: %+v",
				poolType,
				id,
			),
		)
		return resetLastDistribution.Error
	}
	setLastDistribution := tx.Table(
		q.tableName(&models.RoundQueuePool{})+" AS round_queue_pools",
	).Where("round_queue_pools.id = ?", id).Updates(
		map[string]interface{}{
			"updated_at": time.Now().UTC(),
			"last_used":  true,
		},
	)
	if setLastDistribution.Error != nil {
		log.Info().Msg(
			fmt.Sprintf(
				"RoundQueuePoolQueries: error on change set last distribution by type: %+v, id: %+v",
				poolType,
				id,
			),
		)
		return setLastDistribution.Error
	}
	return nil
}

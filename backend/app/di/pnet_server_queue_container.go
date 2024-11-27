package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases/round_queue_pool_pnet"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

func (di *DIContainer) PnetServerChangeDistributionUC() (*round_queue_pool_pnet.RoundQueuePoolPnetUpsert, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &round_queue_pool_pnet.RoundQueuePoolPnetUpsert{
		PNETServerQueries:     db.PNETServerQueries,
		RoundQueuePoolQueries: db.RoundQueuePoolQueries,
	}, nil
}

func (di *DIContainer) PnetServerListDistributionUC() (*round_queue_pool_pnet.RoundQueuePoolPnetList, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &round_queue_pool_pnet.RoundQueuePoolPnetList{
		RoundQueuePoolQueries: db.RoundQueuePoolQueries,
	}, nil
}

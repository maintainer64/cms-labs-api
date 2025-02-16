package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases/round_queue_pool_pnet"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) PnetServerChangeDistributionUC() *round_queue_pool_pnet.RoundQueuePoolPnetUpsert {
	return &round_queue_pool_pnet.RoundQueuePoolPnetUpsert{
		PNETServerQueries:     di.Queries.PNETServerQueries,
		RoundQueuePoolQueries: di.Queries.RoundQueuePoolQueries,
		Logger:                logs.NewZeroLogger(di.ZeroLogConf.SetName("round_queue_pool_pnet.RoundQueuePoolPnetUpsert")),
	}
}

func (di *DIContainer) PnetServerListDistributionUC() *round_queue_pool_pnet.RoundQueuePoolPnetList {
	return &round_queue_pool_pnet.RoundQueuePoolPnetList{
		RoundQueuePoolQueries: di.Queries.RoundQueuePoolQueries,
		Logger:                logs.NewZeroLogger(di.ZeroLogConf.SetName("round_queue_pool_pnet.RoundQueuePoolPnetList")),
	}
}

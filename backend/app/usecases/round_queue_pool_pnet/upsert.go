package round_queue_pool_pnet

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"

	"github.com/thoas/go-funk"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type RoundQueuePoolPnetUpsert struct {
	PNETServerQueries     *queries.PNETServerQueries
	RoundQueuePoolQueries *queries.RoundQueuePoolQueries
}

type RoundQueuePoolPnetUpsertOutputDTO struct {
}

type RoundQueuePoolPnetUpsertResponse = response.Response[RoundQueuePoolPnetUpsertOutputDTO]

func (u *RoundQueuePoolPnetUpsert) pnetServerEntityToStats(item models.PNETServerListItem) ServerStats {
	return ServerStats{
		ID:             item.ID,
		Name:           item.Name,
		UnitRate:       int(item.UnitRate),
		LastCountUsers: int(item.LastCountUsers),
	}
}

func (u *RoundQueuePoolPnetUpsert) Execute() error {
	log.Info().Msg("RoundQueuePoolPnetUpsert start to generate new distribution pnet servers")
	entities, _, err := u.PNETServerQueries.List(queries.PNETServerQueriesListDTO{
		Limit:  queries.MaxLimitCount,
		Offset: 0,
	})
	if err != nil {
		return err
	}
	stats := funk.Map(entities, u.pnetServerEntityToStats).([]ServerStats)
	distribute := ByPriority(stats).GenerateSequencePriorityDistribute()
	return u.RoundQueuePoolQueries.UpsertQueueByPnetServerIds(distribute)
}

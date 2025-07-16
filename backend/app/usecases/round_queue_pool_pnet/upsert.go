package round_queue_pool_pnet

import (
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"

	funk "github.com/thoas/go-funk"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type RoundQueuePoolPnetUpsert struct {
	PNETServerQueries     *queries.PNETServerQueries
	RoundQueuePoolQueries *queries.RoundQueuePoolQueries
	*zerolog.Logger
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
	u.Logger.Info().Msg("RoundQueuePoolPnetUpsert start to generate new distribution pnet servers")
	entities, _, err := u.PNETServerQueries.List(queries.PNETServerQueriesListDTO{
		Limit:  queries.MaxLimitCount,
		Types:  []string{models.ServerTypePnet},
		Offset: 0,
	})
	if err != nil {
		return err
	}
	stats := funk.Map(entities, u.pnetServerEntityToStats).([]ServerStats)
	distribute := ByPriority(stats).GenerateSequencePriorityDistribute(u.Logger)
	return u.RoundQueuePoolQueries.UpsertQueueByPnetServerIds(distribute)
}

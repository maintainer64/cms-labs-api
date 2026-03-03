package round_queue_pool_pnet

import (
	"github.com/rs/zerolog"
	funk "github.com/thoas/go-funk"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type RoundQueuePoolPnetUpsert struct {
	PNETServerQueries     *queries.PNETServerQueries
	RoundQueuePoolQueries *queries.RoundQueuePoolQueries
	*zerolog.Logger
}

type RoundQueuePoolPnetUpsertRequest struct {
	JSONRPC string       `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string       `json:"method" default:"server_queue.upsert" required:"true"`
	Params  *interface{} `json:"params,omitempty"`
	ID      string       `json:"id,omitempty" default:"1" required:"true"`
}

type RoundQueuePoolPnetUpsertOutputDTO struct {
}

type RoundQueuePoolPnetUpsertResponse struct {
	JSONRPC string                            `json:"jsonrpc" default:"2.0" required:"true"`
	Result  RoundQueuePoolPnetUpsertOutputDTO `json:"result,omitempty"`
	Error   interface{}                       `json:"error,omitempty"`
	ID      string                            `json:"id,omitempty" default:"1" required:"true"`
}

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
	if len(entities) == 0 {
		return nil
	}
	stats := funk.Map(entities, u.pnetServerEntityToStats).([]ServerStats)
	distribute := ByPriority(stats).GenerateSequencePriorityDistribute(u.Logger)
	return u.RoundQueuePoolQueries.UpsertQueueByPnetServerIds(distribute)
}

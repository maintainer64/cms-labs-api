package round_queue_pool_pnet

import (
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type RoundQueuePoolPnetList struct {
	RoundQueuePoolQueries *queries.RoundQueuePoolQueries
	*zerolog.Logger
}

type RoundQueuePoolPnetListOutputDTO struct {
	Model []queries.RoundQueuePoolPnetListItem `json:"model" validate:"required"`
}

type RoundQueuePoolPnetListRequest struct {
	JSONRPC string       `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string       `json:"method" default:"server_queue.list" required:"true"`
	Params  *interface{} `json:"params,omitempty"`
	ID      string       `json:"id,omitempty" default:"1" required:"true"`
}

type RoundQueuePoolPnetListResponse struct {
	JSONRPC string                          `json:"jsonrpc" default:"2.0" required:"true"`
	Result  RoundQueuePoolPnetListOutputDTO `json:"result,omitempty"`
	Error   interface{}                     `json:"error,omitempty"`
	ID      string                          `json:"id,omitempty" default:"1" required:"true"`
}

func (u *RoundQueuePoolPnetList) Execute() (RoundQueuePoolPnetListOutputDTO, error) {
	u.Logger.Info().Msg("RoundQueuePoolPnetList execute")
	entities, err := u.RoundQueuePoolQueries.List()
	return RoundQueuePoolPnetListOutputDTO{
		Model: entities,
	}, err
}

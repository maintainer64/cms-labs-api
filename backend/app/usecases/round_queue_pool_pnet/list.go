package round_queue_pool_pnet

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"

	"gitlab.com/a10869/api-modules/backend/app/queries"
)

type RoundQueuePoolPnetList struct {
	RoundQueuePoolQueries *queries.RoundQueuePoolQueries
}

type RoundQueuePoolPnetListOutputDTO struct {
	Model []queries.RoundQueuePoolPnetListItem `json:"model" validate:"required"`
}

type RoundQueuePoolPnetListResponse = response.Response[RoundQueuePoolPnetListOutputDTO]

func (u *RoundQueuePoolPnetList) Execute() (RoundQueuePoolPnetListOutputDTO, error) {
	log.Info().Msg("RoundQueuePoolPnetList execute")
	entities, err := u.RoundQueuePoolQueries.List()
	return RoundQueuePoolPnetListOutputDTO{
		Model: entities,
	}, err
}

package round_queue_pool_pnet

import (
	"fmt"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

type RoundQueuePoolPnetList struct {
	RoundQueuePoolQueries *queries.RoundQueuePoolQueries
}

type RoundQueuePoolPnetListOutputDTO struct {
	Model []queries.RoundQueuePoolPnetListItem `json:"model" validate:"required"`
}

type RoundQueuePoolPnetListResponse = usecases.Response[RoundQueuePoolPnetListOutputDTO]

func (u *RoundQueuePoolPnetList) Execute() (RoundQueuePoolPnetListOutputDTO, error) {
	log.Info().Msg(fmt.Sprintf("RoundQueuePoolPnetList execute"))
	entities, err := u.RoundQueuePoolQueries.List()
	return RoundQueuePoolPnetListOutputDTO{
		Model: entities,
	}, err
}

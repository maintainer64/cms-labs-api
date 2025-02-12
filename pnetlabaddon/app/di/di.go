package di

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gitlab.com/a10869/api-modules/pnetlabaddon/platform/database"
	"gitlab.com/a10869/api-modules/shared/utils"
)

type DIContainer struct {
	Queries *database.Queries
}

func (di *DIContainer) Close() {
	_ = di.Queries.Close()
}

func NewDIContainer() (*DIContainer, error) {
	queries, err := database.OpenDBConnection()
	di := &DIContainer{
		Queries: queries,
	}
	if err != nil {
		log.Warn().Err(err).Msg("failed to open database connection")
		return di, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return di, nil
}

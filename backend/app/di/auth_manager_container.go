package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

func (di *DIContainer) AuthTokenManager() (*auth.TokenManager, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &auth.TokenManager{
		UserQueries:      db.UserQueries,
		UserTokenQueries: db.UserTokenQueries,
	}, nil
}

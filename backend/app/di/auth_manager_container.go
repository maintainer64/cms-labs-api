package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/backend/platform/database"
	"gitlab.com/a10869/api-modules/shared/utils"
)

func (di *DIContainer) authTokenManager(db *database.Queries) *auth.TokenManager {
	return &auth.TokenManager{
		UserQueries:         db.UserQueries,
		UserPasswordQueries: db.UserPasswordQueries,
		TokenAttemptQueries: db.TokenAttemptQueries,
	}
}

func (di *DIContainer) AuthTokenManager() (*auth.TokenManager, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return di.authTokenManager(db), nil
}

func (di *DIContainer) SSOTokenUC() (*auth.SSOTokenUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &auth.SSOTokenUC{
		TokenAttemptQueries: db.TokenAttemptQueries,
		TokenManager:        di.authTokenManager(db),
		PNETServerQueries:   db.PNETServerQueries,
	}, nil
}

func (di *DIContainer) SSOIntrospectUC() (*auth.SSOIntrospectUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &auth.SSOIntrospectUC{
		TokenAttemptQueries: db.TokenAttemptQueries,
		PNETServerQueries:   db.PNETServerQueries,
	}, nil
}

func (di *DIContainer) SSOAuthorizeUC() (*auth.SSOAuthorizeUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &auth.SSOAuthorizeUC{
		TokenAttemptQueries: db.TokenAttemptQueries,
		PNETServerQueries:   db.PNETServerQueries,
	}, nil
}

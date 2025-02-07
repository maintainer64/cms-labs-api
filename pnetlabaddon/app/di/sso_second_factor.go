package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/usecases"
	"gitlab.com/a10869/api-modules/shared/utils"
)

func (di *DIContainer) SSOSecondFactorUC() (*usecases.SSOSecondFactorUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{
			Status:    fiber.StatusInternalServerError,
			Exception: err,
		}
	}
	return &usecases.SSOSecondFactorUC{
		CMSClient:       di.CMSClient(),
		UserQueries:     db.UserQueries,
		UserRoleQueries: db.UserRoleQueries,
	}, nil
}

package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases/tasks"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

func (di *DIContainer) TaskStartup() (*tasks.StartupFiberUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &tasks.StartupFiberUC{
		UserQueries:      db.UserQueries,
		UserTokenQueries: db.UserTokenQueries,
	}, nil
}

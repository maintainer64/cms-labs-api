package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/shared/utils"
)

func (di *DIContainer) PNETServerEditUC() (*usecases.PNETServerEditUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.PNETServerEditUC{
		PNETServerQueries: db.PNETServerQueries,
	}, nil
}

func (di *DIContainer) PNETServerGetUC() (*usecases.PNETServerGetUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.PNETServerGetUC{
		PNETServerQueries: db.PNETServerQueries,
	}, nil
}

func (di *DIContainer) PNETServerListUC() (*usecases.PNETServerListUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.PNETServerListUC{
		PNETServerQueries: db.PNETServerQueries,
	}, nil
}

func (di *DIContainer) PNETServerDeleteUC() (*usecases.PNETServerDeleteUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.PNETServerDeleteUC{
		PNETServerQueries: db.PNETServerQueries,
	}, nil
}

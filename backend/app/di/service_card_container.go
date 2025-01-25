package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

func (di *DIContainer) ServiceCardEditUC() (*usecases.ServiceCardEditUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.ServiceCardEditUC{
		ServiceCardQueries: db.ServiceCardQueries,
	}, nil
}

func (di *DIContainer) ServiceCardGetUC() (*usecases.ServiceCardGetUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.ServiceCardGetUC{
		ServiceCardQueries: db.ServiceCardQueries,
	}, nil
}

func (di *DIContainer) ServiceCardListUC() (*usecases.ServiceCardListUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.ServiceCardListUC{
		ServiceCardQueries: db.ServiceCardQueries,
	}, nil
}

func (di *DIContainer) ServiceCardDeleteUC() (*usecases.ServiceCardDeleteUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.ServiceCardDeleteUC{
		ServiceCardQueries: db.ServiceCardQueries,
	}, nil
}

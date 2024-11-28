package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

func (di *DIContainer) LTIRoutingEditUC() (*usecases.LTIRoutingEditUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIRoutingEditUC{
		LTIRoutingQueries: db.LTIRoutingQueries,
	}, nil
}

func (di *DIContainer) LTIRoutingGetUC() (*usecases.LTIRoutingGetUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIRoutingGetUC{
		LTIRoutingQueries: db.LTIRoutingQueries,
	}, nil
}

func (di *DIContainer) LTIRoutingListUC() (*usecases.LTIRoutingListUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIRoutingListUC{
		LTIRoutingQueries: db.LTIRoutingQueries,
	}, nil
}

func (di *DIContainer) LTIRoutingDeleteUC() (*usecases.LTIRoutingDeleteUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIRoutingDeleteUC{
		LTIRoutingQueries: db.LTIRoutingQueries,
	}, nil
}


package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

func (di *DIContainer) LtiFormEditUC() (*usecases.LtiFormEditUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LtiFormEditUC{
		LTIFromQuery: db.LTIFromQueries,
	}, nil
}

func (di *DIContainer) LtiFormGetUC() (*usecases.LtiFormGetUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LtiFormGetUC{
		LTIFromQuery: db.LTIFromQueries,
	}, nil
}

func (di *DIContainer) LtiFormListUC() (*usecases.LtiFormListUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LtiFormListUC{
		LTIFromQuery: db.LTIFromQueries,
	}, nil
}

func (di *DIContainer) LtiFormDeleteUC() (*usecases.LtiFormDeleteUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LtiFormDeleteUC{
		LTIFromQuery: db.LTIFromQueries,
	}, nil
}

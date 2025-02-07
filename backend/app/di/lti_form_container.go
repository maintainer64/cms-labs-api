package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/shared/utils"
)

func (di *DIContainer) LTIFormEditUC() (*usecases.LTIFormEditUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIFormEditUC{
		LTIFormQueries: db.LTIFormQueries,
	}, nil
}

func (di *DIContainer) LTIFormGetUC() (*usecases.LTIFormGetUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIFormGetUC{
		LTIFormQueries: db.LTIFormQueries,
	}, nil
}

func (di *DIContainer) LTIFormListUC() (*usecases.LTIFormListUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIFormListUC{
		LTIFormQueries: db.LTIFormQueries,
	}, nil
}

func (di *DIContainer) LTIFormDeleteUC() (*usecases.LTIFormDeleteUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIFormDeleteUC{
		LTIFormQueries: db.LTIFormQueries,
	}, nil
}

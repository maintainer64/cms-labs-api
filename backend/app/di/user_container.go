package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
)

func (di *DIContainer) UserEditUC() (*usecases.UserEditUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.UserEditUC{
		UserQueries: db.UserQueries,
	}, nil
}

func (di *DIContainer) UserGetUC() (*usecases.UserGetUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.UserGetUC{
		UserQueries: db.UserQueries,
	}, nil
}

func (di *DIContainer) UserListUC() (*usecases.UserListUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.UserListUC{
		UserQueries: db.UserQueries,
	}, nil
}

func (di *DIContainer) UserPasswordRecoverUC() (*usecases.UserPasswordRecoverUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.UserPasswordRecoverUC{
		UserTokenQueries: db.UserTokenQueries,
	}, nil
}

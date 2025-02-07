package di

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/shared/utils"
)

func (di *DIContainer) LTIAttemptCreateUC() (*usecases.LTIAttemptCreateUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIAttemptCreateUC{
		LTIAttemptQueries:     db.LTIAttemptQueries,
		LaunchData:            db.LTILaunchDataQueries,
		RoundQueuePoolQueries: db.RoundQueuePoolQueries,
		LTIRoutingQueries:     db.LTIRoutingQueries,
		PNETServerQueries:     db.PNETServerQueries,
		UserQueries:           db.UserQueries,
	}, nil
}

func (di *DIContainer) LTIAttemptEditUC() (*usecases.LTIAttemptEditUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIAttemptEditUC{
		LTIAttemptQueries: db.LTIAttemptQueries,
	}, nil
}

func (di *DIContainer) LTIAttemptGetUC() (*usecases.LTIAttemptGetUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIAttemptGetUC{
		LTIAttemptQueries: db.LTIAttemptQueries,
	}, nil
}

func (di *DIContainer) LTIAttemptListUC() (*usecases.LTIAttemptListUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIAttemptListUC{
		LTIAttemptQueries: db.LTIAttemptQueries,
	}, nil
}

func (di *DIContainer) LTIAttemptDeleteUC() (*usecases.LTIAttemptDeleteUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.LTIAttemptDeleteUC{
		LTIAttemptQueries: db.LTIAttemptQueries,
	}, nil
}

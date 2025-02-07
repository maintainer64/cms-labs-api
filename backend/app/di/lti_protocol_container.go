package di

import (
	"github.com/gofiber/fiber/v2"
	datastore "gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
	"gitlab.com/a10869/api-modules/backend/app/usecases/lti_connector"
	"gitlab.com/a10869/api-modules/backend/app/usecases/lti_launch"
	"gitlab.com/a10869/api-modules/backend/app/usecases/lti_login"
	"gitlab.com/a10869/api-modules/shared/utils"
)

func (di *DIContainer) LTIProtocolDatastoreConfig() (*lti_connector.LTIConnectorAPI, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &lti_connector.LTIConnectorAPI{
		LTIFormQueries:       db.LTIFormQueries,
		LTILaunchDataQueries: db.LTILaunchDataQueries,
		UserQueries:          db.UserQueries,
		LTIProtocolDatastoreConfig: &datastore.Config{
			Registrations: db.LTIFormQueries,
			Nonces:        db.LTINonceTokenQueries,
			LaunchData:    db.LTILaunchDataQueries,
			AccessTokens:  db.LTIAccessTokenQueries,
			UserStore:     db.UserQueries,
		},
	}, nil
}

func (di *DIContainer) LTIProtocolLogin() (*lti_login.Login, error) {
	config, err := di.LTIProtocolDatastoreConfig()
	if err != nil {
		return nil, err
	}
	return lti_login.New(config.LTIProtocolDatastoreConfig), nil
}

func (di *DIContainer) LTIProtocolLaunch() (*lti_launch.Launch, error) {
	config, err := di.LTIProtocolDatastoreConfig()
	if err != nil {
		return nil, err
	}
	return lti_launch.New(config.LTIProtocolDatastoreConfig), nil
}

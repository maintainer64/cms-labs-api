package di

import (
	datastore "gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
	"gitlab.com/a10869/api-modules/backend/app/usecases/lti_connector"
	"gitlab.com/a10869/api-modules/backend/app/usecases/lti_launch"
	"gitlab.com/a10869/api-modules/backend/app/usecases/lti_login"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) LTIProtocolDatastoreConfig() *lti_connector.LTIConnectorAPI {
	logger := logs.NewZeroLogger(di.ZeroLogConf.SetName("lti_connector.LTIConnectorAPI"))
	return &lti_connector.LTIConnectorAPI{
		AuthProviderQueries:  di.Queries.AuthProviderQueries,
		LTILaunchDataQueries: di.Queries.LTILaunchDataQueries,
		UserQueries:          di.Queries.UserQueries,
		LTIProtocolDatastoreConfig: &datastore.Config{
			Registrations: di.Queries.AuthProviderQueries,
			Nonces:        di.Queries.LTINonceTokenQueries,
			LaunchData:    di.Queries.LTILaunchDataQueries,
			AccessTokens:  di.Queries.LTIAccessTokenQueries,
			UserStore: &datastore.ExternalDataUpdateQuery{
				UserQueries: di.Queries.UserQueries,
				RoleQueries: di.Queries.RoleQueries,
				Logger:      logger,
			},
		},
		Logger: logger,
	}
}

func (di *DIContainer) LTIProtocolLogin() *lti_login.Login {
	config := di.LTIProtocolDatastoreConfig()
	return lti_login.New(config.LTIProtocolDatastoreConfig)
}

func (di *DIContainer) LTIProtocolLaunch() *lti_launch.Launch {
	config := di.LTIProtocolDatastoreConfig()
	return lti_launch.New(config.LTIProtocolDatastoreConfig)
}

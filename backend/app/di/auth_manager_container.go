package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/backend/platform/database"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) authTokenManager(db *database.Queries) *auth.TokenManager {
	return &auth.TokenManager{
		UserQueries:         db.UserQueries,
		PNETServerQueries:   db.PNETServerQueries,
		UserPasswordQueries: db.UserPasswordQueries,
		TokenAttemptQueries: db.TokenAttemptQueries,
	}
}

func (di *DIContainer) AuthTokenManager() *auth.TokenManager {
	return di.authTokenManager(di.Queries)
}

func (di *DIContainer) SSOTokenUC() *auth.SSOTokenUC {
	return &auth.SSOTokenUC{
		TokenAttemptQueries: di.Queries.TokenAttemptQueries,
		TokenManager:        di.authTokenManager(di.Queries),
		PNETServerQueries:   di.Queries.PNETServerQueries,
		Logger:              logs.NewZeroLogger(di.ZeroLogConf.SetName("auth.SSOTokenUC")),
	}
}

func (di *DIContainer) SSOIntrospectUC() *auth.SSOIntrospectUC {
	return &auth.SSOIntrospectUC{
		TokenAttemptQueries: di.Queries.TokenAttemptQueries,
		PNETServerQueries:   di.Queries.PNETServerQueries,
	}
}

func (di *DIContainer) SSOAuthorizeUC() *auth.SSOAuthorizeUC {
	return &auth.SSOAuthorizeUC{
		TokenAttemptQueries: di.Queries.TokenAttemptQueries,
		PNETServerQueries:   di.Queries.PNETServerQueries,
		Logger:              logs.NewZeroLogger(di.ZeroLogConf.SetName("auth.SSOAuthorizeUC")),
	}
}

func (di *DIContainer) SSOOpenidConfigurationUC() *auth.SSOOpenidConfiguration {
	return &auth.SSOOpenidConfiguration{}
}

func (di *DIContainer) SSOJwksUC() *auth.SSOJwksUC {
	return &auth.SSOJwksUC{}
}

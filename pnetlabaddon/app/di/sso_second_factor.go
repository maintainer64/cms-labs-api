package di

import (
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/usecases"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) SSOSecondFactorUC() *usecases.SSOSecondFactorUC {
	return &usecases.SSOSecondFactorUC{
		CMSClient:        di.CMSClient(),
		UserQueries:      di.Queries.UserQueries,
		UserRoleQueries:  di.Queries.UserRoleQueries,
		GuacamoleQueries: di.Queries.GuacamoleQueries,
		GuacamoleClient:  di.GuacamoleClient(),
		Logger:           logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.SSOSecondFactorUC")),
	}
}

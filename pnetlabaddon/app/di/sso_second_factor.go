package di

import (
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/usecases"
)

func (di *DIContainer) SSOSecondFactorUC() *usecases.SSOSecondFactorUC {
	return &usecases.SSOSecondFactorUC{
		CMSClient:       di.CMSClient(),
		UserQueries:     di.Queries.UserQueries,
		UserRoleQueries: di.Queries.UserRoleQueries,
	}
}

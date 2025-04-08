package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases/tasks"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) TaskStartup() *tasks.StartupFiberUC {
	return &tasks.StartupFiberUC{
		UserQueries:         di.Queries.UserQueries,
		UserPasswordQueries: di.Queries.UserPasswordQueries,
		RoleQueries:         di.Queries.RoleQueries,
		Logger:              logs.NewZeroLogger(di.ZeroLogConf.SetName("tasks.StartupFiberUC")),
	}
}

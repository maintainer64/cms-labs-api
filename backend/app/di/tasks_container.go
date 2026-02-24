package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases/tasks"
	"gitlab.com/a10869/api-modules/backend/pkg/configs"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) TaskStartup() *tasks.StartupFiberUC {
	return &tasks.StartupFiberUC{
		UserDefaultCreateUC: di.TaskUserDefaultCreateUC(),
		ProxmoxSyncUC:       di.TaskProxmoxSyncUC(),
	}
}

func (di *DIContainer) TaskUserDefaultCreateUC() *tasks.UserDefaultCreateUC {
	return &tasks.UserDefaultCreateUC{
		UserQueries:         di.Queries.UserQueries,
		UserPasswordQueries: di.Queries.UserPasswordQueries,
		RoleQueries:         di.Queries.RoleQueries,
		Logger:              logs.NewZeroLogger(di.ZeroLogConf.SetName("tasks.StartupFiberUC")),
	}
}

func (di *DIContainer) TaskProxmoxSyncUC() *tasks.ProxmoxSyncUC {
	return &tasks.ProxmoxSyncUC{
		Debug:                 configs.AppConfig.Debug,
		ProxmoxSyncConfig:     configs.AppConfig.ProxmoxSyncConfig,
		TargetQueries:         di.Queries.TargetQueries,
		TargetRelationQueries: di.Queries.TargetRelationQueries,
		Logger:                logs.NewZeroLogger(di.ZeroLogConf.SetName("tasks.ProxmoxSyncUC")),
	}
}

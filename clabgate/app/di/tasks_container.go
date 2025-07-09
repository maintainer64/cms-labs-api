package di

import (
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) TasksListUC() *usecases.TasksListUC {
	return &usecases.TasksListUC{
		Logger:    logs.NewZeroLogger(di.ZeroLogConf.SetName("tasks.StartupFiberUC")),
		GitClient: di.GitClient(),
	}
}

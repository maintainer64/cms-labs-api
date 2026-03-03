package di

import (
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) TasksListUC() *usecases.TaskListUC {
	return &usecases.TaskListUC{
		Logger:    logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.TaskListUC")),
		GitClient: di.GitClient(),
	}
}

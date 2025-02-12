package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases/tasks"
)

func (di *DIContainer) TaskStartup() *tasks.StartupFiberUC {
	return &tasks.StartupFiberUC{
		UserQueries:         di.Queries.UserQueries,
		UserPasswordQueries: di.Queries.UserPasswordQueries,
	}
}

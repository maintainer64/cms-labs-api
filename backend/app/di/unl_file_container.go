package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) UNLFileSyncUC() *usecases.UNLFileSyncUC {
	return &usecases.UNLFileSyncUC{
		UNLFileQueries: di.Queries.UNLFileQueries,
		Logger:         logs.NewZeroLogger(di.ZeroLogConf.SetName("UNLFileSyncUC")),
	}
}

func (di *DIContainer) UNLFileGetUC() *usecases.UNLFileGetUC {
	return &usecases.UNLFileGetUC{
		UNLFileQueries: di.Queries.UNLFileQueries,
	}
}

func (di *DIContainer) UNLFileListUC() *usecases.UNLFileListUC {
	return &usecases.UNLFileListUC{
		UNLFileQueries: di.Queries.UNLFileQueries,
	}
}

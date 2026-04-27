package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func (di *DIContainer) ServerEditUC() *usecases.ServerEditUC {
	return &usecases.ServerEditUC{
		ServerQueries: di.Queries.ServerQueries,
		RoleQueries:   di.Queries.RoleQueries,
	}
}

func (di *DIContainer) ServerGetUC() *usecases.ServerGetUC {
	return &usecases.ServerGetUC{
		ServerQueries: di.Queries.ServerQueries,
		RoleQueries:   di.Queries.RoleQueries,
	}
}

func (di *DIContainer) ServerListUC() *usecases.ServerListUC {
	return &usecases.ServerListUC{
		ServerQueries: di.Queries.ServerQueries,
		RoleQueries:   di.Queries.RoleQueries,
	}
}

func (di *DIContainer) ServerDeleteUC() *usecases.ServerDeleteUC {
	return &usecases.ServerDeleteUC{
		ServerQueries: di.Queries.ServerQueries,
	}
}

package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func (di *DIContainer) RoleEditUC() *usecases.RoleEditUC {
	return &usecases.RoleEditUC{
		RoleQueries: di.Queries.RoleQueries,
	}
}

func (di *DIContainer) RoleListUC() *usecases.RoleListUC {
	return &usecases.RoleListUC{
		RoleQueries: di.Queries.RoleQueries,
	}
}

func (di *DIContainer) RoleDeleteUC() *usecases.RoleDeleteUC {
	return &usecases.RoleDeleteUC{
		RoleQueries: di.Queries.RoleQueries,
	}
}

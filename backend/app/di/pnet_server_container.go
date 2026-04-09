package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func (di *DIContainer) PNETServerEditUC() *usecases.PNETServerEditUC {
	return &usecases.PNETServerEditUC{
		PNETServerQueries: di.Queries.PNETServerQueries,
		RoleQueries:       di.Queries.RoleQueries,
	}
}

func (di *DIContainer) PNETServerGetUC() *usecases.PNETServerGetUC {
	return &usecases.PNETServerGetUC{
		PNETServerQueries: di.Queries.PNETServerQueries,
		RoleQueries:       di.Queries.RoleQueries,
	}
}

func (di *DIContainer) PNETServerListUC() *usecases.PNETServerListUC {
	return &usecases.PNETServerListUC{
		PNETServerQueries: di.Queries.PNETServerQueries,
		RoleQueries:       di.Queries.RoleQueries,
	}
}

func (di *DIContainer) PNETServerDeleteUC() *usecases.PNETServerDeleteUC {
	return &usecases.PNETServerDeleteUC{
		PNETServerQueries: di.Queries.PNETServerQueries,
	}
}

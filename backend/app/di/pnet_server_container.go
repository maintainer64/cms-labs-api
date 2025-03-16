package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/external"
)

func (di *DIContainer) PNETServerEditUC() *usecases.PNETServerEditUC {
	return &usecases.PNETServerEditUC{
		PNETServerQueries: di.Queries.PNETServerQueries,
	}
}

func (di *DIContainer) PNETServerGetUC() *usecases.PNETServerGetUC {
	return &usecases.PNETServerGetUC{
		PNETServerQueries: di.Queries.PNETServerQueries,
	}
}

func (di *DIContainer) PNETServerListUC() *usecases.PNETServerListUC {
	return &usecases.PNETServerListUC{
		PNETServerQueries: di.Queries.PNETServerQueries,
	}
}

func (di *DIContainer) PNETServerDeleteUC() *usecases.PNETServerDeleteUC {
	return &usecases.PNETServerDeleteUC{
		PNETServerQueries: di.Queries.PNETServerQueries,
	}
}

func (di *DIContainer) PNETServerPingUC() *external.PNETServerPingUC {
	return &external.PNETServerPingUC{
		PNETServerQueries: di.Queries.PNETServerQueries,
		LTIAttemptQueries: di.Queries.LTIAttemptQueries,
	}
}

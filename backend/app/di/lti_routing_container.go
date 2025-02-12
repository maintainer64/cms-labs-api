package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func (di *DIContainer) LTIRoutingEditUC() *usecases.LTIRoutingEditUC {
	return &usecases.LTIRoutingEditUC{
		LTIRoutingQueries: di.Queries.LTIRoutingQueries,
	}
}

func (di *DIContainer) LTIRoutingGetUC() *usecases.LTIRoutingGetUC {
	return &usecases.LTIRoutingGetUC{
		LTIRoutingQueries: di.Queries.LTIRoutingQueries,
	}
}

func (di *DIContainer) LTIRoutingListUC() *usecases.LTIRoutingListUC {
	return &usecases.LTIRoutingListUC{
		LTIRoutingQueries: di.Queries.LTIRoutingQueries,
	}
}

func (di *DIContainer) LTIRoutingDeleteUC() *usecases.LTIRoutingDeleteUC {
	return &usecases.LTIRoutingDeleteUC{
		LTIRoutingQueries: di.Queries.LTIRoutingQueries,
	}
}

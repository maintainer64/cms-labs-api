package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func (di *DIContainer) LTIFormEditUC() *usecases.LTIFormEditUC {
	return &usecases.LTIFormEditUC{
		LTIFormQueries: di.Queries.LTIFormQueries,
	}
}

func (di *DIContainer) LTIFormGetUC() *usecases.LTIFormGetUC {
	return &usecases.LTIFormGetUC{
		LTIFormQueries: di.Queries.LTIFormQueries,
	}
}

func (di *DIContainer) LTIFormListUC() *usecases.LTIFormListUC {
	return &usecases.LTIFormListUC{
		LTIFormQueries: di.Queries.LTIFormQueries,
	}
}

func (di *DIContainer) LTIFormDeleteUC() *usecases.LTIFormDeleteUC {
	return &usecases.LTIFormDeleteUC{
		LTIFormQueries: di.Queries.LTIFormQueries,
	}
}

func (di *DIContainer) LTIFormListSSOUC() *usecases.LTIFormListSSOUC {
	return &usecases.LTIFormListSSOUC{
		LTIFormQueries: di.Queries.LTIFormQueries,
	}
}

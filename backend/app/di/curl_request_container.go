package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/external"
)

func (di *DIContainer) CurlRequestGetUC() *usecases.CurlRequestGetUC {
	return &usecases.CurlRequestGetUC{
		CurlRequestQueries: di.Queries.CurlRequestQueries,
	}
}

func (di *DIContainer) CurlRequestDeleteUC() *usecases.CurlRequestDeleteUC {
	return &usecases.CurlRequestDeleteUC{
		CurlRequestQueries: di.Queries.CurlRequestQueries,
	}
}

func (di *DIContainer) CurlRequestEditUC() *usecases.CurlRequestEditUC {
	return &usecases.CurlRequestEditUC{
		CurlRequestQueries: di.Queries.CurlRequestQueries,
	}
}

func (di *DIContainer) CurlRequestListUC() *usecases.CurlRequestListUC {
	return &usecases.CurlRequestListUC{
		CurlRequestQueries: di.Queries.CurlRequestQueries,
	}
}

func (di *DIContainer) CurlRequestExecuteUC() *external.CurlRequestExecuteUC {
	return &external.CurlRequestExecuteUC{
		CurlRequestQueries: di.Queries.CurlRequestQueries,
		Client:             NewRestyClient(),
	}
}

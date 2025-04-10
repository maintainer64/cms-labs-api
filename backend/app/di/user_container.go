package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
)

func (di *DIContainer) UserEditUC() *usecases.UserEditUC {
	return &usecases.UserEditUC{
		UserQueries: di.Queries.UserQueries,
		RoleQueries: di.Queries.RoleQueries,
	}
}

func (di *DIContainer) UserGetUC() *usecases.UserGetUC {
	return &usecases.UserGetUC{
		UserQueries: di.Queries.UserQueries,
		RoleQueries: di.Queries.RoleQueries,
	}
}

func (di *DIContainer) UserListUC() *usecases.UserListUC {
	return &usecases.UserListUC{
		UserQueries: di.Queries.UserQueries,
		RoleQueries: di.Queries.RoleQueries,
	}
}

func (di *DIContainer) UserPasswordRecoverUC() *auth.UserPasswordRecoverUC {
	return &auth.UserPasswordRecoverUC{
		UserPasswordQueries: di.Queries.UserPasswordQueries,
	}
}

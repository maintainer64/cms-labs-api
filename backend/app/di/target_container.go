package di

import (
	"gitlab.com/a10869/api-modules/backend/app/addons"
	"gitlab.com/a10869/api-modules/backend/app/addons/vault"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/pkg/configs"
)

func (di *DIContainer) FactoryAddonService() *addons.FactoryAddonService {
	return &addons.FactoryAddonService{
		Config:              configs.AppConfig.AddonsConfig,
		VaultClient:         vault.NewClient(configs.AppConfig.Vault),
		TokenAttemptQueries: di.Queries.TokenAttemptQueries,
		PNETServerQueries:   di.Queries.PNETServerQueries,
		UserQueries:         di.Queries.UserQueries,
		TargetQueries:       di.Queries.TargetQueries,
		IssId:               "",
	}
}

func (di *DIContainer) TargetAddonCreateUC() *usecases.TargetAddonCreateUC {
	return &usecases.TargetAddonCreateUC{
		TargetAddonQueries:  di.Queries.TargetAddonQueries,
		TargetQueries:       di.Queries.TargetQueries,
		TargetUserQueries:   di.Queries.TargetUserQueries,
		FactoryAddonService: di.FactoryAddonService(),
	}
}

func (di *DIContainer) TargetAddonDeleteUC() *usecases.TargetAddonDeleteUC {
	return &usecases.TargetAddonDeleteUC{
		TargetAddonQueries:  di.Queries.TargetAddonQueries,
		TargetQueries:       di.Queries.TargetQueries,
		TargetUserQueries:   di.Queries.TargetUserQueries,
		FactoryAddonService: di.FactoryAddonService(),
	}
}

func (di *DIContainer) TargetAddonResetUC() *usecases.TargetAddonResetUC {
	return &usecases.TargetAddonResetUC{
		TargetAddonQueries:  di.Queries.TargetAddonQueries,
		TargetQueries:       di.Queries.TargetQueries,
		TargetUserQueries:   di.Queries.TargetUserQueries,
		FactoryAddonService: di.FactoryAddonService(),
	}
}

func (di *DIContainer) TargetDeleteUC() *usecases.TargetDeleteUC {
	return &usecases.TargetDeleteUC{
		TargetQueries:     di.Queries.TargetQueries,
		TargetUserQueries: di.Queries.TargetUserQueries,
	}
}

func (di *DIContainer) TargetGetUC() *usecases.TargetGetUC {
	return &usecases.TargetGetUC{
		TargetQueries: di.Queries.TargetQueries,
	}
}

func (di *DIContainer) TargetListUC() *usecases.TargetListUC {
	return &usecases.TargetListUC{
		TargetQueries:         di.Queries.TargetQueries,
		TargetUserQueries:     di.Queries.TargetUserQueries,
		TargetRelationQueries: di.Queries.TargetRelationQueries,
	}
}

func (di *DIContainer) TargetRelationCreateUC() *usecases.TargetRelationCreateUC {
	return &usecases.TargetRelationCreateUC{
		TargetRelationQueries: di.Queries.TargetRelationQueries,
		TargetUserQueries:     di.Queries.TargetUserQueries,
	}
}

func (di *DIContainer) TargetRelationDeleteUC() *usecases.TargetRelationDeleteUC {
	return &usecases.TargetRelationDeleteUC{
		TargetRelationQueries: di.Queries.TargetRelationQueries,
		TargetUserQueries:     di.Queries.TargetUserQueries,
	}
}

func (di *DIContainer) TargetUpsertUC() *usecases.TargetUpsertUC {
	return &usecases.TargetUpsertUC{
		TargetQueries:     di.Queries.TargetQueries,
		TargetUserQueries: di.Queries.TargetUserQueries,
	}
}

func (di *DIContainer) TargetUserDeleteUC() *usecases.TargetUserDeleteUC {
	return &usecases.TargetUserDeleteUC{
		TargetQueries:     di.Queries.TargetQueries,
		TargetUserQueries: di.Queries.TargetUserQueries,
	}
}

func (di *DIContainer) TargetUserUpsertUC() *usecases.TargetUserUpsertUC {
	return &usecases.TargetUserUpsertUC{
		TargetQueries:     di.Queries.TargetQueries,
		TargetUserQueries: di.Queries.TargetUserQueries,
	}
}

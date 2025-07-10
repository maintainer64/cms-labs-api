package di

import (
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/clabgate/pkg/configs"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) TopologiesGetUC() (*usecases.TopologiesGetUC, error) {
	kubeQuery, err := queries.NewKubernetesAdmin(configs.AppConfig.K8S)
	if err != nil {
		return nil, err
	}
	return &usecases.TopologiesGetUC{
		Logger:               logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.TopologiesGetUC")),
		KubernetesAdminQuery: kubeQuery,
	}, nil
}

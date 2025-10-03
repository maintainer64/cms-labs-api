package di

import (
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) TopologiesGetUC() (*usecases.TopologiesGetUC, error) {
	kubeQuery, err := di.KubernetesAdmin()
	if err != nil {
		return nil, err
	}
	return &usecases.TopologiesGetUC{
		Logger:               logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.TopologiesGetUC")),
		KubernetesAdminQuery: kubeQuery,
	}, nil
}

func (di *DIContainer) TopologiesCreateUC() (*usecases.TopologyCreateUC, error) {
	kubeQuery, err := di.KubernetesAdmin()
	if err != nil {
		return nil, err
	}
	return &usecases.TopologyCreateUC{
		Logger:               logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.TopologyCreateUC")),
		KubernetesAdminQuery: kubeQuery,
		GitClient:            di.GitClient(),
	}, nil
}

func (di *DIContainer) TopologiesDeleteUC() (*usecases.TopologyDeleteUC, error) {
	kubeQuery, err := di.KubernetesAdmin()
	if err != nil {
		return nil, err
	}
	return &usecases.TopologyDeleteUC{
		Logger:               logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.TopologyDeleteUC")),
		KubernetesAdminQuery: kubeQuery,
	}, nil
}

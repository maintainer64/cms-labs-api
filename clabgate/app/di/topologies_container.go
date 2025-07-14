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

func (di *DIContainer) TopologiesCreateUC() (*usecases.TopologiesCreateUC, error) {
	kubeQuery, err := di.KubernetesAdmin()
	if err != nil {
		return nil, err
	}
	return &usecases.TopologiesCreateUC{
		Logger:               logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.TopologiesCreateUC")),
		KubernetesAdminQuery: kubeQuery,
		GitClient:            di.GitClient(),
	}, nil
}

func (di *DIContainer) TopologiesDeleteUC() (*usecases.TopologiesDeleteUC, error) {
	kubeQuery, err := di.KubernetesAdmin()
	if err != nil {
		return nil, err
	}
	return &usecases.TopologiesDeleteUC{
		Logger:               logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.TopologiesCreateUC")),
		KubernetesAdminQuery: kubeQuery,
	}, nil
}

package di

import (
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) TokenAccessGetUC() (*usecases.TokenAccessGetUC, error) {
	kubeQuery, err := di.KubernetesAdmin()
	if err != nil {
		return nil, err
	}
	return &usecases.TokenAccessGetUC{
		Logger:               logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.TokenAccessGetUC")),
		KubernetesAdminQuery: kubeQuery,
	}, nil
}

func (di *DIContainer) TokenFileYAMLGetUC() (*usecases.TokenFileYAMLGetUC, error) {
	kubeQuery, err := di.KubernetesAdmin()
	if err != nil {
		return nil, err
	}
	return &usecases.TokenFileYAMLGetUC{
		Logger:               logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.TokenFileYAMLGetUC")),
		KubernetesAdminQuery: kubeQuery,
	}, nil
}

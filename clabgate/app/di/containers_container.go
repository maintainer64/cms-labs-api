package di

import (
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/shared/logs"
)

func (di *DIContainer) ContainersGetUC() (*usecases.ContainersGetUC, error) {
	kubeQuery, err := di.KubernetesAdmin()
	if err != nil {
		return nil, err
	}
	return &usecases.ContainersGetUC{
		Logger:                   logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.ContainersGetUC")),
		KubernetesAdminQuery:     kubeQuery,
		KubeDashboardClientQuery: di.KubeDashboardClient(),
	}, nil
}

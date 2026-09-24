package di

import (
	"time"

	"github.com/maintainer64/cms-labs-api/clabgate/app/usecases"
	"github.com/maintainer64/cms-labs-api/clabgate/pkg/configs"
	"github.com/maintainer64/cms-labs-api/shared/logs"
)

func (di *DIContainer) TopologiesGetUC() (*usecases.TopologiesGetUC, error) {
	kubeQuery, err := di.KubernetesAdmin()
	if err != nil {
		return nil, err
	}
	return &usecases.TopologiesGetUC{
		Logger:               logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.TopologiesGetUC")),
		KubernetesAdminQuery: kubeQuery,
		WorkspaceSecret:      configs.AppConfig.Session.WorkspaceSecret,
		WorkspaceGrantTTL:    time.Duration(configs.AppConfig.Session.WorkspaceGrantTTL) * time.Second,
	}, nil
}

func (di *DIContainer) NodeActionsUC() (*usecases.NodeActionsUC, error) {
	kubeQuery, err := di.KubernetesAdmin()
	if err != nil {
		return nil, err
	}
	return &usecases.NodeActionsUC{
		Logger:               logs.NewZeroLogger(di.ZeroLogConf.SetName("usecases.NodeActionsUC")),
		KubernetesAdminQuery: kubeQuery,
	}, nil
}

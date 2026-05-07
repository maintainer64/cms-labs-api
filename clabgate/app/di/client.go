package di

import (
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
)

func (di *DIContainer) KubernetesAdmin() (*queries.KubernetesAdminQuery, error) {
	kubeQuery, err := queries.NewKubernetesAdmin(
		di.ZeroLogConf.SetName("queries.KubernetesAdminQuery"),
	)
	return kubeQuery, err
}

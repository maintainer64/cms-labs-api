package di

import (
	"github.com/go-resty/resty/v2"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/clabgate/pkg/configs"
	"gitlab.com/a10869/api-modules/shared/logs"
)

var (
	NewRestyClient = func() *resty.Client {
		return resty.New()
	}
)

func (di *DIContainer) GitClient() queries.GitCodeRegistry {
	if (configs.AppConfig.GitlabConfig.BaseUrl) != "" {
		return &queries.GitlabCodeRegistryQuery{
			Config: configs.AppConfig.GitlabConfig,
			Client: NewRestyClient(),
			Logger: logs.NewZeroLogger(di.ZeroLogConf.SetName("queries.GitlabCodeRegistryQuery")),
		}
	}
	return nil
}

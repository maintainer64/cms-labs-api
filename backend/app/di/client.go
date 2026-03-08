package di

import (
	resty "github.com/go-resty/resty/v2"
	"gitlab.com/a10869/api-modules/backend/app/addons/vault"
	"gitlab.com/a10869/api-modules/backend/pkg/configs"
	"gitlab.com/a10869/api-modules/shared/logs"
)

var (
	NewRestyClient = func() *resty.Client {
		return resty.New()
	}
	NewVaultClient = func() vault.ClientInterface {
		return vault.NewClient(
			configs.AppConfig.Vault,
			logs.NewZeroLogger(&logs.ZeroLoggerConf{}),
		)
	}
)

package di

import (
	resty "github.com/go-resty/resty/v2"
	"github.com/maintainer64/cms-labs-api/backend/app/addons/vault"
	"github.com/maintainer64/cms-labs-api/backend/pkg/configs"
	"github.com/maintainer64/cms-labs-api/shared/logs"
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

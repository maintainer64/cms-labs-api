package di

import (
	"github.com/go-resty/resty/v2"
	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/configs"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/guacamole_client"
)

var (
	NewRestyClient = func() *resty.Client {
		return resty.New()
	}
)

func (di *DIContainer) CMSClient() *cms_client.CMSClient {
	return cms_client.NewCMSClient(
		configs.AppConfig.CMSClient,
		NewRestyClient(),
	)
}

func (di *DIContainer) GuacamoleClient() *guacamole_client.GuacamoleClient {
	return guacamole_client.NewGuacamoleClient(
		configs.AppConfig.GuacamoleClient,
		NewRestyClient(),
	)
}

package di

import (
	resty "github.com/go-resty/resty/v2"
	"github.com/maintainer64/cms-labs-api/pnetlabaddon/pkg/configs"
	"github.com/maintainer64/cms-labs-api/shared/cms_client"
	"github.com/maintainer64/cms-labs-api/shared/guacamole_client"
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

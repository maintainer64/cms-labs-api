package di

import (
	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/configs"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

func (di *DIContainer) CMSClient() *cms_client.CMSClient {
	return cms_client.NewCMSClient(configs.AppConfig.CMSClient)
}

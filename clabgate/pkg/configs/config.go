package configs

import (
	"os"
	"strconv"

	"gitlab.com/a10869/api-modules/shared/connection"
)

type AppConfigModel struct {
	Debug               bool
	Server              *connection.ServerConfig
	K8S                 *connection.K8SConfig
	GitlabConfig        *connection.GitConfig
	KubeDashboardConfig *connection.KubeDashboardConfig
}

func (c *AppConfigModel) Reload() {
	c.Debug = os.Getenv("DEBUG") == "true"
	c.Server = &connection.ServerConfig{
		Host:              os.Getenv("SERVER_HOST"),
		Port:              os.Getenv("SERVER_PORT"),
		ServerReadTimeout: getEnvInt("SERVER_READ_TIMEOUT"),
		Layer:             os.Getenv("STAGE_STATUS"),
	}
	c.K8S = &connection.K8SConfig{
		ConfigYaml:     os.Getenv("K8S_CONFIG"),
		IssuerOIDC:     os.Getenv("K8S_OIDC_ISSUER"),
		KrewConfigYaml: os.Getenv("K8S_KREW_CONFIG"),
		Namespace:      os.Getenv("K8S_NAMESPACE"),
	}
	c.GitlabConfig = &connection.GitConfig{
		BaseUrl:      os.Getenv("GITLAB_BASE_URL"),
		RepoId:       os.Getenv("GITLAB_REPO_ID"),
		AccessToken:  os.Getenv("GITLAB_ACCESS_KEY"),
		TriggerToken: os.Getenv("GITLAB_TRIGGER_KEY"),
		Branch:       os.Getenv("GITLAB_BRANCH"),
	}
	c.KubeDashboardConfig = &connection.KubeDashboardConfig{
		BaseUrl:     os.Getenv("KUBE_DASHBOARD_BASE_URL"),
		NoVerifySSL: os.Getenv("KUBE_DASHBOARD_NO_VERIFY") == "true",
	}
}

func getEnvInt(name string) int {
	digit, _ := strconv.Atoi(os.Getenv(name))
	return digit
}

func NewAppConfigModel() *AppConfigModel {
	model := &AppConfigModel{}
	model.Reload()
	return model
}

var AppConfig = NewAppConfigModel()

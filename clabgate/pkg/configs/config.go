package configs

import (
	"gitlab.com/a10869/api-modules/shared/connection"
	"os"
	"strconv"
)

type AppConfigModel struct {
	Debug        bool
	Server       *connection.ServerConfig
	K8S          *connection.K8SConfig
	GitlabConfig *connection.GitConfig
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
		ConfigYaml: os.Getenv("K8S_CONFIG"),
		Namespace:  os.Getenv("K8S_NAMESPACE"),
	}
	c.GitlabConfig = &connection.GitConfig{
		BaseUrl: os.Getenv("GITLAB_BASE_URL"),
		RepoId:  os.Getenv("GITLAB_REPO_ID"),
		Token:   os.Getenv("GITLAB_ACCESS_KEY"),
		Branch:  os.Getenv("GITLAB_BRANCH"),
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

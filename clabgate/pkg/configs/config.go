package configs

import (
	"os"
	"strconv"

	"gitlab.com/a10869/api-modules/shared/connection"
)

type AppConfigModel struct {
	Debug  bool
	Server *connection.ServerConfig
}

func (c *AppConfigModel) Reload() {
	c.Debug = os.Getenv("DEBUG") == "true"
	c.Server = &connection.ServerConfig{
		Host:              os.Getenv("SERVER_HOST"),
		Port:              os.Getenv("SERVER_PORT"),
		ServerReadTimeout: getEnvInt("SERVER_READ_TIMEOUT"),
		Layer:             os.Getenv("STAGE_STATUS"),
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

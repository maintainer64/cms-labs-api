package connection

import (
	"fmt"
	"os"

	"github.com/goccy/go-json"
)

type ProxmoxSyncConfig struct {
	Servers []ProxmoxServerConfig `json:"servers"`
}

type ProxmoxServerConfig struct {
	Url   string `json:"url"`
	Token string `json:"token"`
}

func GetProxmoxConfig(pathConfig string) *ProxmoxSyncConfig {
	var config ProxmoxSyncConfig
	if pathConfig == "" {
		return &ProxmoxSyncConfig{}
	}
	rawConfig, err := os.ReadFile(pathConfig)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(rawConfig, &config); err != nil {
		panic(fmt.Errorf("failed to parse ProxmoxConfig: %w", err))
	}
	return &config
}

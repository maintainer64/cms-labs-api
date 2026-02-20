package connection

import (
	"fmt"

	"github.com/goccy/go-json"
)

type ProxmoxSyncConfig struct {
	Servers []ProxmoxServerConfig `json:"servers"`
}

type ProxmoxServerConfig struct {
	Url   string `json:"url"`
	Token string `json:"token"`
}

func GetProxmoxConfig(rawConfig string) *ProxmoxSyncConfig {
	var config ProxmoxSyncConfig
	if rawConfig == "" {
		return &config
	}
	if err := json.Unmarshal([]byte(rawConfig), &config); err != nil {
		panic(fmt.Errorf("failed to parse ProxmoxConfig: %w", err))
	}
	return &config
}

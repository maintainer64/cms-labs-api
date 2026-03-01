package connection

import (
	"os"

	"github.com/goccy/go-json"
)

type AddonType string

const (
	PostgreSQLAddon AddonType = "postgresql"
	MySQLAddon      AddonType = "mysql"
	S3Addon         AddonType = "s3"
	HarborAddon     AddonType = "harbor"
	KubernetesAddon AddonType = "kubernetes"
	VaultAddon      AddonType = "vault"
)

type AddonConfig struct {
	ID     string         `json:"id"`
	Tag    string         `json:"tag"`
	Name   string         `json:"name"`
	Type   AddonType      `json:"type"`
	Params map[string]any `json:"params"`
}

type AddonsConfig struct {
	Addons []AddonConfig `json:"addons"`
}

func GetAddonsConfig(pathConfig string) *AddonsConfig {
	if pathConfig == "" {
		return &AddonsConfig{}
	}
	var cfg AddonsConfig
	rawConfig, err := os.ReadFile(pathConfig)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(rawConfig, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}

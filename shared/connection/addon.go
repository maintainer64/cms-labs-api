package connection

import "github.com/goccy/go-json"

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

func GetAddonsConfig(rawConfig string) *AddonsConfig {
	if rawConfig == "" {
		return &AddonsConfig{}
	}
	var cfg AddonsConfig
	if err := json.Unmarshal([]byte(rawConfig), &cfg); err != nil {
		panic(err)
	}
	return &cfg
}

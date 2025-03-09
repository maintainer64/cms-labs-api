package guacamole_client

type GuacamoleClientConfig struct {
	Debug             bool   `json:"debug"`
	MaxTimeoutSeconds int64  `json:"max_timeout_seconds"`
	BaseUrl           string `json:"base_url"`
}

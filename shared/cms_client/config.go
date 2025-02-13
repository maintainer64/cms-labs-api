package cms_client

type CMSClientConfig struct {
	Debug             bool   `json:"debug"`
	MaxTimeoutSeconds int64  `json:"max_timeout_seconds"`
	ClientID          string `json:"client_id"`
	Token             string `json:"token"`
	BaseUrl           string `json:"base_url"`
}

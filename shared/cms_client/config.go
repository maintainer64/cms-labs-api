package cms_client

type CMSClientConfig struct {
	Debug             bool   `json:"debug"`
	MaxTimeoutSeconds int64  `json:"max_timeout_seconds"`
	ClientID          string `json:"client_id"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	BaseUrl           string `json:"base_url"`
}

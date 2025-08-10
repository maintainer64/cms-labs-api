package connection

type DBConfig struct {
	Type                   string
	User                   string
	Password               string
	Host                   string
	Port                   string
	Name                   string
	SSL                    string
	MaxConnections         int
	MaxIdleConnections     int
	MaxLifetimeConnections int
	TablePrefix            string
}

type ServerConfig struct {
	Host              string
	Port              string
	ServerReadTimeout int
	Layer             string
}

type K8SConfig struct {
	ConfigYaml     string
	KrewConfigYaml string
	IssuerOIDC     string
	Namespace      string
}

type GitConfig struct {
	BaseUrl      string
	AccessToken  string
	TriggerToken string
	RepoId       string
	Branch       string
}

type KubeDashboardConfig struct {
	BaseUrl     string
	NoVerifySSL bool
}

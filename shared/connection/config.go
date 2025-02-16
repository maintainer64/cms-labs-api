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

package configs

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	Host              string
	Port              string
	ServerReadTimeout int
	Layer             string
}

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
}

type JWTConfig struct {
	SecretKey                string
	SecretKeyExpireMinutes   int
	SecretRefresh            string
	SecretRefreshExpireHours int
}

type AppConfigModel struct {
	Debug  bool
	Server *ServerConfig
	DB     *DBConfig
	JWT    *JWTConfig
}

func getEnvInt(name string) int {
	digit, _ := strconv.Atoi(os.Getenv(name))
	return digit
}

func NewAppConfigModel() *AppConfigModel {
	debug := os.Getenv("DEBUG") == "true"

	return &AppConfigModel{
		Debug: debug,
		Server: &ServerConfig{
			Host:              os.Getenv("SERVER_HOST"),
			Port:              os.Getenv("SERVER_PORT"),
			ServerReadTimeout: getEnvInt("SERVER_READ_TIMEOUT"),
			Layer:             os.Getenv("STAGE_STATUS"),
		},
		DB: &DBConfig{
			Type:                   os.Getenv("DB_TYPE"),
			User:                   os.Getenv("DB_USER"),
			Password:               os.Getenv("DB_PASSWORD"),
			Host:                   os.Getenv("DB_HOST"),
			Port:                   os.Getenv("DB_PORT"),
			Name:                   os.Getenv("DB_NAME"),
			SSL:                    os.Getenv("DB_SSL_MODE"),
			MaxConnections:         getEnvInt("DB_MAX_CONNECTIONS"),
			MaxIdleConnections:     getEnvInt("DB_MAX_IDLE_CONNECTIONS"),
			MaxLifetimeConnections: getEnvInt("DB_MAX_LIFETIME_CONNECTIONS"),
		},
		JWT: &JWTConfig{
			SecretKey:                os.Getenv("JWT_SECRET_KEY"),
			SecretKeyExpireMinutes:   getEnvInt("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"),
			SecretRefresh:            os.Getenv("JWT_REFRESH_KEY"),
			SecretRefreshExpireHours: getEnvInt("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT"),
		},
	}
}

var AppConfig = NewAppConfigModel()

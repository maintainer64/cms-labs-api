package utils

import (
	"fmt"

	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/configs"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "mysql":
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			configs.AppConfig.DB.User,
			configs.AppConfig.DB.Password,
			configs.AppConfig.DB.Host,
			configs.AppConfig.DB.Port,
			configs.AppConfig.DB.Name,
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			configs.AppConfig.Server.Host,
			configs.AppConfig.Server.Port,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}

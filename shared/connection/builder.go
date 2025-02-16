package connection

import "fmt"

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string, dbConfig *DBConfig, fiberConfig *ServerConfig) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "mysql":
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Name,
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			fiberConfig.Host,
			fiberConfig.Port,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}

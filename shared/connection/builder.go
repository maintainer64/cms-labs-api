package connection

import "fmt"

func UrlBuilderMySql(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
}

func UrlBuilderFiber(fiberConfig *ServerConfig) string {
	return fmt.Sprintf(
		"%s:%s",
		fiberConfig.Host,
		fiberConfig.Port,
	)
}

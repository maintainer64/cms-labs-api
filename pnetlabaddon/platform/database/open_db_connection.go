package database

import (
	"os"

	"gorm.io/gorm"

	"gitlab.com/a10869/api-modules/pnetlabaddon/app/queries"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries
	*queries.UserRoleQueries
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define Database connection variables.
	var (
		db  *gorm.DB
		err error
	)

	// Get DB_TYPE value from .env file.
	dbType := os.Getenv("DB_TYPE")

	// Define a new Database connection with right DB type.
	switch dbType {
	case "mysql":
		db, err = MysqlConnection()
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries:     &queries.UserQueries{DB: db},
		UserRoleQueries: &queries.UserRoleQueries{DB: db},
	}, nil
}

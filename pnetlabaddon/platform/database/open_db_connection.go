package database

import (
	"os"

	"gorm.io/gorm"

	"gitlab.com/a10869/api-modules/pnetlabaddon/app/queries"
)

// Queries struct for collect all app queries.
type Queries struct {
	*gorm.DB
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
		DB:              db,
		UserQueries:     &queries.UserQueries{DB: db},
		UserRoleQueries: &queries.UserRoleQueries{DB: db},
	}, nil
}

func (q *Queries) Close() error {
	dbInstance, err := q.DB.DB()
	if err != nil {
		return err
	}
	return dbInstance.Close()
}

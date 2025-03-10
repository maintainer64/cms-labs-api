package database

import (
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/backend/pkg/configs"
	"gitlab.com/a10869/api-modules/shared/connection"

	"gorm.io/gorm"

	"gitlab.com/a10869/api-modules/pnetlabaddon/app/queries"
)

// Queries struct for collect all app queries.
type Queries struct {
	*gorm.DB
	*queries.UserQueries
	*queries.UserRoleQueries
	*queries.GuacamoleQueries
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection(l *zerolog.Logger) (*Queries, error) {
	// Define Database connection variables.
	var (
		db  *gorm.DB
		err error
	)

	db, err = connection.MysqlConnection(configs.AppConfig.DB, l)

	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		DB:               db,
		UserQueries:      &queries.UserQueries{DB: db},
		UserRoleQueries:  &queries.UserRoleQueries{DB: db},
		GuacamoleQueries: &queries.GuacamoleQueries{DB: db, Logger: l},
	}, nil
}

func (q *Queries) Close() error {
	dbInstance, err := q.DB.DB()
	if err != nil {
		return err
	}
	return dbInstance.Close()
}

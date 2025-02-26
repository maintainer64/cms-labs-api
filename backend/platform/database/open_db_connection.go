package database

import (
	"os"

	"github.com/rs/zerolog"

	"gorm.io/gorm"

	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

// Queries struct for collect all app queries.
type Queries struct {
	*gorm.DB
	*lti_query.LTIFormQueries
	*lti_query.LTINonceTokenQueries
	*lti_query.LTIAccessTokenQueries
	*lti_query.LTILaunchDataQueries
	*queries.UserQueries
	*queries.UserPasswordQueries
	*queries.PNETServerQueries
	*queries.RoundQueuePoolQueries
	*queries.LTIRoutingQueries
	*queries.LTIAttemptQueries
	*queries.ServiceCardQueries
	*queries.TokenAttemptQueries
	*queries.UNLFileQueries
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection(l *zerolog.Logger) (*Queries, error) {
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
		db, err = MysqlConnection(l)
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		DB:                    db,
		LTIFormQueries:        &lti_query.LTIFormQueries{DB: db, Logger: l},
		LTINonceTokenQueries:  &lti_query.LTINonceTokenQueries{DB: db, Logger: l},
		LTIAccessTokenQueries: &lti_query.LTIAccessTokenQueries{DB: db, Logger: l},
		LTILaunchDataQueries:  &lti_query.LTILaunchDataQueries{DB: db, Logger: l},
		UserQueries:           &queries.UserQueries{DB: db, Logger: l},
		UserPasswordQueries:   &queries.UserPasswordQueries{DB: db, Logger: l},
		PNETServerQueries:     &queries.PNETServerQueries{DB: db, Logger: l},
		RoundQueuePoolQueries: &queries.RoundQueuePoolQueries{DB: db, Logger: l},
		LTIRoutingQueries:     &queries.LTIRoutingQueries{DB: db, Logger: l},
		LTIAttemptQueries:     &queries.LTIAttemptQueries{DB: db, Logger: l},
		ServiceCardQueries:    &queries.ServiceCardQueries{DB: db, Logger: l},
		TokenAttemptQueries:   &queries.TokenAttemptQueries{DB: db, Logger: l},
		UNLFileQueries:        &queries.UNLFileQueries{DB: db, Logger: l},
	}, nil
}

func (q *Queries) Close() error {
	dbInstance, err := q.DB.DB()
	if err != nil {
		return err
	}
	return dbInstance.Close()
}

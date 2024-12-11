package database

import (
	"os"

	"gorm.io/gorm"

	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

// Queries struct for collect all app queries.
type Queries struct {
	*lti_query.LTIFormQueries
	*lti_query.LTINonceTokenQueries
	*lti_query.LTIAccessTokenQueries
	*lti_query.LTILaunchDataQueries
	*queries.UserQueries
	*queries.UserTokenQueries
	*queries.PNETServerQueries
	*queries.RoundQueuePoolQueries
	*queries.LTIRoutingQueries
	*queries.LTIAttemptQueries
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
		LTIFormQueries:        &lti_query.LTIFormQueries{DB: db},
		LTINonceTokenQueries:  &lti_query.LTINonceTokenQueries{DB: db},
		LTIAccessTokenQueries: &lti_query.LTIAccessTokenQueries{DB: db},
		LTILaunchDataQueries:  &lti_query.LTILaunchDataQueries{DB: db},
		UserQueries:           &queries.UserQueries{DB: db},
		UserTokenQueries:      &queries.UserTokenQueries{DB: db},
		PNETServerQueries:     &queries.PNETServerQueries{DB: db},
		RoundQueuePoolQueries: &queries.RoundQueuePoolQueries{DB: db},
		LTIRoutingQueries:     &queries.LTIRoutingQueries{DB: db},
		LTIAttemptQueries:     &queries.LTIAttemptQueries{DB: db},
	}, nil
}

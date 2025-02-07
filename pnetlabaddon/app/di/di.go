package di

import (
	"gitlab.com/a10869/api-modules/pnetlabaddon/platform/database"
)

type DIContainer struct {
}

func NewDIContainer() *DIContainer {
	return &DIContainer{}
}

func (di *DIContainer) Queries() (*database.Queries, error) {
	return database.OpenDBConnection()
}

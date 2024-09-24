package database

import (
	"fmt"
	"time"

	"gitlab.com/a10869/api-modules/backend/pkg/configs"
	"gitlab.com/a10869/api-modules/backend/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

// MysqlConnection func for connection to Mysql database.
func MysqlConnection() (*gorm.DB, error) {

	// Build Mysql connection URL.
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}

	// Define database connection for Mysql.
	db, err := gorm.Open(
		mysql.Open(mysqlConnURL),
		&gorm.Config{
			QueryFields: true,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error, not opened connected to database, %w", err)
	}

	sqlDB.SetMaxIdleConns(configs.AppConfig.DB.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(configs.AppConfig.DB.MaxConnections)
	sqlDB.SetConnMaxLifetime(time.Duration(configs.AppConfig.DB.MaxLifetimeConnections) * time.Minute)
	return db, nil
}

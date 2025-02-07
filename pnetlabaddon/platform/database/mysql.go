package database

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm/schema"

	. "gitlab.com/a10869/api-modules/shared/logs"
	_ "gitlab.com/a10869/api-modules/shared/logs"

	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/configs"
	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
)

var (
	log = NewZeroLogger("db")
)

type DBLogger struct {
}

func (l *DBLogger) Printf(format string, ctx ...interface{}) {
	log.Info().Msg(fmt.Sprintf(format, ctx...))
}

// MysqlConnection func for connection to Mysql database.
func MysqlConnection() (*gorm.DB, error) {

	// Build Mysql connection URL.
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}

	log.Debug().Msg(fmt.Sprintf("Table prefix MysqlConnection %+v", configs.AppConfig.DB.TablePrefix))
	log.Debug().Msg(fmt.Sprintf("Mysql DSN %+v", mysqlConnURL))

	// Define database connection for Mysql.
	db, err := gorm.Open(
		mysql.Open(mysqlConnURL),
		&gorm.Config{
			QueryFields: true,
			Logger: logger.New(
				&DBLogger{},
				logger.Config{
					SlowThreshold:             200 * time.Millisecond,
					LogLevel:                  logger.Warn,
					IgnoreRecordNotFoundError: false,
					Colorful:                  false,
				}),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: configs.AppConfig.DB.TablePrefix,
			},
		},
	)
	if err != nil {
		errText := fmt.Sprintf("error, not connected to database, %+v", err)
		log.Error().Msg(errText)
		return nil, errors.New(errText)
	}

	sqlDB, err := db.DB()
	if err != nil {
		errText := fmt.Sprintf("error, not opened connected to database, %+v", err)
		log.Error().Msg(errText)
		return nil, errors.New(errText)
	}

	sqlDB.SetMaxIdleConns(configs.AppConfig.DB.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(configs.AppConfig.DB.MaxConnections)
	sqlDB.SetConnMaxLifetime(time.Duration(configs.AppConfig.DB.MaxLifetimeConnections) * time.Minute)
	return db, nil
}

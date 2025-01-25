package main

import (
	"fmt"
	"io"
	"os"

	"gitlab.com/a10869/api-modules/backend/app/models"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("mysql").Load(
		&models.LTIAccessToken{},
		&models.LTIForm{},
		&models.LTILaunchData{},
		&models.LTINonceToken{},
		&models.PNETServer{},
		&models.User{},
		&models.UserToken{},
		&models.RoundQueuePool{},
		&models.LTIRouting{},
		&models.LTIAttempt{},
		&models.ServiceCard{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	_, _ = io.WriteString(os.Stdout, stmts)
}

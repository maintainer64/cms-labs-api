package main

import (
	"fmt"
	"io"
	"os"

	"gitlab.com/a10869/api-modules/backend/app/models"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	// usage on `atlas migrate diff --env gorm`
	// usage on `atlas migrate apply --url "$POSTGRES_URL"`
	stmts, err := gormschema.New("mysql").Load(
		&models.LTIForm{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}

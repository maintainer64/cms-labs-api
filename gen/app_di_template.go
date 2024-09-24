package main

const AppDITemplate = `package di

import (
	"github.com/gofiber/fiber/v2"
	"{{.Root}}/app/usecases"
	"{{.Root}}/pkg/utils"
)

func (di *DIContainer) {{.Name}}EditUC() (*usecases.{{.Name}}EditUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.{{.Name}}EditUC{
		{{.Name}}Query: db.{{.Name}}Queries,
	}, nil
}

func (di *DIContainer) {{.Name}}GetUC() (*usecases.{{.Name}}GetUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.{{.Name}}GetUC{
		{{.Name}}Query: db.{{.Name}}Queries,
	}, nil
}

func (di *DIContainer) {{.Name}}ListUC() (*usecases.{{.Name}}ListUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.{{.Name}}ListUC{
		{{.Name}}Query: db.{{.Name}}Queries,
	}, nil
}

func (di *DIContainer) {{.Name}}DeleteUC() (*usecases.{{.Name}}DeleteUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, utils.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.{{.Name}}DeleteUC{
		{{.Name}}Query: db.{{.Name}}Queries,
	}, nil
}

`

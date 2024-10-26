package main

const AppControllersTemplate = `package controllers

import (
	"github.com/gofiber/fiber/v2"
	"{{.Root}}/app/di"
	"{{.Root}}/app/usecases"
	"{{.Root}}/pkg/utils"
)

// {{.Name}}Create func for creates a new {{.Name}}.
// @Description Create {{.NameSnake}}.
// @Summary create {{.NameSnake}}
// @Tags {{.Name}}
// @Accept json
// @Produce json
// @Param form body usecases.{{.Name}}EditInputDTO true "{{.NameSnake}} form info"
// @Success 200 {object} usecases.{{.Name}}EditResponse
// @Security ApiKeyAuth
// @Router /v1/{{.NameDash}}/upsert [post]
func {{.Name}}Create(c *fiber.Ctx) error {
	dto := usecases.{{.Name}}EditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().{{.Name}}EditUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// {{.Name}}List func for view of list {{.Name}}.
// @Description List {{.NameSnake}}.
// @Summary list {{.NameSnake}}
// @Tags {{.Name}}
// @Accept json
// @Produce json
// @Param form body usecases.{{.Name}}ListInputDTO true "{{.NameSnake}} list info"
// @Success 200 {object} usecases.{{.Name}}ListResponse
// @Security ApiKeyAuth
// @Router /v1/{{.NameDash}}/list [post]
func {{.Name}}List(c *fiber.Ctx) error {
	dto := usecases.{{.Name}}ListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().{{.Name}}ListUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// {{.Name}}Delete func for delete {{.Name}}.
// @Description Delete {{.NameSnake}}.
// @Summary delete {{.NameSnake}}
// @Tags {{.Name}}
// @Accept json
// @Produce json
// @Param form body usecases.{{.Name}}DeleteInputDTO true "{{.NameSnake}} id"
// @Success 200 {object} usecases.{{.Name}}DeleteResponse
// @Security ApiKeyAuth
// @Router /v1/{{.NameDash}}/delete [post]
func {{.Name}}Delete(c *fiber.Ctx) error {
	dto := usecases.{{.Name}}DeleteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().{{.Name}}DeleteUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// {{.Name}}Get func for full model {{.Name}}.
// @Description get {{.NameSnake}}.
// @Summary get {{.NameSnake}}
// @Tags {{.Name}}
// @Accept json
// @Produce json
// @Param form body usecases.{{.Name}}GetInputDTO true "{{.NameSnake}} id"
// @Success 200 {object} usecases.{{.Name}}GetResponse
// @Security ApiKeyAuth
// @Router /v1/{{.NameDash}}/get [post]
func {{.Name}}Get(c *fiber.Ctx) error {
	dto := usecases.{{.Name}}GetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().{{.Name}}GetUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

`

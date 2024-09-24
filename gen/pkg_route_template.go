package main

const PKGRouteTemplate = `package routes

import (
	"github.com/gofiber/fiber/v2"
	"{{.Root}}/app/controllers"
)

func V1{{.Name}}Routes(a *fiber.App) {
	group := a.Group("/api/v1/{{.NameDash}}")
	group.Post("/upsert", controllers.{{.Name}}Create)
	group.Post("/list", controllers.{{.Name}}List)
	group.Post("/delete", controllers.{{.Name}}Delete)
	group.Post("/get", controllers.{{.Name}}Get)
}

`

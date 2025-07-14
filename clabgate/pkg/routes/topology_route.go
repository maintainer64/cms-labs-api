package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/clabgate/app/controllers"
)

func V1TopologyRoute(a *fiber.App) {
	group := a.Group("/clabgate/api/v1/topologies")
	group.Post("/get", controllers.TopologiesGet)
	group.Post("/create", controllers.TopologiesCreate)
	group.Post("/delete", controllers.TopologiesDelete)
}

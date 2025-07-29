package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/clabgate/app/controllers"
)

func V1ContainersRoute(a *fiber.App) {
	group := a.Group("/clabgate/api/v1/containers")
	group.Post("/get", controllers.ContainersDeviceGet)
}

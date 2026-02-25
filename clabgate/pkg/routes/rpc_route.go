package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/clabgate/app/controllers"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
)

// V1RpcRoute func for describe group of jsonrpc 2.0 protocol.
func V1RpcRoute(a *fiber.App) {
	rpc := jsonrpc.NewJsonRPCServer("/clabgate/api/v1/rpc", a)
	rpc.Method("task.list", controllers.TaskList)
	rpc.Method("topology.get", controllers.TopologyGet)
	rpc.Method("topology.create", controllers.TopologyCreate)
	rpc.Method("topology.delete", controllers.TopologyDelete)
	rpc.Method("container.get", controllers.ContainerGet)
}

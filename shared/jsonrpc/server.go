package jsonrpc

import "github.com/gofiber/fiber/v2"

type Server struct {
	Fiber   *fiber.App
	RpcPath string
	Methods map[string]Handler
}

func NewJsonRPCServer(rpcPath string, app *fiber.App) *Server {
	server := &Server{
		Fiber:   app,
		RpcPath: rpcPath,
		Methods: map[string]Handler{},
	}
	server.Fiber.Use([]string{server.RpcPath, server.RpcPath + "/*"}, server.requestProcessing)
	return server
}

func (s *Server) Method(name string, handler Handler) {
	s.Methods[name] = handler
}

func (s *Server) requestProcessing(c *fiber.Ctx) error {
	req := RequestJSONRPC{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"jsonrpc": "2.0",
			"error":   fiber.Map{"code": ParseError, "message": "Parse error"},
			"id":      nil,
		})
	}
	handler, exists := s.Methods[req.Method]
	if !exists {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"jsonrpc": "2.0",
			"error":   fiber.Map{"code": MethodNotFound, "message": "Method not found"},
			"id":      req.ID,
		})
	}
	result, err := handler(&Ctx{
		FiberCtx: c,
		Params:   req.Params,
	})
	switch e := err.(type) {
	case RpcError:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"jsonrpc": "2.0",
			"error": fiber.Map{
				"code":    RPCLogicError,
				"message": "Validation error",
				"data": fiber.Map{
					"code":    e.Code,
					"message": e.Message,
				},
			},
			"id": req.ID,
		})
	case RpcValidatorError:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"jsonrpc": "2.0",
			"error": fiber.Map{
				"code":    InvalidParams,
				"message": e.Exception.Error(),
			},
			"id": req.ID,
		})
	case nil:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"jsonrpc": "2.0",
			"result":  result,
			"id":      req.ID,
		})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"jsonrpc": "2.0",
			"error":   fiber.Map{"code": InternalError, "message": e.Error()},
			"id":      req.ID,
		})
	}
}

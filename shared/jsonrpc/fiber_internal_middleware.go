package jsonrpc

import (
	fiber "github.com/gofiber/fiber/v2"
)

func InternalFormatterNew(debug bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
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
				"id": "1",
			})
		case RpcValidatorError:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"jsonrpc": "2.0",
				"error": fiber.Map{
					"code":    InvalidParams,
					"message": e.Exception.Error(),
				},
				"id": "1",
			})
		case nil:
			return e
		default:
			if debug {
				return e
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"jsonrpc": "2.0",
				"error": fiber.Map{
					"code":    InternalError,
					"message": "Internal Server Error",
				},
				"id": "1",
			})
		}
	}
}

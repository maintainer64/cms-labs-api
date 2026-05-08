// Package logs общие компоненты для Fiber логирования
package logs

import (
	"strings"

	"github.com/gofiber/contrib/fiberzerolog"
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
)

func NewFiberZerologLogger() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		c := &jsonrpc.Ctx{FiberCtx: ctx}
		logger := NewZeroLogger(NewZeroLoggerConf(c))
		return fiberzerolog.New(
			fiberzerolog.Config{
				Logger: logger,
				Next: func(c *fiber.Ctx) bool {
					path := c.Path()
					return strings.HasPrefix(path, "/api/docs") ||
						strings.HasPrefix(path, "/clabgate/api/docs")
				},
			},
		)(c.FiberCtx)
	}
}

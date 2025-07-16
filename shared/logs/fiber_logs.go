// Package logs общие компоненты для Fiber логирования
package logs

import (
	"github.com/gofiber/contrib/fiberzerolog"
	fiber "github.com/gofiber/fiber/v2"
)

func NewFiberZerologLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := NewZeroLogger(NewZeroLoggerConf(c))
		return fiberzerolog.New(
			fiberzerolog.Config{
				Logger: logger,
			},
		)(c)
	}
}

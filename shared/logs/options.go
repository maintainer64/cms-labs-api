// Package logs общие компоненты для Fiber логирования
package logs

import (
	"github.com/gofiber/fiber/v2"
)

// NewOptionsFiberNoContent Функция отправляет 204 на все OPTIONS запросы для браузера
func NewOptionsFiberNoContent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() != "OPTIONS" {
			// Продолжаем обработку запроса, если это не OPTIONS
			return c.Next()
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

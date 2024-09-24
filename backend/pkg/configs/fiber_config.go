package configs

import (
	"time"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(AppConfig.Server.ServerReadTimeout),
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
}

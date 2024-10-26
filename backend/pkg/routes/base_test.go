package routes

import (
	"fmt"
	"io"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gitlab.com/a10869/api-modules/backend/pkg/middleware"
	"gitlab.com/a10869/api-modules/backend/platform/database"
	"gorm.io/gorm"
)

func FiberAppTest() (*fiber.App, *gorm.DB) {
	// Load .env.test file from the root folder.
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}

	// Define a new Fiber app.
	app := fiber.New()
	middleware.FiberMiddleware(app)
	db, _ := database.MysqlConnection()
	return app, db
}

func FiberRequestPayload(payload any) io.Reader {
	return strings.NewReader(fmt.Sprint(payload))
}

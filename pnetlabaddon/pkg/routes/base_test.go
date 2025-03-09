package routes

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/h2non/gock"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/di"

	"gitlab.com/a10869/api-modules/shared/logs"

	"github.com/goccy/go-json"

	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/middleware"
	"gitlab.com/a10869/api-modules/pnetlabaddon/platform/database"
	"gorm.io/gorm"
)

type FiberTestHTTP struct {
	App *fiber.App
	DB  *gorm.DB
}

func (f *FiberTestHTTP) Request(method string, route string, body io.Reader, authorization string) (int, string) {
	req := httptest.NewRequest(method, route, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)

	// Perform the request plain with the app.
	resp, _ := f.App.Test(req, -1)
	responseString := ""
	if b, err := io.ReadAll(resp.Body); err == nil {
		responseString = string(b)
	}
	return resp.StatusCode, responseString
}

func NewFiberTestHTTP() *FiberTestHTTP {
	// Load .env.test file from the root folder.
	_ = godotenv.Load("../../.env")
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}
	// Load .env local file form the additional
	configs.AppConfig.Reload()

	// Inject Resty
	di.NewRestyClient = func() *resty.Client {
		return resty.NewWithClient(
			&http.Client{Transport: gock.DefaultTransport},
		)
	}

	// Define a new Fiber app.
	app := fiber.New()
	middleware.FiberMiddleware(app)
	FiberRoutes(app)
	db, _ := database.MysqlConnection(logs.NewZeroLogger(&logs.ZeroLoggerConf{}))
	return &FiberTestHTTP{
		App: app,
		DB:  db,
	}
}

func FiberJSON(data any) string {
	jsonByte, _ := json.Marshal(data)
	return string(jsonByte)
}

func FiberRequestPayload(data any) io.Reader {
	return strings.NewReader(fmt.Sprint(FiberJSON(data)))
}

package routes

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/a10869/api-modules/shared/connection"

	resty "github.com/go-resty/resty/v2"
	"github.com/h2non/gock"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/di"

	"gitlab.com/a10869/api-modules/shared/logs"

	json "github.com/goccy/go-json"

	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/configs"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gitlab.com/a10869/api-modules/pnetlabaddon/pkg/middleware"
	"gorm.io/gorm"
)

type FiberTestHTTP struct {
	App *fiber.App
	DB  *gorm.DB
}

type FiberTestHttpRequest struct {
	Method        string
	Route         string
	Body          io.Reader
	Authorization string
	ContentType   string
}

func (f *FiberTestHTTP) Request(
	r *FiberTestHttpRequest,
) (int, string) {
	req := httptest.NewRequest(r.Method, r.Route, r.Body)
	if r.ContentType == "" {
		r.ContentType = "application/json"
	}
	req.Header.Set("Content-Type", r.ContentType)
	req.Header.Set("Authorization", r.Authorization)

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
	// Fix relative paths for config files - find project root
	projectRoot := findPnetlabTestProjectRoot()
	if projectRoot != "" {
		if path := os.Getenv("ADDONS_PATH_CONFIG"); path != "" && !filepath.IsAbs(path) {
			os.Setenv("ADDONS_PATH_CONFIG", filepath.Join(projectRoot, path))
		}
		if path := os.Getenv("PROXMOX_PATH_CONFIG"); path != "" && !filepath.IsAbs(path) {
			os.Setenv("PROXMOX_PATH_CONFIG", filepath.Join(projectRoot, path))
		}
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
	db, _ := connection.MysqlConnection(configs.AppConfig.DB, logs.NewZeroLogger(&logs.ZeroLoggerConf{}))
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

func findPnetlabTestProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		goModPath := filepath.Join(dir, "go.mod")
		// #nosec G304
		if _, err := os.ReadFile(goModPath); err == nil {
			if isPnetlabTestProjectRoot(dir) {
				return dir
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

func isPnetlabTestProjectRoot(dir string) bool {
	// #nosec G304
	data, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil {
		return false
	}
	content := string(data)
	return !strings.Contains(content, "/backend") &&
		!strings.Contains(content, "/shared") &&
		!strings.Contains(content, "/pnetlabaddon") &&
		!strings.Contains(content, "/clabgate")
}

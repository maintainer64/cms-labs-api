package routes

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http/httptest"
	"net/url"
	"strings"

	"gitlab.com/a10869/api-modules/shared/logs"

	"github.com/goccy/go-json"

	"github.com/google/uuid"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/pkg/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gitlab.com/a10869/api-modules/backend/pkg/middleware"
	"gitlab.com/a10869/api-modules/backend/platform/database"
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

func (f *FiberTestHTTP) AuthorizationUser(userID uint, serverID uint) string {
	if userID == 0 {
		entity := models.User{}
		entity.Email = uuid.New().String() + "@admin.com"
		entity.UserRole = models.UsersRoleAdmin
		f.DB.Create(&entity)
		userID = entity.ID
	}
	container, _ := di.NewDIContainer(&logs.ZeroLoggerConf{})
	uc := container.AuthTokenManager()
	token, _ := uc.NewJWTByUserId("", userID, serverID, nil)
	return "Bearer " + token.AccessToken
}

func (f *FiberTestHTTP) AuthorizationServiceBasic() (string, string) {
	entity := models.PNETServer{}
	entity.Type = models.ServerTypeOpenID
	entity.Name = uuid.New().String() + "_server"
	entity.Url = "https://localhost"
	entity.IsActive = true
	entity.Token = uuid.New().String()
	entity.ClientID = uuid.New().String()
	f.DB.Create(&entity)
	return "Basic " + base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", entity.ClientID, entity.Token)),
	), entity.ClientID
}

func NewFiberTestHTTP() *FiberTestHTTP {
	// Load .env.test file from the root folder.
	_ = godotenv.Load("../../.env")
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}
	// Load .env local file form the additional
	configs.AppConfig.Reload()

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

func FiberRequestFormPayload(data map[string]string) io.Reader {
	// Создаем form-data запрос
	form := url.Values{}
	for key, value := range data {
		form.Add(key, value)
	}
	return strings.NewReader(form.Encode())
}

package routes

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/external"
)

func TestCurlRequestCreate(t *testing.T) {
	description := "Create new CurlRequest"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)

	// Тестовые данные
	dto := usecases.CurlRequestEditInputDTO{
		Name:       "Test Request " + uuid.New().String(),
		URL:        "https://example.com/api",
		Method:     "POST",
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       `{"key":"value"}`,
		RawRequest: `POST /api HTTP/1.1\nContent-Type: application/json\n\n{"key":"value"}`,
	}

	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/curl-request/upsert",
		Body:          FiberRequestPayload(dto),
		Authorization: authHeader,
	})

	bodyModel := usecases.CurlRequestEditResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, 200, statusCode, description)
	assert.Greater(t, bodyModel.Result.ID, uint(0), description)

	// Проверка, что запрос действительно создан в базе данных
	var createdEntity models.CurlRequest
	f.DB.First(&createdEntity, bodyModel.Result.ID)
	assert.Equal(t, dto.Name, createdEntity.Name, description)
	assert.Equal(t, dto.URL, createdEntity.URL, description)
	assert.Equal(t, dto.Method, createdEntity.Method, description)
}

func TestCurlRequestGet(t *testing.T) {
	description := "get CurlRequest"
	f := NewFiberTestHTTP()

	// Создаем тестовый запрос
	entity := models.CurlRequest{
		CurlRequestBase: models.CurlRequestBase{
			Name: "Test Request " + uuid.New().String(),
			URL:  "https://example.com/api",
		},
		CurlRequestSecret: models.CurlRequestSecret{
			Method:  "GET",
			Headers: map[string]string{"Authorization": "Bearer token"},
			Body:    "",
			Raw:     "GET /api HTTP/1.1\nAuthorization: Bearer token",
		},
	}
	f.DB.Create(&entity)

	authHeader := f.AuthorizationUser(0, 0)
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "POST",
		Route:  "/api/v1/curl-request/get",
		Body: FiberRequestPayload(map[string]any{
			"id": entity.ID,
		}),
		Authorization: authHeader,
	})

	bodyModel := usecases.CurlRequestGetResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, 200, statusCode, description)
	assert.Equal(t, entity.ID, bodyModel.Result.Model.ID, description)
	assert.Equal(t, entity.Name, bodyModel.Result.Model.Name, description)
	assert.Equal(t, entity.URL, bodyModel.Result.Model.URL, description)
}

func TestCurlRequestGetNotFound(t *testing.T) {
	description := "not found CurlRequest"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)
	expectedBody := map[string]interface{}{
		"error": true,
		"msg":   "CurlRequest not found",
	}

	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "POST",
		Route:  "/api/v1/curl-request/get",
		Body: FiberRequestPayload(map[string]any{
			"id": 9999999,
		}),
		Authorization: authHeader,
	})

	assert.Equal(t, 404, statusCode, description)
	assert.Equal(t, FiberJSON(expectedBody), body, description)
}

func TestCurlRequestList(t *testing.T) {
	f := NewFiberTestHTTP()

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.CurlRequest{})

	// Создаем тестовые запросы
	entity1 := models.CurlRequest{
		CurlRequestBase: models.CurlRequestBase{
			Name: "Request A " + uuid.New().String(),
			URL:  "https://example.com/a",
		},
	}
	f.DB.Create(&entity1)

	entity2 := models.CurlRequest{
		CurlRequestBase: models.CurlRequestBase{
			Name: "Request B " + uuid.New().String(),
			URL:  "https://example.com/b",
		},
	}
	f.DB.Create(&entity2)

	entity3 := models.CurlRequest{
		CurlRequestBase: models.CurlRequestBase{
			Name: "Another Request " + uuid.New().String(),
			URL:  "https://another.com",
		},
	}
	f.DB.Create(&entity3)

	authHeader := f.AuthorizationUser(0, 0)

	tests := []struct {
		description string
		body        map[string]any
		ids         []uint
	}{
		{
			description: "List all requests",
			body: map[string]any{
				"limit":  100,
				"offset": 0,
			},
			ids: []uint{entity3.ID, entity2.ID, entity1.ID}, // Ожидаем в обратном порядке created_at
		},
		{
			description: "Search by name",
			body: map[string]any{
				"search": "Request A",
				"limit":  100,
				"offset": 0,
			},
			ids: []uint{entity1.ID},
		},
		{
			description: "Search by URL",
			body: map[string]any{
				"search": "another.com",
				"limit":  100,
				"offset": 0,
			},
			ids: []uint{entity3.ID},
		},
		{
			description: "Pagination",
			body: map[string]any{
				"limit":  1,
				"offset": 1,
			},
			ids: []uint{entity2.ID},
		},
	}

	for _, test := range tests {
		statusCode, body := f.Request(&FiberTestHttpRequest{
			Method:        "POST",
			Route:         "/api/v1/curl-request/list",
			Body:          FiberRequestPayload(test.body),
			Authorization: authHeader,
		})

		bodyModel := usecases.CurlRequestListResponse{}
		_ = json.Unmarshal([]byte(body), &bodyModel)

		ids := make([]uint, 0)
		for _, model := range bodyModel.Result.Model {
			ids = append(ids, model.ID)
		}

		assert.Equal(t, 200, statusCode, test.description)
		assert.Equal(t, test.ids, ids, test.description)
	}
}

func TestCurlRequestDelete(t *testing.T) {
	f := NewFiberTestHTTP()

	// Создаем тестовый запрос
	entity := models.CurlRequest{
		CurlRequestBase: models.CurlRequestBase{
			Name: "Request to delete " + uuid.New().String(),
			URL:  "https://example.com/delete",
		},
	}
	f.DB.Create(&entity)

	authHeader := f.AuthorizationUser(0, 0)

	tests := []struct {
		description string
		id          uint
	}{
		{
			description: "Successfully delete CurlRequest",
			id:          entity.ID,
		},
		{
			description: "Fail to delete non-existent CurlRequest",
			id:          999999,
		},
	}

	for _, test := range tests {
		statusCode, body := f.Request(&FiberTestHttpRequest{
			Method: "POST",
			Route:  "/api/v1/curl-request/delete",
			Body: FiberRequestPayload(map[string]any{
				"id": test.id,
			}),
			Authorization: authHeader,
		})

		response := usecases.CurlRequestDeleteResponse{}
		_ = json.Unmarshal([]byte(body), &response)
		assert.Equal(t, 200, statusCode, test.description)
		assert.Equal(t, test.id, response.Result.ID, test.description)

		// Проверяем, что запрос удален из базы данных
		var deletedEntity models.CurlRequest
		f.DB.First(&deletedEntity, test.id)
		assert.Equal(t, uint(0), deletedEntity.ID, test.description)
	}
}

func TestCurlRequestUnauthorizedAccess(t *testing.T) {
	f := NewFiberTestHTTP()

	tests := []struct {
		description  string
		method       string
		route        string
		body         map[string]any
		expectedCode int
	}{
		{
			description:  "Unauthorized access to create",
			method:       "POST",
			route:        "/api/v1/curl-request/upsert",
			body:         map[string]any{"name": "Test", "url": "https://test.com"},
			expectedCode: 401,
		},
		{
			description:  "Unauthorized access to get",
			method:       "POST",
			route:        "/api/v1/curl-request/get",
			body:         map[string]any{"id": 1},
			expectedCode: 401,
		},
		{
			description:  "Unauthorized access to list",
			method:       "POST",
			route:        "/api/v1/curl-request/list",
			body:         map[string]any{},
			expectedCode: 401,
		},
		{
			description:  "Unauthorized access to delete",
			method:       "POST",
			route:        "/api/v1/curl-request/delete",
			body:         map[string]any{"id": 1},
			expectedCode: 401,
		},
	}

	for _, test := range tests {
		statusCode, _ := f.Request(&FiberTestHttpRequest{
			Method: test.method,
			Route:  test.route,
			Body:   FiberRequestPayload(test.body),
		})

		assert.Equal(t, test.expectedCode, statusCode, test.description)
	}
}

func TestApplyOverrides(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		overrides map[string]string
		expected  string
	}{
		{
			name:      "Simple replacement",
			input:     "https://api.example.com/users/$userId",
			overrides: map[string]string{"userId": "123"},
			expected:  "https://api.example.com/users/123",
		},
		{
			name:      "Multiple replacements",
			input:     "Bearer $token, User: $userId",
			overrides: map[string]string{"token": "abc123", "userId": "456"},
			expected:  "Bearer abc123, User: 456",
		},
		{
			name:      "No placeholders",
			input:     "Just a normal string",
			overrides: map[string]string{"key": "value"},
			expected:  "Just a normal string",
		},
		{
			name:      "Placeholder without override",
			input:     "Keep $undefined as is",
			overrides: map[string]string{"other": "value"},
			expected:  "Keep $undefined as is",
		},
		{
			name:      "Special characters in replacement",
			input:     "Replace $key",
			overrides: map[string]string{"key": "a/b?c=d&e=f"},
			expected:  "Replace a/b?c=d&e=f",
		},
		{
			name:      "Multiple same placeholders",
			input:     "$key $key $key",
			overrides: map[string]string{"key": "value"},
			expected:  "value value value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := external.ApplyOverrides(tt.input, tt.overrides)
			assert.Equal(t, tt.expected, result, "Test case: %s", tt.name)
		})
	}
}

func TestCurlRequestExecute(t *testing.T) {
	// Инициализация тестового приложения
	f := NewFiberTestHTTP()

	authHeader, _ := f.AuthorizationServiceBasic()

	// Создаем тестовый запрос в БД
	testRequest := models.CurlRequest{
		CurlRequestBase: models.CurlRequestBase{
			Name: "Test Request",
			URL:  "https://api.example.com/users/$userId",
		},
		CurlRequestSecret: models.CurlRequestSecret{
			Method:  "GET",
			Headers: map[string]string{"Authorization": "Bearer $token"},
			Timeout: 10,
		},
	}
	f.DB.Create(&testRequest)

	// Мок внешнего API
	defer gock.Off()
	gock.New("https://api.example.com").
		MatchHeader("Authorization", "Bearer valid_token_123").
		Get("/users/42").
		Reply(200).
		JSON(map[string]interface{}{"id": 42, "name": "John Doe"})

	tests := []struct {
		name           string
		payload        external.CurlRequestExecuteInputDTO
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Success with overrides",
			payload: external.CurlRequestExecuteInputDTO{
				CurlRequestID: testRequest.ID,
				Overrides: map[string]string{
					"userId": "42",
					"token":  "valid_token_123",
				},
			},
			expectedStatus: 200,
			expectedBody: map[string]interface{}{
				"status":  200,
				"headers": map[string]interface{}{"Content-Type": "application/json"},
				"body":    map[string]interface{}{"id": float64(42), "name": "John Doe"},
			},
		},
		{
			name: "Missing required override",
			payload: external.CurlRequestExecuteInputDTO{
				CurlRequestID: testRequest.ID,
				Overrides: map[string]string{
					"userId": "42",
					// token не указан
				},
			},
			expectedStatus: 500,
			expectedBody: map[string]interface{}{
				"error": true,
				"msg":   "cannot match any request",
			},
		},
		{
			name: "Invalid request ID",
			payload: external.CurlRequestExecuteInputDTO{
				CurlRequestID: 999999, // несуществующий ID
				Overrides: map[string]string{
					"userId": "42",
					"token":  "valid_token_123",
				},
			},
			expectedStatus: 500,
			expectedBody: map[string]interface{}{
				"error": true,
				"msg":   "CurlRequest not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Подготовка запроса
			statusCode, body := f.Request(&FiberTestHttpRequest{
				Method:        "POST",
				Route:         "/api/v1/curl-request/execute",
				Body:          FiberRequestPayload(tt.payload),
				Authorization: authHeader,
			})

			// Проверка статуса
			assert.Equal(t, tt.expectedStatus, statusCode)

			if tt.expectedStatus == 200 {
				// Для успешного случая проверяем структуру ответа
				assert.Equal(t, tt.expectedBody["status"], statusCode)

				// Проверяем body
				expectedBody := tt.expectedBody["body"].(map[string]interface{})
				actualBody := make(map[string]interface{})
				_ = json.Unmarshal([]byte(body), &actualBody)
				assert.Equal(t, expectedBody, actualBody)
			} else {
				// Для ошибок проверяем наличие полей error и msg
				// Проверяем body
				actualBody := make(map[string]interface{})
				_ = json.Unmarshal([]byte(body), &actualBody)
				assert.Equal(t, tt.expectedBody["error"], actualBody["error"])
				assert.Contains(t, actualBody["msg"], tt.expectedBody["msg"])
			}
		})
	}

	// Проверяем что все моки были вызваны
	assert.True(t, gock.IsDone())
}

package routes

import (
	"testing"
	"time"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"

	"github.com/google/uuid"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/external"

	"gitlab.com/a10869/api-modules/backend/app/models"
)

func TestV1PNETServerRouteGet(t *testing.T) {
	description := "get server"
	f := NewFiberTestHTTP()
	entity := models.PNETServer{}
	entity.Name = "Server"
	entity.Url = "https://localhost"
	entity.Type = models.ServerTypePnet
	entity.IsActive = true
	entity.MinutesForDisconnect = 10
	entity.MaxCountUsersLimit = 20
	entity.UnitRate = 1
	entity.Token = uuid.New().String()
	entity.ClientID = uuid.New().String()
	entity.CreatedAt = time.Now().UTC()
	entity.UpdatedAt = time.Now().UTC()
	f.DB.Create(&entity)
	authHeader := f.AuthorizationUser(0, 0)
	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "POST",
		Route:  "/api/v1/pnet-server/get",
		Body: FiberRequestPayload(map[string]any{
			"id": entity.ID,
		}),
		Authorization: authHeader,
	})
	bodyModel := usecases.PNETServerGetResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, entity.ID, bodyModel.Result.Model.ID, description)
}

func TestV1PNETServerRouteGetNotFound(t *testing.T) {
	description := "not found pnet"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)
	expectedBody := map[string]interface{}{
		"error": true,
		"msg":   "PNETServer not found",
	}
	expectedCode := 404
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "POST",
		Route:  "/api/v1/pnet-server/get",
		Body: FiberRequestPayload(map[string]any{
			"id": 9999999,
		}),
		Authorization: authHeader,
	})

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, FiberJSON(expectedBody), body, description)
}

func TestV1PNETServerRouteSearch(t *testing.T) {
	f := NewFiberTestHTTP()
	now := time.Now().UTC()
	entity := models.PNETServer{}
	entity.Name = "Server " + uuid.New().String()
	entity.Url = "https://localhost"
	entity.Type = models.ServerTypePnet
	entity.IsActive = true
	entity.MinutesForDisconnect = 10
	entity.MaxCountUsersLimit = 20
	entity.LastOnlineStatus = &now
	entity.UnitRate = 20
	entity.Token = uuid.New().String()
	entity.ClientID = uuid.New().String()
	f.DB.Create(&entity)
	entity2 := models.PNETServer{}
	entity2.Name = entity.Name
	entity2.UnitRate = 10
	entity2.Url = "https://anotherhost"
	entity2.Type = models.ServerTypeOpenID
	entity2.IsActive = false
	entity2.Token = uuid.New().String()
	entity2.ClientID = uuid.New().String()
	f.DB.Create(&entity2)
	authHeader := f.AuthorizationUser(0, 0)
	tests := []struct {
		description string
		body        map[string]any
		ids         []uint
	}{
		{
			description: "Search entity order from ALL",
			body: map[string]any{
				"search":   entity.Name,
				"limit":    100,
				"offset":   0,
				"status":   "all",
				"order_by": "createdAt",
			},
			ids: []uint{entity2.ID, entity.ID},
		},
		{
			description: "Search entity 2 from ALL",
			body: map[string]any{
				"search":   entity.Name,
				"limit":    1,
				"offset":   0,
				"status":   "all",
				"order_by": "createdAt",
			},
			ids: []uint{entity2.ID},
		},
		{
			description: "Search entity 1 from ALL",
			body: map[string]any{
				"search":   entity.Name,
				"limit":    1,
				"offset":   1,
				"status":   "all",
				"order_by": "createdAt",
			},
			ids: []uint{entity.ID},
		},
		{
			description: "Search entity 1 from active",
			body: map[string]any{
				"search":   entity.Name,
				"limit":    100,
				"offset":   0,
				"status":   "active",
				"order_by": "createdAt",
			},
			ids: []uint{entity.ID},
		},
		{
			description: "Order from unitRate",
			body: map[string]any{
				"search":   entity.Name,
				"limit":    100,
				"offset":   0,
				"status":   "all",
				"order_by": "unitRate",
			},
			ids: []uint{entity.ID, entity2.ID},
		},
	}
	for _, test := range tests {
		statusCode, body := f.Request(&FiberTestHttpRequest{
			Method:        "POST",
			Route:         "/api/v1/pnet-server/list",
			Body:          FiberRequestPayload(test.body),
			Authorization: authHeader,
		})
		bodyModel := usecases.PNETServerListResponse{}
		_ = json.Unmarshal([]byte(body), &bodyModel)
		ids := make([]uint, 0)
		for _, model := range bodyModel.Result.Model {
			ids = append(ids, model.ID)
		}
		assert.Equal(t, statusCode, 200, test.description)
		assert.Equal(t, test.ids, ids, test.description)
	}
}

func TestV1PNETServerRouteDelete(t *testing.T) {
	f := NewFiberTestHTTP()
	now := time.Now().UTC()
	entity := models.PNETServer{}
	entity.Name = "Server " + uuid.New().String()
	entity.Url = "https://localhost"
	entity.Type = models.ServerTypePnet
	entity.IsActive = true
	entity.MinutesForDisconnect = 10
	entity.MaxCountUsersLimit = 20
	entity.LastOnlineStatus = &now
	entity.UnitRate = 20
	entity.Token = uuid.New().String()
	entity.ClientID = uuid.New().String()
	f.DB.Create(&entity)

	authHeader := f.AuthorizationUser(0, 0)
	tests := []struct {
		description string
		id          uint
		statusCode  int
		expectedErr string
	}{
		{
			description: "Successfully delete PNETServer",
			id:          entity.ID,
			statusCode:  200,
		},
		{
			description: "Fail to delete non-existent PNETServer",
			id:          999999,
		},
	}

	for _, test := range tests {
		statusCode, body := f.Request(&FiberTestHttpRequest{
			Method: "POST",
			Route:  "/api/v1/pnet-server/delete",
			Body: FiberRequestPayload(map[string]any{
				"id": test.id,
			}),
			Authorization: authHeader,
		})
		response := usecases.PNETServerDeleteResponse{}
		_ = json.Unmarshal([]byte(body), &response)
		assert.Equal(t, 200, statusCode, test.description)
		assert.Equal(t, test.id, response.Result.ID, test.description)

		// Проверяем, что сервер успешно удален
		var deletedEntity models.PNETServer
		f.DB.First(&deletedEntity, test.id)
		var idNotFound uint = 0
		assert.Equal(t, idNotFound, deletedEntity.ID, test.description)
	}
}

func TestV1PNETServerRouteCreate(t *testing.T) {
	description := "Create new PNET server"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)

	// Тестовые данные
	now := time.Now().UTC()
	entity := models.PNETServer{}

	entity.Name = "Server " + uuid.New().String()
	entity.Url = "https://localhost"
	entity.Type = models.ServerTypePnet
	entity.IsActive = true
	entity.MinutesForDisconnect = 10
	entity.MaxCountUsersLimit = 20
	entity.LastOnlineStatus = &now
	entity.UnitRate = 20
	entity.ClientID = uuid.New().String()

	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "POST",
		Route:  "/api/v1/pnet-server/upsert",
		Body: FiberRequestPayload(usecases.PNETServerEditInputDTO{
			ClientID:             entity.ClientID,
			Type:                 entity.Type,
			Name:                 entity.Name,
			Url:                  entity.Url,
			IsActive:             entity.IsActive,
			MinutesForDisconnect: entity.MinutesForDisconnect,
			MaxCountUsersLimit:   entity.MaxCountUsersLimit,
			UnitRate:             entity.UnitRate,
		}),
		Authorization: authHeader,
	})

	bodyModel := usecases.PNETServerEditResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, statusCode, 200, description)
	assert.Equal(t, true, bodyModel.Result.ID > 0, description)

	// Проверка, что сервер действительно создан в базе данных
	var createdEntity models.PNETServer
	f.DB.First(&createdEntity, bodyModel.Result.ID)
	assert.Equal(t, entity.Name, createdEntity.Name, description)
	assert.Equal(t, entity.Url, createdEntity.Url, description)
	assert.Equal(t, entity.Type, createdEntity.Type, description)
	assert.Equal(t, entity.IsActive, createdEntity.IsActive, description)
	assert.Equal(t, entity.MinutesForDisconnect, createdEntity.MinutesForDisconnect, description)
	assert.Equal(t, entity.MaxCountUsersLimit, createdEntity.MaxCountUsersLimit, description)
	assert.Equal(t, entity.UnitRate, createdEntity.UnitRate, description)
}

func TestV1PNETServerPingSuccess(t *testing.T) {
	description := "ping PNET server successfully"
	f := NewFiberTestHTTP()

	// Create a test PNET server
	authHeader, clientID := f.AuthorizationServiceBasic()
	var serverEntity models.PNETServer

	f.DB.Where("client_id = ?", clientID).Find(&serverEntity)

	userEntity := models.User{}
	userEntity.Email = uuid.New().String() + "@example.com"
	userEntity.UserRole = models.UsersRoleStudent
	f.DB.Create(&userEntity)

	ltiRoutingEntity := models.LTIRouting{}
	ltiRoutingEntity.LTITitle = uuid.New().String() + "@example.com"
	f.DB.Create(&ltiRoutingEntity)

	// Create a test LTI attempt
	attemptEntity := models.LTIAttempt{}
	attemptEntity.AttemptID = uuid.New().String()
	attemptEntity.UserID = userEntity.ID
	attemptEntity.LTIRoutingID = ltiRoutingEntity.ID
	attemptEntity.PNETServerID = serverEntity.ID
	attemptEntity.ExpiredAt = time.Now().UTC().Add(30 * time.Minute)
	f.DB.Create(&attemptEntity)

	input := external.PNETServerPingInputDTO{
		Attempts: []external.AttemptDTO{
			{
				AttemptID: attemptEntity.AttemptID,
			},
		},
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/pnet-server/ping",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := response.Response[external.PNETServerPingOutputDTO]{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, len(input.Attempts), bodyModel.Result.Count, description)

	// Verify that the server's last online status and count were updated
	var updatedServer models.PNETServer
	err := f.DB.First(&updatedServer, serverEntity.ID).Error
	assert.NoError(t, err, description)
	assert.NotNil(t, updatedServer.LastOnlineStatus, description)
	assert.Equal(t, len(input.Attempts), updatedServer.LastCountUsers, description)

	// Verify that the attempt's expired_at was extended
	var updatedAttempt models.LTIAttempt
	err = f.DB.First(&updatedAttempt, attemptEntity.ID).Error
	assert.NoError(t, err, description)
	assert.True(t, updatedAttempt.ExpiredAt.After(attemptEntity.ExpiredAt), description)
}

func TestV1PNETServerPingUnauthorized(t *testing.T) {
	description := "ping PNET server unauthorized"
	f := NewFiberTestHTTP()

	input := external.PNETServerPingInputDTO{
		Attempts: []external.AttemptDTO{
			{
				AttemptID: "test-attempt-id",
				UserEmail: "user@example.com",
				UserID:    1,
			},
		},
	}

	expectedCode := 401 // Assuming 401 is returned for unauthorized access
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "POST",
		Route:  "/api/v1/pnet-server/ping",
		Body:   FiberRequestPayload(input),
	})

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "Unauthorized", description)
}

func TestV1PNETServerPingInvalidAttempt(t *testing.T) {
	description := "ping PNET server with invalid attempt"
	f := NewFiberTestHTTP()

	authHeader, _ := f.AuthorizationServiceBasic()

	input := external.PNETServerPingInputDTO{
		Attempts: []external.AttemptDTO{
			{
				AttemptID: "invalid-attempt-id",
				UserEmail: "user@example.com",
				UserID:    1,
			},
		},
	}

	expectedCode := 200 // Assuming 200 is returned even if some attempts are invalid
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/pnet-server/ping",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := response.Response[external.PNETServerPingOutputDTO]{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 0, bodyModel.Result.Count, description)
}

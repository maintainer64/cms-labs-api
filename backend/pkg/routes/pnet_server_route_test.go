package routes

import (
	"testing"
	"time"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"gitlab.com/a10869/api-modules/backend/app/usecases"

	"gitlab.com/a10869/api-modules/backend/app/models"

	"github.com/stretchr/testify/assert"
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
	authHeader := f.AuthorizationUser(0, 0, "")
	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/pnet-server/get",
		FiberRequestPayload(map[string]any{
			"id": entity.ID,
		}),
		authHeader,
	)
	bodyModel := usecases.PNETServerGetResponse{}
	json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, entity.ID, bodyModel.Result.Model.ID, description)
}

func TestV1PNETServerRouteGetNotFound(t *testing.T) {
	description := "not found pnet"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0, "")
	expectedBody := map[string]interface{}{
		"error": true,
		"msg":   "PNETServer not found",
	}
	expectedCode := 404
	statusCode, body := f.Request(
		"POST",
		"/api/v1/pnet-server/get",
		FiberRequestPayload(map[string]any{
			"id": 9999999,
		}),
		authHeader,
	)

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
	authHeader := f.AuthorizationUser(0, 0, "")
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
		statusCode, body := f.Request(
			"POST",
			"/api/v1/pnet-server/list",
			FiberRequestPayload(test.body),
			authHeader,
		)
		bodyModel := usecases.PNETServerListResponse{}
		json.Unmarshal([]byte(body), &bodyModel)
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

	authHeader := f.AuthorizationUser(0, 0, "")
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
		statusCode, body := f.Request(
			"POST",
			"/api/v1/pnet-server/delete",
			FiberRequestPayload(map[string]any{
				"id": test.id,
			}),
			authHeader,
		)
		response := usecases.PNETServerDeleteResponse{}
		json.Unmarshal([]byte(body), &response)
		assert.Equal(t, 200, statusCode, test.description)
		assert.Equal(t, test.id, response.Result.ID, test.description)

		// Проверяем, что сервер успешно удален
		var deletedEntity models.PNETServer
		f.DB.First(&deletedEntity, test.id)
		var idNotFound uint = 0
		assert.Equal(t, idNotFound, deletedEntity.ID, test.description)
	}
}

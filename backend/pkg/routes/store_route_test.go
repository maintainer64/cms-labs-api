package routes

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/models/types"
)

func TestStoreGet(t *testing.T) {
	description := "get user store"
	f := NewFiberTestHTTP()

	testUser := models.User{
		UserBase: models.UserBase{
			Email: uuid.New().String() + "@user.com",
			Name:  "Test User",
		},
		UserSecret: models.UserSecret{
			Store: types.UserStore{
				"theme": "dark",
				"prefs": map[string]interface{}{"notifications": true},
				"word":  "Русские символы в БД",
			},
		},
	}
	f.DB.Create(&testUser)

	authHeader := f.AuthorizationUser(testUser.ID, 0)

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "GET",
		Route:         "/api/v1/global-store",
		Authorization: authHeader,
	})

	var response types.UserStore
	err := json.Unmarshal([]byte(body), &response)
	assert.NoError(t, err, description)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, "dark", response["theme"], description)
	assert.True(t, response["prefs"].(map[string]interface{})["notifications"].(bool), description)
}

func TestStoreGetEmpty(t *testing.T) {
	description := "get empty user store"
	f := NewFiberTestHTTP()

	testUser := models.User{
		UserBase: models.UserBase{
			Email: uuid.New().String() + "@user.com",
			Name:  "Test User",
		},
	}
	f.DB.Create(&testUser)

	authHeader := f.AuthorizationUser(testUser.ID, 0)

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "GET",
		Route:         "/api/v1/global-store",
		Authorization: authHeader,
	})

	var response types.UserStore
	err := json.Unmarshal([]byte(body), &response)
	assert.NoError(t, err, description)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Empty(t, response, description)
}

func TestStoreSet(t *testing.T) {
	description := "set user store"
	f := NewFiberTestHTTP()

	testUser := models.User{
		UserBase: models.UserBase{
			Email: uuid.New().String() + "@user.com",
			Name:  "Test User",
		},
	}
	f.DB.Create(&testUser)

	authHeader := f.AuthorizationUser(testUser.ID, 0)

	input := map[string]interface{}{
		"theme": "light",
		"prefs": map[string]interface{}{
			"notifications": false,
			"language":      "en",
		},
	}
	jsonData, _ := json.Marshal(input)

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/global-store",
		Body:          strings.NewReader(string(jsonData)),
		Authorization: authHeader,
	})

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Empty(t, body, description)

	// Verify database record
	var user models.User
	err := f.DB.First(&user, testUser.ID).Error
	assert.NoError(t, err, description)
	assert.Equal(t, "light", user.Store["theme"], description)
	assert.False(t, user.Store["prefs"].(map[string]interface{})["notifications"].(bool), description)
	assert.Equal(t, "en", user.Store["prefs"].(map[string]interface{})["language"].(string), description)
}

func TestStoreSetInvalidJSON(t *testing.T) {
	description := "set user store with invalid JSON"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)

	expectedCode := 400
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/global-store",
		Body:          strings.NewReader("{invalid json}"),
		Authorization: authHeader,
	})

	var response map[string]interface{}
	err := json.Unmarshal([]byte(body), &response)
	assert.NoError(t, err, description)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, response["msg"], "invalid", description)
}

func TestStoreUnauthorized(t *testing.T) {
	tests := []struct {
		method       string
		route        string
		description  string
		expectedCode int
	}{
		{
			method:       "GET",
			route:        "/api/v1/global-store",
			description:  "get store unauthorized",
			expectedCode: 401,
		},
		{
			method:       "POST",
			route:        "/api/v1/global-store",
			description:  "set store unauthorized",
			expectedCode: 401,
		},
	}

	f := NewFiberTestHTTP()

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			statusCode, body := f.Request(&FiberTestHttpRequest{
				Method: test.method,
				Route:  test.route,
			})

			var response map[string]interface{}
			err := json.Unmarshal([]byte(body), &response)
			assert.NoError(t, err, test.description)

			assert.Equal(t, test.expectedCode, statusCode, test.description)
			assert.Contains(t, response["msg"], "Invalid token", test.description)
		})
	}
}

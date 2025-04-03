package routes

import (
	"fmt"
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func TestV1ServiceCardCreate(t *testing.T) {
	description := "create service card"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)
	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.ServiceCard{})
	input := usecases.ServiceCardEditInputDTO{
		ImageUrl:    "https://example.com/image.jpg",
		Url:         "https://example.com",
		Name:        "Test Service",
		Description: "Test Description",
		Order:       1,
		IsActive:    true,
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/service-card/upsert",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.ServiceCardEditResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.NotZero(t, bodyModel.Result.ID, description)

	// Verify database record
	var entity models.ServiceCard
	err := f.DB.First(&entity, bodyModel.Result.ID).Error
	assert.NoError(t, err, description)
	assert.Equal(t, input.Name, entity.Name, description)
	assert.Equal(t, input.Url, entity.Url, description)
	assert.Equal(t, input.Description, entity.Description, description)
}

func TestV1ServiceCardList(t *testing.T) {
	description := "list service cards"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)
	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.ServiceCard{})

	entities := make([]models.ServiceCard, 0)
	for i := 1; i <= 3; i++ {
		entity := models.ServiceCard{}
		entity.ImageUrl = fmt.Sprintf("image%d.jpg", i)
		entity.Url = fmt.Sprintf("service%d.com", i)
		entity.Name = fmt.Sprintf("Service %d", i)
		entity.Description = fmt.Sprintf("%d service", i)
		entity.Order = uint(i + 1)
		entity.IsActive = true
		entities = append(entities, entity)
		f.DB.Create(&entity)
	}

	input := usecases.ServiceCardListInputDTO{
		Limit:  10,
		Offset: 0,
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/service-card/list",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.ServiceCardListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, int64(3), bodyModel.Result.TotalCount, description)
	assert.Len(t, bodyModel.Result.Model, 3, description)

	// Verify response content
	for i, item := range bodyModel.Result.Model {
		assert.Equal(t, entities[len(entities)-i-1].Name, item.Name, description)
		assert.Equal(t, entities[len(entities)-i-1].Url, item.Url, description)
	}
}

func TestV1ServiceCardListPagination(t *testing.T) {
	description := "list service cards with pagination"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)
	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.ServiceCard{})

	// Create 3 test records
	for i := 1; i <= 3; i++ {
		entity := models.ServiceCard{}
		entity.ImageUrl = fmt.Sprintf("image%d.jpg", i)
		entity.Url = fmt.Sprintf("service%d.com", i)
		entity.Name = fmt.Sprintf("Service %d", i)
		entity.Description = fmt.Sprintf("%d service", i)
		entity.Order = uint(i + 1)
		entity.IsActive = true
		f.DB.Create(&entity)
	}

	input := usecases.ServiceCardListInputDTO{
		Limit:  2,
		Offset: 1,
	}

	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/service-card/list",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.ServiceCardListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, 200, statusCode, description)
	assert.Equal(t, int64(3), bodyModel.Result.TotalCount, description)
	assert.Len(t, bodyModel.Result.Model, 2, description)
	assert.Equal(t, "Service 2", bodyModel.Result.Model[0].Name, description)
}

func TestV1ServiceCardListEmpty(t *testing.T) {
	description := "list empty service cards"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.ServiceCard{})

	input := usecases.ServiceCardListInputDTO{
		Search: "NonExistingService",
		Limit:  10,
	}

	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/service-card/list",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.ServiceCardListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, 200, statusCode, description)
	assert.Equal(t, int64(0), bodyModel.Result.TotalCount, description)
	assert.Len(t, bodyModel.Result.Model, 0, description)
}

func TestV1ServiceCardGet(t *testing.T) {
	description := "get service card"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.ServiceCard{})

	// Create a test service card

	entity := models.ServiceCard{}
	entity.ImageUrl = "https://example.com/image.jpg"
	entity.Url = "https://example.com"
	entity.Name = "Test Service"
	entity.Description = "Test Description"
	entity.Order = 1
	entity.IsActive = true
	f.DB.Create(&entity)

	input := usecases.ServiceCardGetInputDTO{
		ID: entity.ID, // Use the ID of the created entity
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/service-card/get",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.ServiceCardGetResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, entity.ID, bodyModel.Result.Model.ID, description)
	assert.Equal(t, entity.Name, bodyModel.Result.Model.Name, description)
	assert.Equal(t, entity.Url, bodyModel.Result.Model.Url, description)
	assert.Equal(t, entity.Description, bodyModel.Result.Model.Description, description)
	assert.Equal(t, entity.Order, bodyModel.Result.Model.Order, description)
	assert.Equal(t, entity.IsActive, bodyModel.Result.Model.IsActive, description)
}

func TestV1ServiceCardGetNotFound(t *testing.T) {
	description := "get non-existent service card"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.ServiceCard{})

	// Use a non-existent ID
	input := usecases.ServiceCardGetInputDTO{
		ID: 9999,
	}

	expectedCode := 404 // Assuming 404 is returned for not found
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/service-card/get",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "not found", description) // Adjust based on your error response format
}

func TestV1ServiceCardDelete(t *testing.T) {
	description := "delete service card"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.ServiceCard{})

	// Create a test service card

	entity := models.ServiceCard{}
	entity.ImageUrl = "https://example.com/image.jpg"
	entity.Url = "https://example.com"
	entity.Name = "Test Service"
	entity.Description = "Test Description"
	entity.Order = 1
	entity.IsActive = true
	f.DB.Create(&entity)

	input := usecases.ServiceCardDeleteInputDTO{
		ID: entity.ID, // Use the ID of the created entity
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/service-card/delete",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.ServiceCardDeleteResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, entity.ID, bodyModel.Result.ID, description)

	var entityDB models.ServiceCard
	result := f.DB.First(&entityDB, entity.ID)
	assert.Equal(t, result.Error.Error(), "record not found", description)
}

func TestV1NotFound(t *testing.T) {
	description := "not found route"
	f := NewFiberTestHTTP()
	expectedCode := 404
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "POST",
		Route:  "/api/v1/not-found-route",
		Body:   FiberRequestPayload(""),
	})
	bodyModel := map[string]interface{}{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.True(t, bodyModel["error"].(bool), description)
	assert.Equal(t, "sorry, endpoint is not found", bodyModel["msg"].(string), description)
}

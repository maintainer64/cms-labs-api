package routes

import (
	"testing"
	"time"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func TestV1UserCreate(t *testing.T) {
	description := "create user"
	f := NewFiberTestHTTP()

	entityDB := models.User{}
	entityDB.Email = uuid.New().String() + "@example.com"
	f.DB.Create(&entityDB)

	roleStudent := models.Role{}
	roleStudent.Code = "student" + uuid.New().String()
	roleStudent.Name = "Student"
	f.DB.Create(&roleStudent)

	roleAdmin := models.Role{}
	roleAdmin.Code = "admin" + uuid.New().String()
	roleAdmin.Name = "Admin"
	f.DB.Create(&roleAdmin)

	roleUser := models.RoleRelation{}
	roleUser.RoleID = roleStudent.ID
	roleUser.UserID = &entityDB.ID

	f.DB.Create(&roleUser)

	input := usecases.UserEditInputDTO{
		ID:        entityDB.ID,
		Name:      "John Doe",
		Email:     entityDB.Email,
		LTIUserID: "lti123",
		Roles:     []uint{roleAdmin.ID},
		GroupName: "Group A",
		IsActive:  true,
	}

	authHeader := f.AuthorizationUser(0, 0)

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/user/upsert",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.UserEditResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.NotZero(t, bodyModel.Result.ID, description)

	// Verify database record
	var entity models.User
	err := f.DB.First(&entity, bodyModel.Result.ID).Error
	assert.NoError(t, err, description)
	assert.Equal(t, input.Name, entity.Name, description)
	assert.Equal(t, input.Email, entity.Email, description)
	assert.Equal(t, input.LTIUserID, entity.LTIUserID, description)
	assert.Equal(t, input.GroupName, entity.GroupName, description)
	assert.True(t, entity.IsActive(), description)

	var entityRelationRole models.RoleRelation
	err = f.DB.Where("user_id = ?", entity.ID).Find(&entityRelationRole).Error
	assert.NoError(t, err, description)
	assert.Equal(t, entityRelationRole.RoleID, roleAdmin.ID, description)
}

func TestV1UserCreateUnauthorized(t *testing.T) {
	description := "create user unauthorized"
	f := NewFiberTestHTTP()

	// No auth header provided
	input := usecases.UserEditInputDTO{
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		LTIUserID: "lti123",
		GroupName: "Group A",
		IsActive:  true,
	}

	expectedCode := 401 // Assuming 401 is returned for unauthorized access
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "POST",
		Route:  "/api/v1/user/upsert",
		Body:   FiberRequestPayload(input),
	})

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "token is malformed: token contains an invalid number of segments", description)
}

func TestV1UserCreateDeactivated(t *testing.T) {
	description := "create deactivated user"
	f := NewFiberTestHTTP()

	entityDB := models.User{}
	entityDB.Email = uuid.New().String() + "@example.com"
	entityDB.DeletedAt = nil
	f.DB.Create(&entityDB)

	authHeader := f.AuthorizationUser(0, 0)

	input := usecases.UserEditInputDTO{
		ID:        entityDB.ID,
		Name:      "John Doe",
		Email:     entityDB.Email,
		LTIUserID: "lti123",
		GroupName: "Group A",
		IsActive:  false,
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/user/upsert",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.UserEditResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.NotZero(t, bodyModel.Result.ID, description)

	// Verify database record
	var entity models.User
	err := f.DB.First(&entity, bodyModel.Result.ID).Error
	assert.NoError(t, err, description)
	assert.False(t, entity.IsActive(), description)
}

func TestV1UserGetSuccess(t *testing.T) {
	description := "get user successfully"
	f := NewFiberTestHTTP()

	// Create a test user
	entityDB := models.User{}
	entityDB.Email = uuid.New().String() + "@example.com"
	entityDB.Name = "John Doe"
	entityDB.GroupName = "Group A"
	entityDB.LTIUserID = "lti123"
	entityDB.DeletedAt = nil
	f.DB.Create(&entityDB)

	authHeader := f.AuthorizationUser(0, 0)

	input := usecases.UserGetInputDTO{
		ID: entityDB.ID,
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/user/get",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.UserGetResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, entityDB.ID, bodyModel.Result.Model.ID, description)
	assert.Equal(t, entityDB.Name, bodyModel.Result.Model.Name, description)
	assert.Equal(t, entityDB.Email, bodyModel.Result.Model.Email, description)
	assert.Equal(t, entityDB.GroupName, bodyModel.Result.Model.GroupName, description)
	assert.Equal(t, entityDB.LTIUserID, bodyModel.Result.Model.LTIUserID, description)
	assert.True(t, bodyModel.Result.Model.IsActive(), description)
}

func TestV1UserGetNotFound(t *testing.T) {
	description := "get non-existent user"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0)

	// Use a non-existent ID
	input := usecases.UserGetInputDTO{
		ID: 9999,
	}

	expectedCode := 404 // Assuming 404 is returned for not found
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/user/get",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "not found", description) // Adjust based on your error response format
}

func TestV1UserGetInactiveUser(t *testing.T) {
	description := "get inactive user"
	f := NewFiberTestHTTP()

	// Create an inactive user
	deletedAt := time.Now()
	entityDB := models.User{}
	entityDB.Email = uuid.New().String() + "@example.com"
	entityDB.Name = "John Doe"
	entityDB.GroupName = "Group A"
	entityDB.LTIUserID = "lti123"
	entityDB.DeletedAt = &deletedAt
	f.DB.Create(&entityDB)

	authHeader := f.AuthorizationUser(0, 0)

	input := usecases.UserGetInputDTO{
		ID: entityDB.ID,
	}

	expectedCode := 200 // Assuming 200 is returned even for inactive users
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/user/get",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.UserGetResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, entityDB.ID, bodyModel.Result.Model.ID, description)
	assert.False(t, bodyModel.Result.Model.IsActive(), description)
}

func TestV1UserListSuccess(t *testing.T) {
	description := "list users successfully"
	f := NewFiberTestHTTP()

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.User{})

	// Create test users
	entityDB1 := models.User{}
	entityDB1.Email = uuid.New().String() + "@example.com"
	entityDB1.Name = "John Doe"
	entityDB1.GroupName = "Group A"
	entityDB1.LTIUserID = "lti123"
	f.DB.Create(&entityDB1)

	entityDB2 := models.User{}
	entityDB2.Email = uuid.New().String() + "@example.com"
	entityDB2.Name = "Jane Doe"
	entityDB2.GroupName = "Group B"
	entityDB2.LTIUserID = "lti456"
	f.DB.Create(&entityDB2)

	authHeader := f.AuthorizationUser(0, 0)

	input := usecases.UserListInputDTO{
		Search: "",
		Limit:  10,
		Offset: 0,
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/user/list",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.UserListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 3, len(bodyModel.Result.Model), description)
}

func TestV1UserListFilterBySearch(t *testing.T) {
	description := "list users with search filter"
	f := NewFiberTestHTTP()

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.User{})

	// Create test users
	entityDB1 := models.User{}
	entityDB1.Email = uuid.New().String() + "@example.com"
	entityDB1.Name = "John Doe"
	entityDB1.GroupName = "Group A"
	entityDB1.LTIUserID = "lti123"
	f.DB.Create(&entityDB1)

	entityDB2 := models.User{}
	entityDB2.Email = uuid.New().String() + "@example.com"
	entityDB2.Name = "Jane Doe"
	entityDB2.GroupName = "Group B"
	entityDB2.LTIUserID = "lti456"
	f.DB.Create(&entityDB2)

	authHeader := f.AuthorizationUser(0, 0)

	input := usecases.UserListInputDTO{
		Search: "John",
		Limit:  10,
		Offset: 0,
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/user/list",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.UserListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 1, len(bodyModel.Result.Model), description)
	assert.Equal(t, entityDB1.Name, bodyModel.Result.Model[0].Model.Name, description)
}

func TestV1UserListFilterByIDs(t *testing.T) {
	description := "list users with ID filter"
	f := NewFiberTestHTTP()

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.User{})

	// Create test users
	entityDB1 := models.User{}
	entityDB1.Email = uuid.New().String() + "@example.com"
	entityDB1.Name = "John Doe"
	entityDB1.GroupName = "Group A"
	entityDB1.LTIUserID = "lti123"
	f.DB.Create(&entityDB1)

	entityDB2 := models.User{}
	entityDB2.Email = uuid.New().String() + "@example.com"
	entityDB2.Name = "Jane Doe"
	entityDB2.GroupName = "Group B"
	entityDB2.LTIUserID = "lti456"
	f.DB.Create(&entityDB2)

	authHeader := f.AuthorizationUser(0, 0)

	input := usecases.UserListInputDTO{
		Search:  "",
		UserIds: []uint{entityDB1.ID},
		Limit:   10,
		Offset:  0,
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/user/list",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.UserListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 1, len(bodyModel.Result.Model), description)
	assert.Equal(t, entityDB1.ID, bodyModel.Result.Model[0].Model.ID, description)
}

func TestV1UserListPagination(t *testing.T) {
	description := "list users with pagination"
	f := NewFiberTestHTTP()

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(&models.User{})

	// Create test users
	entityDB1 := models.User{}
	entityDB1.Email = uuid.New().String() + "@example.com"
	entityDB1.Name = "John Doe"
	entityDB1.GroupName = "Group A"
	entityDB1.LTIUserID = "lti123"
	f.DB.Create(&entityDB1)

	entityDB2 := models.User{}
	entityDB2.Email = uuid.New().String() + "@example.com"
	entityDB2.Name = "Jane Doe"
	entityDB2.GroupName = "Group B"
	entityDB2.LTIUserID = "lti456"
	f.DB.Create(&entityDB2)

	authHeader := f.AuthorizationUser(0, 0)

	input := usecases.UserListInputDTO{
		Search: "",
		Limit:  1,
		Offset: 2,
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/user/list",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})
	bodyModel := usecases.UserListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 1, len(bodyModel.Result.Model), description)
	assert.Equal(t, entityDB1.ID, bodyModel.Result.Model[0].Model.ID, description)
}

func TestV1UserListUnauthorized(t *testing.T) {
	description := "list users unauthorized"
	f := NewFiberTestHTTP()

	input := usecases.UserListInputDTO{
		Search:  "",
		UserIds: []uint{},
		Limit:   10,
		Offset:  0,
	}

	expectedCode := 401 // Assuming 401 is returned for unauthorized access
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "POST",
		Route:  "/api/v1/user/list",
		Body:   FiberRequestPayload(input),
	})

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "token is malformed: token contains an invalid number of segments", description)
}

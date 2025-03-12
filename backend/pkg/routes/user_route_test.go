package routes

import (
	"testing"
	"time"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func TestV1UserCreate(t *testing.T) {
	description := "create user"
	f := NewFiberTestHTTP()

	entityDB := models.User{}
	entityDB.Email = uuid.New().String() + "@example.com"
	entityDB.UserRole = models.UsersRoleAdmin
	f.DB.Create(&entityDB)

	input := usecases.UserEditInputDTO{
		ID:        entityDB.ID,
		Name:      "John Doe",
		Email:     entityDB.Email,
		LTIUserID: "lti123",
		UserRole:  models.UsersRoleStudent,
		GroupName: "Group A",
		IsActive:  true,
	}

	authHeader := f.AuthorizationUser(entityDB.ID, 0, "")

	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/user/upsert",
		FiberRequestPayload(input),
		authHeader,
	)
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
	assert.Equal(t, input.UserRole, entity.UserRole, description)
	assert.Equal(t, input.GroupName, entity.GroupName, description)
	assert.True(t, entity.IsActive(), description)
}

func TestV1UserCreateUnauthorized(t *testing.T) {
	description := "create user unauthorized"
	f := NewFiberTestHTTP()

	// No auth header provided
	input := usecases.UserEditInputDTO{
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		LTIUserID: "lti123",
		UserRole:  "student",
		GroupName: "Group A",
		IsActive:  true,
	}

	expectedCode := 401 // Assuming 401 is returned for unauthorized access
	statusCode, body := f.Request(
		"POST",
		"/api/v1/user/upsert",
		FiberRequestPayload(input),
		"",
	)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "token is malformed: token contains an invalid number of segments", description)
}

func TestV1UserCreateDeactivated(t *testing.T) {
	description := "create deactivated user"
	f := NewFiberTestHTTP()

	entityDB := models.User{}
	entityDB.Email = uuid.New().String() + "@example.com"
	entityDB.UserRole = models.UsersRoleAdmin
	entityDB.DeletedAt = nil
	f.DB.Create(&entityDB)

	authHeader := f.AuthorizationUser(entityDB.ID, 0, "")

	input := usecases.UserEditInputDTO{
		ID:        entityDB.ID,
		Name:      "John Doe",
		Email:     entityDB.Email,
		LTIUserID: "lti123",
		UserRole:  models.UsersRoleStudent,
		GroupName: "Group A",
		IsActive:  false,
	}

	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/user/upsert",
		FiberRequestPayload(input),
		authHeader,
	)
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
	entityDB.UserRole = models.UsersRoleAdmin
	entityDB.DeletedAt = nil
	f.DB.Create(&entityDB)

	authHeader := f.AuthorizationUser(entityDB.ID, 0, "")

	input := usecases.UserGetInputDTO{
		ID: entityDB.ID,
	}

	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/user/get",
		FiberRequestPayload(input),
		authHeader,
	)
	bodyModel := usecases.UserGetResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, entityDB.ID, bodyModel.Result.Model.ID, description)
	assert.Equal(t, entityDB.Name, bodyModel.Result.Model.Name, description)
	assert.Equal(t, entityDB.Email, bodyModel.Result.Model.Email, description)
	assert.Equal(t, entityDB.UserRole, bodyModel.Result.Model.UserRole, description)
	assert.Equal(t, entityDB.GroupName, bodyModel.Result.Model.GroupName, description)
	assert.Equal(t, entityDB.LTIUserID, bodyModel.Result.Model.LTIUserID, description)
	assert.True(t, bodyModel.Result.Model.IsActive(), description)
}

func TestV1UserGetNotFound(t *testing.T) {
	description := "get non-existent user"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0, "")

	// Use a non-existent ID
	input := usecases.UserGetInputDTO{
		ID: 9999,
	}

	expectedCode := 404 // Assuming 404 is returned for not found
	statusCode, body := f.Request(
		"POST",
		"/api/v1/user/get",
		FiberRequestPayload(input),
		authHeader,
	)

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
	entityDB.UserRole = models.UsersRoleStudent
	entityDB.DeletedAt = &deletedAt
	f.DB.Create(&entityDB)

	authHeader := f.AuthorizationUser(0, 0, "")

	input := usecases.UserGetInputDTO{
		ID: entityDB.ID,
	}

	expectedCode := 200 // Assuming 200 is returned even for inactive users
	statusCode, body := f.Request(
		"POST",
		"/api/v1/user/get",
		FiberRequestPayload(input),
		authHeader,
	)
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
	entityDB1.UserRole = models.UsersRoleAdmin
	entityDB1.GroupName = "Group A"
	entityDB1.LTIUserID = "lti123"
	f.DB.Create(&entityDB1)

	entityDB2 := models.User{}
	entityDB2.Email = uuid.New().String() + "@example.com"
	entityDB2.Name = "Jane Doe"
	entityDB2.UserRole = models.UsersRoleInstructor
	entityDB2.GroupName = "Group B"
	entityDB2.LTIUserID = "lti456"
	f.DB.Create(&entityDB2)

	authHeader := f.AuthorizationUser(entityDB1.ID, 0, "")

	input := usecases.UserListInputDTO{
		Search: "",
		Limit:  10,
		Offset: 0,
	}

	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/user/list",
		FiberRequestPayload(input),
		authHeader,
	)
	bodyModel := usecases.UserListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 2, len(bodyModel.Result.Model), description)
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
	entityDB1.UserRole = models.UsersRoleAdmin
	entityDB1.GroupName = "Group A"
	entityDB1.LTIUserID = "lti123"
	f.DB.Create(&entityDB1)

	entityDB2 := models.User{}
	entityDB2.Email = uuid.New().String() + "@example.com"
	entityDB2.Name = "Jane Doe"
	entityDB2.UserRole = models.UsersRoleInstructor
	entityDB2.GroupName = "Group B"
	entityDB2.LTIUserID = "lti456"
	f.DB.Create(&entityDB2)

	authHeader := f.AuthorizationUser(entityDB1.ID, 0, "")

	input := usecases.UserListInputDTO{
		Search: "John",
		Limit:  10,
		Offset: 0,
	}

	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/user/list",
		FiberRequestPayload(input),
		authHeader,
	)
	bodyModel := usecases.UserListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 1, len(bodyModel.Result.Model), description)
	assert.Equal(t, entityDB1.Name, bodyModel.Result.Model[0].Name, description)
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
	entityDB1.UserRole = models.UsersRoleAdmin
	entityDB1.GroupName = "Group A"
	entityDB1.LTIUserID = "lti123"
	f.DB.Create(&entityDB1)

	entityDB2 := models.User{}
	entityDB2.Email = uuid.New().String() + "@example.com"
	entityDB2.Name = "Jane Doe"
	entityDB2.UserRole = models.UsersRoleInstructor
	entityDB2.GroupName = "Group B"
	entityDB2.LTIUserID = "lti456"
	f.DB.Create(&entityDB2)

	authHeader := f.AuthorizationUser(entityDB1.ID, 0, "")

	input := usecases.UserListInputDTO{
		Search:  "",
		UserIds: []uint{entityDB1.ID},
		Limit:   10,
		Offset:  0,
	}

	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/user/list",
		FiberRequestPayload(input),
		authHeader,
	)
	bodyModel := usecases.UserListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 1, len(bodyModel.Result.Model), description)
	assert.Equal(t, entityDB1.ID, bodyModel.Result.Model[0].ID, description)
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
	entityDB1.UserRole = models.UsersRoleAdmin
	entityDB1.GroupName = "Group A"
	entityDB1.LTIUserID = "lti123"
	f.DB.Create(&entityDB1)

	entityDB2 := models.User{}
	entityDB2.Email = uuid.New().String() + "@example.com"
	entityDB2.Name = "Jane Doe"
	entityDB2.UserRole = models.UsersRoleInstructor
	entityDB2.GroupName = "Group B"
	entityDB2.LTIUserID = "lti456"
	f.DB.Create(&entityDB2)

	authHeader := f.AuthorizationUser(entityDB1.ID, 0, "")

	input := usecases.UserListInputDTO{
		Search: "",
		Limit:  1,
		Offset: 1,
	}

	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/user/list",
		FiberRequestPayload(input),
		authHeader,
	)
	bodyModel := usecases.UserListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 1, len(bodyModel.Result.Model), description)
	assert.Equal(t, entityDB1.ID, bodyModel.Result.Model[0].ID, description)
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
	statusCode, body := f.Request(
		"POST",
		"/api/v1/user/list",
		FiberRequestPayload(input),
		"",
	)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "token is malformed: token contains an invalid number of segments", description)
}

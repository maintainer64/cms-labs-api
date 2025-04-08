package routes

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func TestRoleUpsert(t *testing.T) {
	description := "successful role upsert (create)"
	f := NewFiberTestHTTP()

	authHeader := f.AuthorizationUser(0, 0)

	input := usecases.RoleEditInputDTO{
		Name: "Test Role " + uuid.New().String(),
		Code: "test-role-" + uuid.New().String(),
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/role/upsert",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})

	response := usecases.RoleEditResponse{}
	_ = json.Unmarshal([]byte(body), &response)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.NotZero(t, response.Result.ID, description)

	// Verify role was created in DB
	var role models.Role
	f.DB.Where("id = ?", response.Result.ID).First(&role)
	assert.Equal(t, input.Name, role.Name, description)
	assert.Equal(t, input.Code, role.Code, description)
}

func TestRoleUpsertUpdate(t *testing.T) {
	description := "successful role upsert (update)"
	f := NewFiberTestHTTP()

	authHeader := f.AuthorizationUser(0, 0)

	// Create existing role to update
	existingRole := models.Role{}
	existingRole.Name = "Old Role Name"
	existingRole.Code = "old-role-code"
	f.DB.Create(&existingRole)

	input := usecases.RoleEditInputDTO{
		ID:   existingRole.ID,
		Name: "Updated Role Name",
		Code: "updated-role-code",
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/role/upsert",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})

	response := usecases.RoleEditResponse{}
	_ = json.Unmarshal([]byte(body), &response)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, existingRole.ID, response.Result.ID, description)

	// Verify role was updated in DB
	var role models.Role
	f.DB.Where("id = ?", existingRole.ID).First(&role)
	assert.Equal(t, input.Name, role.Name, description)
	assert.Equal(t, input.Code, role.Code, description)
	assert.NotEqual(t, existingRole.UpdatedAt, role.UpdatedAt, description)
}

func TestRoleUpsertUnauthorized(t *testing.T) {
	description := "role upsert without admin privileges"
	f := NewFiberTestHTTP()

	// Create regular user (non-admin)
	regularUser := models.User{}
	regularUser.Email = "user" + uuid.New().String() + "@example.com"
	f.DB.Create(&regularUser)
	authHeader := f.AuthorizationUser(regularUser.ID, 0)

	input := usecases.RoleEditInputDTO{
		Name: "Test Role",
		Code: "test-role",
	}

	expectedCode := fiber.StatusUnauthorized
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/role/upsert",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "User with current role is not allow action", description)
}

func TestRoleList(t *testing.T) {
	description := "successful role list"
	f := NewFiberTestHTTP()

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(models.Role{})

	// Create some test roles
	for i := 0; i < 5; i++ {
		role := models.Role{}
		role.Name = fmt.Sprintf("Role %d", i)
		role.Name = fmt.Sprintf("role-%d", i)
		f.DB.Create(&role)
	}

	input := usecases.RoleListInputDTO{}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method: "POST",
		Route:  "/api/v1/role/list",
		Body:   FiberRequestPayload(input),
	})

	response := usecases.RoleListResponse{}
	_ = json.Unmarshal([]byte(body), &response)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Len(t, response.Result.Model, 5, description)
	assert.Equal(t, 5, response.Result.TotalCount, description)
}

func TestRoleDelete(t *testing.T) {
	description := "successful role delete"
	f := NewFiberTestHTTP()

	authHeader := f.AuthorizationUser(0, 0)

	// Create role to delete
	roleToDelete := models.Role{}
	roleToDelete.Name = "To Delete"
	roleToDelete.Code = "to-delete"
	f.DB.Create(&roleToDelete)

	input := usecases.RoleDeleteInputDTO{
		ID: roleToDelete.ID,
	}

	expectedCode := 200
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/role/delete",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})

	response := usecases.RoleDeleteResponse{}
	_ = json.Unmarshal([]byte(body), &response)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, roleToDelete.ID, response.Result.ID, description)

	// Verify role was deleted from DB
	var role models.Role
	result := f.DB.Where("id = ?", roleToDelete.ID).First(&role)
	assert.Error(t, result.Error, description)
	assert.Contains(t, result.Error.Error(), "record not found", description)
}

func TestRoleDeleteUnauthorized(t *testing.T) {
	description := "role delete without admin privileges"
	f := NewFiberTestHTTP()

	// Create regular user (non-admin)
	regularUser := models.User{}
	regularUser.Email = "user" + uuid.New().String() + "@example.com"
	f.DB.Create(&regularUser)
	authHeader := f.AuthorizationUser(regularUser.ID, 0)

	// Create role that would be deleted if authorized
	role := models.Role{}
	role.Name = "Test Role"
	role.Code = "test-role"
	f.DB.Create(&role)

	input := usecases.RoleDeleteInputDTO{
		ID: role.ID,
	}

	expectedCode := fiber.StatusUnauthorized
	statusCode, body := f.Request(&FiberTestHttpRequest{
		Method:        "POST",
		Route:         "/api/v1/role/delete",
		Body:          FiberRequestPayload(input),
		Authorization: authHeader,
	})

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "User with current role is not allow action", description)

	// Verify role was NOT deleted from DB
	var dbRole models.Role
	result := f.DB.Where("id = ?", role.ID).First(&dbRole)
	assert.NoError(t, result.Error, description)
}

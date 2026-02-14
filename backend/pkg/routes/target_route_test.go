package routes

import (
	"testing"

	json "github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/models/types"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func TestV1TargetUpsertCreate(t *testing.T) {
	desc := "create target"
	f := NewTestHTTP()
	authAdmin := f.AuthorizationUser(0, nil)

	input := usecases.TargetUpsertInputDTO{
		ID:          nil, // создание
		Name:        "Test Target " + uuid.New().String(),
		Description: ptrString("Test Description"),
		Type:        "service",
		Links:       types.JsonStore{"doc": "https://example.com"},
		Tags:        types.JsonStore{"env": "prod"},
	}

	code, body := f.Rpc(&TestRpcRequest{
		Method:        "target.upsert",
		Params:        input,
		Authorization: authAdmin,
	})
	var resp usecases.TargetUpsertResponse
	err := json.Unmarshal([]byte(body), &resp)
	assert.NoError(t, err, desc)
	assert.Equal(t, 200, code, desc)
	assert.NotEmpty(t, resp.Result.ID, desc)

	// Проверка в БД
	var target models.Target
	err = f.DB.First(&target, "id = ?", resp.Result.ID).Error
	assert.NoError(t, err, desc)
	assert.Equal(t, input.Name, target.Name, desc)
	assert.Equal(t, *input.Description, *target.Description, desc)

	// Проверка, что создатель стал редактором
	var tu models.TargetUser
	err = f.DB.First(&tu, "target_id = ?", resp.Result.ID).Error
	assert.NoError(t, err, desc)
	roles, ok := tu.Roles["roles"].([]interface{})
	assert.True(t, ok, desc)
	assert.Contains(t, roles, models.UserRoleEditor, desc)
}

func TestV1TargetUpsertUpdate(t *testing.T) {
	desc := "update target"
	f := NewTestHTTP()
	authAdmin := f.AuthorizationUser(0, nil)

	// Сначала создадим цель
	createInput := usecases.TargetUpsertInputDTO{
		ID:          nil,
		Name:        "Original" + uuid.New().String(),
		Description: ptrString("Original desc"),
		Type:        "service",
		Links:       types.JsonStore{},
		Tags:        types.JsonStore{},
	}
	code, body := f.Rpc(&TestRpcRequest{
		Method:        "target.upsert",
		Params:        createInput,
		Authorization: authAdmin,
	})
	var createResp usecases.TargetUpsertResponse
	json.Unmarshal([]byte(body), &createResp)
	targetID := createResp.Result.ID

	// Обновляем
	updateInput := usecases.TargetUpsertInputDTO{
		ID:          &targetID,
		Name:        "Updated" + uuid.New().String(),
		Description: ptrString("Updated desc"),
		Type:        "server",
		Links:       types.JsonStore{"new": "link"},
		Tags:        types.JsonStore{"env": "stage"},
	}

	code, body = f.Rpc(&TestRpcRequest{
		Method:        "target.upsert",
		Params:        updateInput,
		Authorization: authAdmin,
	})
	var updateResp usecases.TargetUpsertResponse
	err := json.Unmarshal([]byte(body), &updateResp)
	assert.NoError(t, err, desc)
	assert.Equal(t, 200, code, desc)
	assert.Equal(t, targetID, updateResp.Result.ID, desc)

	// Проверка в БД
	var target models.Target
	f.DB.First(&target, "id = ?", targetID)
	// Имя не является изменяемым
	assert.Equal(t, createInput.Name, target.Name, desc)
	assert.Equal(t, "Updated desc", *target.Description, desc)
}

func TestV1TargetUpsertUpdate_NotEditor(t *testing.T) {
	desc := "update target by non-editor should fail"
	f := NewTestHTTP()
	authAdmin := f.AuthorizationUser(0, nil)
	regularUser := models.User{}
	regularUser.Email = "user" + uuid.New().String() + "@example.com"
	f.DB.Create(&regularUser)
	authStudent := f.AuthorizationUser(regularUser.ID, nil)

	// Создаём цель от админа
	createInput := usecases.TargetUpsertInputDTO{
		ID:          nil,
		Name:        "Target" + uuid.New().String(),
		Description: ptrString("desc"),
		Type:        "service",
	}
	code, body := f.Rpc(&TestRpcRequest{
		Method:        "target.upsert",
		Params:        createInput,
		Authorization: authAdmin,
	})
	var createResp usecases.TargetUpsertResponse
	json.Unmarshal([]byte(body), &createResp)
	targetID := createResp.Result.ID

	// Пытаемся обновить от студента (не редактор)
	updateInput := usecases.TargetUpsertInputDTO{
		ID:          &targetID,
		Name:        "Hacked" + uuid.New().String(),
		Description: ptrString("evil"),
		Type:        "service",
	}
	code, body = f.Rpc(&TestRpcRequest{
		Method:        "target.upsert",
		Params:        updateInput,
		Authorization: authStudent,
	})
	assert.Equal(t, 500, code, desc) // код ответа 500, ошибка в JSON-RPC
	var rpcError struct {
		Error struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				Code string `json:"code"`
			} `json:"data"`
		} `json:"error"`
	}
	json.Unmarshal([]byte(body), &rpcError)
	assert.Equal(t, "permission_denied", rpcError.Error.Data.Code, desc)
}

func TestV1TargetUpsertCreate_ForbiddenRole(t *testing.T) {
	desc := "create target by student should be forbidden"
	f := NewTestHTTP()
	regularUser := models.User{}
	regularUser.Email = "user" + uuid.New().String() + "@example.com"
	f.DB.Create(&regularUser)
	authStudent := f.AuthorizationUser(regularUser.ID, nil) // student

	input := usecases.TargetUpsertInputDTO{
		ID:          nil,
		Name:        "Target" + uuid.New().String(),
		Description: ptrString("desc"),
		Type:        "service",
	}
	code, body := f.Rpc(&TestRpcRequest{
		Method:        "target.upsert",
		Params:        input,
		Authorization: authStudent,
	})
	assert.Equal(t, 500, code, desc)
	var rpcError struct {
		Error struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				Code string `json:"code"`
			} `json:"data"`
		} `json:"error"`
	}
	json.Unmarshal([]byte(body), &rpcError)
	assert.Equal(t, "permission_denied", rpcError.Error.Data.Code, desc)
}

func TestV1TargetGet(t *testing.T) {
	desc := "get target by ID"
	f := NewTestHTTP()
	auth := f.AuthorizationUser(0, nil)

	target := models.Target{
		ID:          uuid.New().String(),
		Name:        "Test Target" + uuid.New().String(),
		Description: ptrString("Test Description"),
		Type:        "service",
		Links:       types.JsonStore{},
		Tags:        types.JsonStore{},
	}
	f.DB.Create(&target)

	input := usecases.TargetGetInputDTO{ID: target.ID}
	code, body := f.Rpc(&TestRpcRequest{
		Method:        "target.get",
		Params:        input,
		Authorization: auth,
	})
	var resp usecases.TargetGetResponse
	err := json.Unmarshal([]byte(body), &resp)
	assert.NoError(t, err, desc)
	assert.Equal(t, 200, code, desc)
	assert.Equal(t, target.ID, resp.Result.ID, desc)
	assert.Equal(t, target.Name, resp.Result.Name, desc)
	assert.Equal(t, *target.Description, *resp.Result.Description, desc)
}

func TestV1TargetDelete(t *testing.T) {
	desc := "delete target"
	f := NewTestHTTP()
	regularUser := models.User{}
	regularUser.Email = "user" + uuid.New().String() + "@example.com"
	f.DB.Create(&regularUser)
	authUser := f.AuthorizationUser(regularUser.ID, nil)

	// Создаём цель
	target := models.Target{
		ID:          uuid.New().String(),
		Name:        "ToDelete" + uuid.New().String(),
		Description: ptrString("desc"),
		Type:        "service",
	}
	f.DB.Create(&target)

	targetUser := models.TargetUser{
		TargetUserBase: models.TargetUserBase{
			TargetID: target.ID,
			UserID:   regularUser.ID,
			Roles: types.JsonStore{
				"roles": []string{models.UserRoleEditor},
			},
		},
	}
	f.DB.Create(&targetUser)

	input := usecases.TargetDeleteInputDTO{ID: target.ID}
	code, body := f.Rpc(&TestRpcRequest{
		Method:        "target.delete",
		Params:        input,
		Authorization: authUser,
	})
	var resp usecases.TargetDeleteResponse
	err := json.Unmarshal([]byte(body), &resp)
	assert.NoError(t, err, desc)
	assert.Equal(t, 200, code, desc)

	// Проверка удаления
	var count int64
	f.DB.Model(&models.Target{}).Where("id = ?", target.ID).Count(&count)
	assert.Equal(t, int64(0), count, desc)
}

func TestV1TargetDelete_NotEditor(t *testing.T) {
	desc := "delete target by non-editor should fail"
	f := NewTestHTTP()
	authAdmin := f.AuthorizationUser(0, nil)
	regularUser := models.User{}
	regularUser.Email = "user" + uuid.New().String() + "@example.com"
	f.DB.Create(&regularUser)
	authStudent := f.AuthorizationUser(regularUser.ID, nil)

	// Создаём цель от админа (он становится редактором)
	createInput := usecases.TargetUpsertInputDTO{
		ID:          nil,
		Name:        "Target" + uuid.New().String(),
		Description: ptrString("desc"),
		Type:        "service",
	}
	code, body := f.Rpc(&TestRpcRequest{
		Method:        "target.upsert",
		Params:        createInput,
		Authorization: authAdmin,
	})
	var createResp usecases.TargetUpsertResponse
	json.Unmarshal([]byte(body), &createResp)
	targetID := createResp.Result.ID

	// Пытаемся удалить от студента
	input := usecases.TargetDeleteInputDTO{ID: targetID}
	code, body = f.Rpc(&TestRpcRequest{
		Method:        "target.delete",
		Params:        input,
		Authorization: authStudent,
	})
	assert.Equal(t, 500, code, desc)
	var rpcError struct {
		Error struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				Code string `json:"code"`
			} `json:"data"`
		} `json:"error"`
	}
	json.Unmarshal([]byte(body), &rpcError)
	assert.Equal(t, "permission_denied", rpcError.Error.Data.Code, desc)
}

func ptrString(s string) *string {
	return &s
}

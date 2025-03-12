package routes

import (
	"fmt"
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func TestV1UNLFileListSuccess(t *testing.T) {
	description := "list UNL files successfully"
	f := NewFiberTestHTTP()

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(models.UNLFile{})

	// Create test UNL files
	entityDB1 := models.UNLFile{}
	entityDB1.Path = "/path/to/file1"
	entityDB1.Type = "txt"
	f.DB.Create(&entityDB1)

	entityDB2 := models.UNLFile{}
	entityDB2.Path = "/path/to/file2"
	entityDB2.Type = "pdf"
	f.DB.Create(&entityDB2)

	authHeader := f.AuthorizationUser(0, 0, "") // Assuming user ID 1 is authorized

	input := usecases.UNLFileListInputDTO{
		Search: "",
		Type:   []string{},
		Limit:  2,
		Offset: 0,
	}

	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/unl-file/list",
		FiberRequestPayload(input),
		authHeader,
	)
	bodyModel := usecases.UNLFileListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 2, len(bodyModel.Result.Model), description)
	assert.Equal(t, entityDB2.ID, bodyModel.Result.Model[0].ID, description)
	assert.Equal(t, entityDB2.Path, bodyModel.Result.Model[0].Path, description)
	assert.Equal(t, entityDB1.ID, bodyModel.Result.Model[1].ID, description)
	assert.Equal(t, entityDB1.Path, bodyModel.Result.Model[1].Path, description)
}

func TestV1UNLFileListFilterByType(t *testing.T) {
	description := "list UNL files with type filter"
	f := NewFiberTestHTTP()

	// Clear Table
	f.DB.Where("id > ?", 0).Delete(models.UNLFile{})

	// Create test UNL files
	entityDB1 := models.UNLFile{}
	entityDB1.Path = "/path/to/file1"
	entityDB1.Type = "txt"
	f.DB.Create(&entityDB1)

	entityDB2 := models.UNLFile{}
	entityDB2.Path = "/path/to/file2"
	entityDB2.Type = "pdf"
	f.DB.Create(&entityDB2)

	authHeader := f.AuthorizationUser(0, 0, "")

	input := usecases.UNLFileListInputDTO{
		Search: "",
		Type:   []string{"txt"},
		Limit:  10,
		Offset: 0,
	}

	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/unl-file/list",
		FiberRequestPayload(input),
		authHeader,
	)
	bodyModel := usecases.UNLFileListResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 1, len(bodyModel.Result.Model), description)
	assert.Equal(t, entityDB1.ID, bodyModel.Result.Model[0].ID, description)
	assert.Equal(t, entityDB1.Type, bodyModel.Result.Model[0].Type, description)
}

func TestV1UNLFileGetSuccess(t *testing.T) {
	description := "get UNL file successfully"
	f := NewFiberTestHTTP()

	// Create a test UNL file
	entityDB := models.UNLFile{}
	entityDB.Path = "/path/to/file"
	entityDB.Type = "txt"
	f.DB.Create(&entityDB)

	authHeader := f.AuthorizationUser(0, 0, "")

	input := usecases.UNLFileGetInputDTO{
		ID: entityDB.ID,
	}

	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/unl-file/get",
		FiberRequestPayload(input),
		authHeader,
	)
	bodyModel := usecases.UNLFileGetResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, entityDB.ID, bodyModel.Result.Model.ID, description)
	assert.Equal(t, entityDB.Path, bodyModel.Result.Model.Path, description)
	assert.Equal(t, entityDB.Type, bodyModel.Result.Model.Type, description)
}

func TestV1UNLFileGetNotFound(t *testing.T) {
	description := "get non-existent UNL file"
	f := NewFiberTestHTTP()
	authHeader := f.AuthorizationUser(0, 0, "")

	// Use a non-existent ID
	input := usecases.UNLFileGetInputDTO{
		ID: 9999,
	}

	expectedCode := 404 // Assuming 404 is returned for not found
	statusCode, body := f.Request(
		"POST",
		"/api/v1/unl-file/get",
		FiberRequestPayload(input),
		authHeader,
	)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "not found", description)
}

func TestV1UNLFileSyncSuccess(t *testing.T) {
	description := "sync UNL files successfully"
	f := NewFiberTestHTTP()

	authHeader := f.AuthorizationServiceBasic()

	input := usecases.UNLFileSyncInputDTO{
		Branch:     "main",
		Repository: "https://github.com/maintainer64/yandex-music-to-discord-macos-native.git",
	}

	expectedCode := 200
	statusCode, body := f.Request(
		"POST",
		"/api/v1/unl-file/sync",
		FiberRequestPayload(input),
		authHeader,
	)
	bodyModel := usecases.UNLFileSyncResponse{}
	_ = json.Unmarshal([]byte(body), &bodyModel)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Equal(t, 10, bodyModel.Result.Count, description)
	var entities []models.UNLFile
	f.DB.Where("deleted_at is null").Order("created_at desc").Find(&entities)
	assert.Equal(t, 10, len(entities), description)
	expectedFiles := []map[string]string{
		{
			"path": "richUpdater.go",
			"type": "go",
		},
		{
			"path": "pkg/nowplayingmacos/nowplaying.go",
			"type": "go",
		},
		{
			"path": "pkg/discordimageuploader/uploader.go",
			"type": "go",
		},
		{
			"path": "pkg/discordimageuploader/search.go",
			"type": "go",
		},
		{
			"path": "pkg/discordimageuploader/cache.go",
			"type": "go",
		},
		{
			"path": "pkg/configurator/configurator.go",
			"type": "go",
		},
		{
			"path": "pkg/cache/cache.go",
			"type": "go",
		},
		{
			"path": "main.go",
			"type": "go",
		},
		{
			"path": "go.sum",
			"type": "sum",
		},
		{
			"path": "go.mod",
			"type": "mod",
		},
	}
	for index, entity := range entities {
		assert.Equal(t, expectedFiles[index]["path"], entity.Path, fmt.Sprintf("%s_%d", description, index))
		assert.Equal(t, expectedFiles[index]["type"], entity.Type, fmt.Sprintf("%s_%d", description, index))
	}
}

func TestV1UNLFileSyncInvalidRepository(t *testing.T) {
	description := "sync UNL files with invalid repository"
	f := NewFiberTestHTTP()

	authHeader := f.AuthorizationServiceBasic()

	input := usecases.UNLFileSyncInputDTO{
		Branch:     "main",
		Repository: "invalid-repo",
	}

	expectedCode := 500 // Assuming 500 is returned for internal server error
	statusCode, body := f.Request(
		"POST",
		"/api/v1/unl-file/sync",
		FiberRequestPayload(input),
		authHeader,
	)

	assert.Equal(t, expectedCode, statusCode, description)
	assert.Contains(t, body, "not clone repository", description)
}

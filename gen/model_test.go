package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelNew(t *testing.T) {
	outputData := NewTemplateContextDTO(&CliParams{
		ProjectRoot: "github.com/simple-app/backend",
		ModelName:   "UserFormLegacy",
		Directory:   "/opt/app",
	})
	expectedData := &TemplateContextDTO{
		Root:           "github.com/simple-app/backend",
		Name:           "UserFormLegacy",
		FilesDirectory: "/opt/app",
		NameSnake:      "user_form_legacy",
		NameDash:       "user-form-legacy",
	}
	assert.Equal(t, expectedData, outputData)
}

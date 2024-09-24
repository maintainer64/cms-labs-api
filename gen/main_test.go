package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func tempDirectory() (string, error) {
	return os.MkdirTemp(".", "tmp")
}

func TestCommandLineArgumentMain(t *testing.T) {
	directory, err := tempDirectory()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(directory)
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }() // os.Args is a "global variable", so keep the state from before the test, and restore it after.

	fmt.Println("Expected: \nTwo arguments are:  3   fnameTestname")
	fmt.Println("Got:")
	os.Args = []string{"main.go", "-model", "Book", "-root", "github.com/project", "-directory", directory}
	main()
	var createdFiles []string
	err = filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				panic(err)
			}
			if !info.IsDir() {
				createdFiles = append(
					createdFiles,
					strings.TrimPrefix(path, strings.TrimPrefix(directory, "./")),
				)
			}
			return nil
		})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(
		t,
		[]string{
			"/app/controllers/book_controller.go",
			"/app/di/book_container.go",
			"/app/models/book_model.go",
			"/app/queries/book_query.go",
			"/app/usecases/book_delete.go",
			"/app/usecases/book_edit.go",
			"/app/usecases/book_get.go",
			"/app/usecases/book_list.go",
			"/pkg/routes/book_route.go",
		},
		createdFiles,
	)
}

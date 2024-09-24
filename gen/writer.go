package main

import (
	"errors"
	"os"
	"path/filepath"
)

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

func createDirectories(filePath string) error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0o750)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteFile(file *GenerateFile) error {
	if checkFileExists(file.FullFileName) {
		return nil
	}
	err := createDirectories(file.FullFileName)
	if err != nil {
		return err
	}
	fin, err := os.Create(file.FullFileName)
	if err != nil {
		return err
	}
	defer fin.Close()
	_, err = fin.WriteString(file.Content)
	return err
}

func WriteFileBulk(files []GenerateFile) error {
	for _, file := range files {
		err := WriteFile(&file)
		if err != nil {
			return err
		}
	}
	return nil
}

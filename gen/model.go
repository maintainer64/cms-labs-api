package main

import (
	"regexp"
	"strings"
)

// CliParams region
type CliParams struct {
	ProjectRoot string // путь для import проекта
	ModelName   string // название модели в БД
	Directory   string // директория для проекта
	Type        string // тип генерации
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ModelNameSnake название_маленькими_буквами в БД
func (m *CliParams) ModelNameSnake() string {
	snake := matchFirstCap.ReplaceAllString(m.ModelName, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// ModelNameDash название-маленькими-буквами в маршрутах
func (m *CliParams) ModelNameDash() string {
	dash := matchFirstCap.ReplaceAllString(m.ModelName, "${1}-${2}")
	dash = matchAllCap.ReplaceAllString(dash, "${1}-${2}")
	return strings.ToLower(dash)
}

// CliParams end region

// TemplateContextDTO region
type TemplateContextDTO struct {
	Root           string // путь для import проекта
	Name           string // название модели в БД
	FilesDirectory string // директория для проекта
	NameSnake      string // название_маленькими_буквами в БД
	NameDash       string // название-маленькими-буквами в маршрутах
}

func NewTemplateContextDTO(params *CliParams) *TemplateContextDTO {
	return &TemplateContextDTO{
		Root:           params.ProjectRoot,
		Name:           params.ModelName,
		FilesDirectory: params.Directory,
		NameSnake:      params.ModelNameSnake(),
		NameDash:       params.ModelNameDash(),
	}
}

type FileDeclaration struct {
	Filename          string
	Template          string
	RelativeDirectory string
}

// TemplateContextDTO end region

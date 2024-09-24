package main

const AppModelsTemplate = `package models

// {{.Name}}Base struct to describe {{.Name}} object.
type {{.Name}}Base struct {

}

type {{.Name}}Secret struct {

}

type {{.Name}}ListItem struct {
	Base
	{{.Name}}Base
}

// TableName переопределяет название таблицы для {{.Name}}ListItem на ` + "`{{.NameSnake}}s`" + `
func ({{.Name}}ListItem) TableName() string {
	return "{{.NameSnake}}s"
}

type {{.Name}} struct {
	Base
	{{.Name}}Base
	{{.Name}}Secret
}

`

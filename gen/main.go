// Package main автоматический генератор моделей проекта
package main

import (
	"flag"
	"fmt"
)

var files = []FileDeclaration{
	{
		Filename:          "{{.NameSnake}}_controller.go",
		Template:          AppControllersTemplate,
		RelativeDirectory: "app/controllers",
	},
	{
		Filename:          "{{.NameSnake}}_container.go",
		Template:          AppDITemplate,
		RelativeDirectory: "app/di",
	},
	{
		Filename:          "{{.NameSnake}}_model.go",
		Template:          AppModelsTemplate,
		RelativeDirectory: "app/models",
	},
	{
		Filename:          "{{.NameSnake}}_query.go",
		Template:          AppQueriesTemplate,
		RelativeDirectory: "app/queries",
	},
	{
		Filename:          "{{.NameSnake}}_route.go",
		Template:          PKGRouteTemplate,
		RelativeDirectory: "pkg/routes",
	},
	{
		Filename:          "{{.NameSnake}}_delete.go",
		Template:          PKGUseCaseDelete,
		RelativeDirectory: "app/usecases",
	},
	{
		Filename:          "{{.NameSnake}}_edit.go",
		Template:          PKGUseCaseEdit,
		RelativeDirectory: "app/usecases",
	},
	{
		Filename:          "{{.NameSnake}}_get.go",
		Template:          PKGUseCaseGet,
		RelativeDirectory: "app/usecases",
	},
	{
		Filename:          "{{.NameSnake}}_list.go",
		Template:          PKGUseCaseList,
		RelativeDirectory: "app/usecases",
	},
}

func main() {
	var params CliParams
	flag.StringVar(&params.Directory, "directory", "", "Директория проекта")
	flag.StringVar(&params.ProjectRoot, "root", "", "Путь gitlab.com/a10869/... импорта модулей")
	flag.StringVar(&params.ModelName, "model", "", "Название модели")
	flag.Parse()
	if params.ModelName == "" {
		panic("Не передана директория пакета")
	}
	fmt.Printf(
		"Generate model %s on module %s into directory %s",
		params.ModelName,
		params.ProjectRoot,
		params.Directory,
	)
	content, err := GenerateBulk(NewTemplateContextDTO(&params), files)
	if err != nil {
		panic(err)
	}
	err = WriteFileBulk(content)
	if err != nil {
		panic(err)
	}
}

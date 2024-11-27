package main

const PKGUseCaseDelete = `package usecases

import (
	"{{.Root}}/app/queries"
)

type {{.Name}}DeleteUC struct {
	{{.Name}}Queries *queries.{{.Name}}Queries
}

type {{.Name}}DeleteInputDTO struct {
	ID uint ` + "`json:\"id\" required:\"true\"`" + `
}

type {{.Name}}DeleteResponse = Response[{{.Name}}DeleteInputDTO]

func (u *{{.Name}}DeleteUC) Execute(dto {{.Name}}DeleteInputDTO) ({{.Name}}DeleteInputDTO, error) {
	err := u.{{.Name}}Queries.Delete(dto.ID)
	return dto, err
}

`

const PKGUseCaseEdit = `package usecases

import (
	"{{.Root}}/app/models"
	"{{.Root}}/app/queries"
)

type {{.Name}}EditUC struct {
	{{.Name}}Queries *queries.{{.Name}}Queries
}

type {{.Name}}EditInputDTO struct {
	ID       uint   ` + "`json:\"id\"`" + `
	Name     string ` + "`json:\"name\" validate:\"required\"`" + `
	Url      string ` + "`json:\"url\" validate:\"required\"`" + `
	IsActive bool   ` + "`json:\"is_active\"`" + `
	UnitRate uint   ` + "`json:\"unit_rate\"`" + `
	Token    string ` + "`json:\"token\" validate:\"required\"`" + `
}

type {{.Name}}EditOutputDTO struct {
	ID uint ` + "`json:\"id\" required:\"true\"`" + `
}

type {{.Name}}EditResponse = Response[{{.Name}}EditOutputDTO]

func (u *{{.Name}}EditUC) Execute(dto {{.Name}}EditInputDTO) ({{.Name}}EditOutputDTO, error) {
	entity := &models.{{.Name}}{}
	entity.ID = dto.ID
    /*
        TODO: Add attributes set to upsert {{.Name}}
	*/
	err := u.{{.Name}}Queries.Upsert(entity)
	return {{.Name}}EditOutputDTO{ID: entity.ID}, err
}

`

const PKGUseCaseGet = `package usecases

import (
	"{{.Root}}/app/models"
	"{{.Root}}/app/queries"
)

type {{.Name}}GetUC struct {
	{{.Name}}Queries *queries.{{.Name}}Queries
}

type {{.Name}}GetInputDTO struct {
	ID uint ` + "`json:\"id\" required:\"true\"`" + `
}

type {{.Name}}GetOutputDTO struct {
	Model models.{{.Name}} ` + "`json:\"model\" required:\"true\"`" + `
}

type {{.Name}}GetResponse = Response[{{.Name}}GetOutputDTO]

func (u *{{.Name}}GetUC) Execute(dto {{.Name}}GetInputDTO) ({{.Name}}GetOutputDTO, error) {
	form, err := u.{{.Name}}Queries.Get(dto.ID)
	return {{.Name}}GetOutputDTO{
		Model: form,
	}, err
}

`

const PKGUseCaseList = `package usecases

import (
	"{{.Root}}/app/models"
	"{{.Root}}/app/queries"
)

type {{.Name}}ListUC struct {
	{{.Name}}Queries *queries.{{.Name}}Queries
}

type {{.Name}}ListInputDTO struct {
	Search string ` + "`json:\"search\"`" + `
	Limit  int ` + "`json:\"limit\"`" + `
	Offset int ` + "`json:\"offset\"`" + `
}

type {{.Name}}ListOutputDTO struct {
	Model []models.{{.Name}}ListItem ` + "`json:\"model\" validate:\"required\"`" + `
	TotalCount int64                      ` + "`json:\"total_count\" validate:\"required\"`" + `
}

type {{.Name}}ListResponse = Response[{{.Name}}ListOutputDTO]

func (u *{{.Name}}ListUC) Execute(dto {{.Name}}ListInputDTO) ({{.Name}}ListOutputDTO, error) {
	entities, count, err := u.{{.Name}}Queries.List(dto.Search, dto.Limit, dto.Offset)
	return {{.Name}}ListOutputDTO{
		Model: entities,
		TotalCount: count
	}, err
}

`

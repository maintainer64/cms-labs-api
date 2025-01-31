package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFunc(t *testing.T) {
	inputData := TemplateContextDTO{
		Root:           "github.com/simple-app/backend",
		Name:           "UserFormLegacy",
		FilesDirectory: "/opt/directory",
		NameSnake:      "user_form_legacy",
		NameDash:       "user-form-legacy",
	}
	outputData, err := GenerateBulk(&inputData, files)
	assert.Nil(t, err)
	assert.Equal(t, []GenerateFile{
		{
			FullFileName: "/opt/directory/app/controllers/user_form_legacy_controller.go",
			Content: `package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simple-app/backend/app/di"
	"github.com/simple-app/backend/app/usecases"
	"github.com/simple-app/backend/pkg/response"
)

// UserFormLegacyCreate func for creates a new UserFormLegacy.
// @Description Create user_form_legacy.
// @Summary create user_form_legacy
// @Tags UserFormLegacy
// @Accept json
// @Produce json
// @Param form body usecases.UserFormLegacyEditInputDTO true "user_form_legacy form info"
// @Success 200 {object} usecases.UserFormLegacyEditResponse
// @Security ApiKeyAuth
// @Router /v1/user-form-legacy/upsert [post]
func UserFormLegacyCreate(c *fiber.Ctx) error {
	dto := usecases.UserFormLegacyEditInputDTO{}
	err := response.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().UserFormLegacyEditUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return response.FiberSuccessResponse{Result: output}
}

// UserFormLegacyList func for view of list UserFormLegacy.
// @Description List user_form_legacy.
// @Summary list user_form_legacy
// @Tags UserFormLegacy
// @Accept json
// @Produce json
// @Param form body usecases.UserFormLegacyListInputDTO true "user_form_legacy list info"
// @Success 200 {object} usecases.UserFormLegacyListResponse
// @Security ApiKeyAuth
// @Router /v1/user-form-legacy/list [post]
func UserFormLegacyList(c *fiber.Ctx) error {
	dto := usecases.UserFormLegacyListInputDTO{}
	err := response.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().UserFormLegacyListUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return response.FiberSuccessResponse{Result: output}
}

// UserFormLegacyDelete func for delete UserFormLegacy.
// @Description Delete user_form_legacy.
// @Summary delete user_form_legacy
// @Tags UserFormLegacy
// @Accept json
// @Produce json
// @Param form body usecases.UserFormLegacyDeleteInputDTO true "user_form_legacy id"
// @Success 200 {object} usecases.UserFormLegacyDeleteResponse
// @Security ApiKeyAuth
// @Router /v1/user-form-legacy/delete [post]
func UserFormLegacyDelete(c *fiber.Ctx) error {
	dto := usecases.UserFormLegacyDeleteInputDTO{}
	err := response.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().UserFormLegacyDeleteUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return response.FiberSuccessResponse{Result: output}
}

// UserFormLegacyGet func for full model UserFormLegacy.
// @Description get user_form_legacy.
// @Summary get user_form_legacy
// @Tags UserFormLegacy
// @Accept json
// @Produce json
// @Param form body usecases.UserFormLegacyGetInputDTO true "user_form_legacy id"
// @Success 200 {object} usecases.UserFormLegacyGetResponse
// @Security ApiKeyAuth
// @Router /v1/user-form-legacy/get [post]
func UserFormLegacyGet(c *fiber.Ctx) error {
	dto := usecases.UserFormLegacyGetInputDTO{}
	err := response.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	uc, err := di.NewDIContainer().UserFormLegacyGetUC()
	if err != nil {
		return err
	}
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return response.FiberSuccessResponse{Result: output}
}

`,
		},
		{
			FullFileName: "/opt/directory/app/di/user_form_legacy_container.go",
			Content: `package di

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simple-app/backend/app/usecases"
	"github.com/simple-app/backend/pkg/response"
)

func (di *DIContainer) UserFormLegacyEditUC() (*usecases.UserFormLegacyEditUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, response.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.UserFormLegacyEditUC{
		UserFormLegacyQueries: db.UserFormLegacyQueries,
	}, nil
}

func (di *DIContainer) UserFormLegacyGetUC() (*usecases.UserFormLegacyGetUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, response.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.UserFormLegacyGetUC{
		UserFormLegacyQueries: db.UserFormLegacyQueries,
	}, nil
}

func (di *DIContainer) UserFormLegacyListUC() (*usecases.UserFormLegacyListUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, response.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.UserFormLegacyListUC{
		UserFormLegacyQueries: db.UserFormLegacyQueries,
	}, nil
}

func (di *DIContainer) UserFormLegacyDeleteUC() (*usecases.UserFormLegacyDeleteUC, error) {
	db, err := di.Queries()
	if err != nil {
		return nil, response.FiberValidationException{Status: fiber.StatusInternalServerError, Exception: err}
	}
	return &usecases.UserFormLegacyDeleteUC{
		UserFormLegacyQueries: db.UserFormLegacyQueries,
	}, nil
}

`,
		},
		{
			FullFileName: "/opt/directory/app/models/user_form_legacy_model.go",
			Content: `package models

// UserFormLegacyBase struct to describe UserFormLegacy object.
type UserFormLegacyBase struct {

}

type UserFormLegacySecret struct {

}

type UserFormLegacyListItem struct {
	Base
	UserFormLegacyBase
}

// TableName переопределяет название таблицы для UserFormLegacyListItem на ` + "`user_form_legacys`" + `
func (UserFormLegacyListItem) TableName() string {
	return "user_form_legacys"
}

type UserFormLegacy struct {
	Base
	UserFormLegacyBase
	UserFormLegacySecret
}

`,
		},
		{
			FullFileName: "/opt/directory/app/queries/user_form_legacy_query.go",
			Content: `package queries

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/simple-app/backend/app/models"
	"github.com/simple-app/backend/pkg/response"
	"gorm.io/gorm"
	"time"
)

type UserFormLegacyQueries struct {
	*gorm.DB
}

func (q *UserFormLegacyQueries) Get(id uint) (models.UserFormLegacy, error) {
	var entity models.UserFormLegacy
	result := q.First(&entity, id)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return entity, response.FiberValidationException{
			Status:    fiber.StatusNotFound,
			Exception: errors.New("UserFormLegacy not found"),
		}
	}
	return entity, result.Error
}

func (q *UserFormLegacyQueries) Upsert(entity *models.UserFormLegacy) error {
	if entity == nil {
		return nil
	}
	entityDB := models.UserFormLegacy{}
	q.Where("id = ?", entity.ID).Find(&entityDB)
	if entityDB.ID != 0 {
		// Update
		log.Debug().Msg(fmt.Sprintf("UserFormLegacyQueries: entity update: %+v", entityDB))
		log.Info().Msg(fmt.Sprintf("UserFormLegacyQueries: entity update id=%+v", entityDB.ID))
		entity.ID = entityDB.ID
		entity.CreatedAt = entityDB.CreatedAt
		entity.UpdatedAt = time.Now().UTC()
        /*
            TODO: Complete code with UserFormLegacy entity
        */
		result := q.Save(&entity)
		return result.Error
	} else {
		// Create
		log.Debug().Msg(fmt.Sprintf("UserFormLegacyQueries: entity create: %+v", entity))
		log.Info().Msg(fmt.Sprintf("UserFormLegacyQueries: entity create name=%+v", entity.Name))
		entity.ID = 0
		entity.CreatedAt = time.Now().UTC()
		entity.UpdatedAt = time.Now().UTC()
		/*
		    Complete code with UserFormLegacy entity
		*/
		result := q.Create(entity)
		return result.Error
	}
}

func (q *UserFormLegacyQueries) List(
	search string,
	limit int,
	offset int,
) ([]models.UserFormLegacyListItem, int64, error) {
	var entities []models.UserFormLegacyListItem
	result := q.listFilter(
		search,
		q.Limit(MaxLimitCount).Offset(0),
	).Find(&entities)
	count := result.RowsAffected
	log.Debug().Msg(fmt.Sprintf("UserFormLegacyQueries list: count %+v", count))
	result = q.listFilter(
		search,
		q.Limit(limit).Offset(offset),
	).Find(&entities)
	log.Debug().Msg(fmt.Sprintf("UserFormLegacyQueries list: entities %+v", entities))
	return entities, count, result.Error
}

func (q *UserFormLegacyQueries) listFilter(search string, tx *gorm.DB) *gorm.DB {
	tx = tx.Order(` + "`created_at desc`" + `)
	if search == "" {
		return tx
	}
	tx = tx.Or("id = ?", search)
	return tx
}

func (q *UserFormLegacyQueries) Delete(id uint) error {
	_ = q.Where("id = ?", id).Delete(&models.UserFormLegacy{})
	log.Debug().Msg(fmt.Sprintf("UserFormLegacyQueries: delete entity by id: %+v", id))
	return nil
}

`,
		},
		{
			FullFileName: "/opt/directory/pkg/routes/user_form_legacy_route.go",
			Content: `package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simple-app/backend/app/controllers"
)

func V1UserFormLegacyRoutes(a *fiber.App) {
	group := a.Group("/api/v1/user-form-legacy")
	group.Post("/upsert", controllers.UserFormLegacyCreate)
	group.Post("/list", controllers.UserFormLegacyList)
	group.Post("/delete", controllers.UserFormLegacyDelete)
	group.Post("/get", controllers.UserFormLegacyGet)
}

`,
		},
		{
			FullFileName: "/opt/directory/app/usecases/user_form_legacy_delete.go",
			Content: `package usecases

import (
	"github.com/simple-app/backend/app/queries"
)

type UserFormLegacyDeleteUC struct {
	UserFormLegacyQueries *queries.UserFormLegacyQueries
}

type UserFormLegacyDeleteInputDTO struct {
	ID uint ` + "`json:\"id\" required:\"true\"`" + `
}

type UserFormLegacyDeleteResponse = response.Response[UserFormLegacyDeleteInputDTO]

func (u *UserFormLegacyDeleteUC) Execute(dto UserFormLegacyDeleteInputDTO) (UserFormLegacyDeleteInputDTO, error) {
	err := u.UserFormLegacyQueries.Delete(dto.ID)
	return dto, err
}

`,
		},
		{
			FullFileName: "/opt/directory/app/usecases/user_form_legacy_edit.go",
			Content: `package usecases

import (
	"github.com/simple-app/backend/app/models"
	"github.com/simple-app/backend/app/queries"
)

type UserFormLegacyEditUC struct {
	UserFormLegacyQueries *queries.UserFormLegacyQueries
}

type UserFormLegacyEditInputDTO struct {
	ID       uint   ` + "`json:\"id\"`" + `
	Name     string ` + "`json:\"name\" validate:\"required\"`" + `
	Url      string ` + "`json:\"url\" validate:\"required\"`" + `
	IsActive bool   ` + "`json:\"is_active\"`" + `
	UnitRate uint   ` + "`json:\"unit_rate\"`" + `
	Token    string ` + "`json:\"token\" validate:\"required\"`" + `
}

type UserFormLegacyEditOutputDTO struct {
	ID uint ` + "`json:\"id\" required:\"true\"`" + `
}

type UserFormLegacyEditResponse = response.Response[UserFormLegacyEditOutputDTO]

func (u *UserFormLegacyEditUC) Execute(dto UserFormLegacyEditInputDTO) (UserFormLegacyEditOutputDTO, error) {
	entity := &models.UserFormLegacy{}
	entity.ID = dto.ID
    /*
        TODO: Add attributes set to upsert UserFormLegacy
	*/
	err := u.UserFormLegacyQueries.Upsert(entity)
	return UserFormLegacyEditOutputDTO{ID: entity.ID}, err
}

`,
		},
		{
			FullFileName: "/opt/directory/app/usecases/user_form_legacy_get.go",
			Content: `package usecases

import (
	"github.com/simple-app/backend/app/models"
	"github.com/simple-app/backend/app/queries"
)

type UserFormLegacyGetUC struct {
	UserFormLegacyQueries *queries.UserFormLegacyQueries
}

type UserFormLegacyGetInputDTO struct {
	ID uint ` + "`json:\"id\" required:\"true\"`" + `
}

type UserFormLegacyGetOutputDTO struct {
	Model models.UserFormLegacy ` + "`json:\"model\" required:\"true\"`" + `
}

type UserFormLegacyGetResponse = response.Response[UserFormLegacyGetOutputDTO]

func (u *UserFormLegacyGetUC) Execute(dto UserFormLegacyGetInputDTO) (UserFormLegacyGetOutputDTO, error) {
	form, err := u.UserFormLegacyQueries.Get(dto.ID)
	return UserFormLegacyGetOutputDTO{
		Model: form,
	}, err
}

`,
		},
		{
			FullFileName: "/opt/directory/app/usecases/user_form_legacy_list.go",
			Content: `package usecases

import (
	"github.com/simple-app/backend/app/models"
	"github.com/simple-app/backend/app/queries"
)

type UserFormLegacyListUC struct {
	UserFormLegacyQueries *queries.UserFormLegacyQueries
}

type UserFormLegacyListInputDTO struct {
	Search string ` + "`json:\"search\"`" + `
	Limit  int ` + "`json:\"limit\"`" + `
	Offset int ` + "`json:\"offset\"`" + `
}

type UserFormLegacyListOutputDTO struct {
	Model []models.UserFormLegacyListItem ` + "`json:\"model\" validate:\"required\"`" + `
	TotalCount int64                      ` + "`json:\"total_count\" validate:\"required\"`" + `
}

type UserFormLegacyListResponse = response.Response[UserFormLegacyListOutputDTO]

func (u *UserFormLegacyListUC) Execute(dto UserFormLegacyListInputDTO) (UserFormLegacyListOutputDTO, error) {
	entities, count, err := u.UserFormLegacyQueries.List(dto.Search, dto.Limit, dto.Offset)
	return UserFormLegacyListOutputDTO{
		Model: entities,
		TotalCount: count,
	}, err
}

`,
		},
	}, outputData)
}

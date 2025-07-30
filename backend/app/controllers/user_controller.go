package controllers

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/backend/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// UserCreate func for creates a new User.
// @Description Create user. Roles[admin, instructor]
// @Summary create user
// @Tags User
// @Accept json
// @Produce json
// @Param form body usecases.UserEditInputDTO true "user form info"
// @Success 200 {object} usecases.UserEditResponse
// @Security ApiKeyAuth
// @Router /v1/user/upsert [post]
func UserCreate(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return err
	}
	dto := usecases.UserEditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.UserEditUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// UserList func for view of list User.
// @Description List user. Roles[admin, instructor]
// @Summary list user
// @Tags User
// @Accept json
// @Produce json
// @Param form body usecases.UserListInputDTO true "user list info"
// @Success 200 {object} usecases.UserListResponse
// @Security ApiKeyAuth
// @Router /v1/user/list [post]
func UserList(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return err
	}
	dto := usecases.UserListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.UserListUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// UserGet func for full model User.
// @Description get user. Roles[admin, instructor]
// @Summary get user
// @Tags User
// @Accept json
// @Produce json
// @Param form body usecases.UserGetInputDTO true "user id"
// @Success 200 {object} usecases.UserGetResponse
// @Security ApiKeyAuth
// @Router /v1/user/get [post]
func UserGet(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return err
	}
	dto := usecases.UserGetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.UserGetUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

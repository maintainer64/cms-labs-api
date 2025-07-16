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

// RoleUpsert func for create and update a few Role.
// @Description Upset role. Roles[admin]
// @Summary upsert role
// @Tags Role
// @Accept json
// @Produce json
// @Param form body usecases.RoleEditInputDTO true "user form info"
// @Success 200 {object} usecases.RoleEditResponse
// @Security ApiKeyAuth
// @Router /v1/role/upsert [post]
func RoleUpsert(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.RoleEditInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.RoleEditUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// RoleList func for view of list Role.
// @Description List user. Roles[admin, instructor, student, any]
// @Summary list role
// @Tags Role
// @Accept json
// @Produce json
// @Param form body usecases.RoleListInputDTO true "user list info"
// @Success 200 {object} usecases.RoleListResponse
// @Router /v1/role/list [post]
func RoleList(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.RoleListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.RoleListUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// RoleDelete func for delete Role.
// @Description Delete role. Roles: [admin]
// @Summary delete role
// @Tags Role
// @Accept json
// @Produce json
// @Param form body usecases.RoleDeleteInputDTO true "pnet_server id"
// @Success 200 {object} usecases.RoleDeleteResponse
// @Security ApiKeyAuth
// @Router /v1/role/delete [post]
func RoleDelete(c *fiber.Ctx) error {
	diLoggerConf := logs.NewZeroLoggerConf(c)
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin},
	); err != nil {
		return err
	}
	dto := usecases.RoleDeleteInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.RoleDeleteUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

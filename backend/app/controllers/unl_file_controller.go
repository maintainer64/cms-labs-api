package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/backend/app/di"
	"gitlab.com/a10869/api-modules/backend/app/usecases"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// UNLFileList func for view of list UNLFile.
// @Description List unl_file.
// @Summary list unl_file
// @Tags UNLFile
// @Accept json
// @Produce json
// @Param form body usecases.UNLFileListInputDTO true "unl_file list info"
// @Success 200 {object} usecases.UNLFileListResponse
// @Security ApiKeyAuth
// @Router /v1/unl-file/list [post]
func UNLFileList(c *fiber.Ctx) error {
	dto := usecases.UNLFileListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.UNLFileListUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// UNLFileGet func for full model UNLFile.
// @Description get unl_file.
// @Summary get unl_file
// @Tags UNLFile
// @Accept json
// @Produce json
// @Param form body usecases.UNLFileGetInputDTO true "unl_file id"
// @Success 200 {object} usecases.UNLFileGetResponse
// @Security ApiKeyAuth
// @Router /v1/unl-file/get [post]
func UNLFileGet(c *fiber.Ctx) error {
	dto := usecases.UNLFileGetInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.UNLFileGetUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// UNLFileSync func for sync from git models UNLFile.
// @Description sync from git unl_file.
// @Summary sync from git unl_file
// @Tags UNLFile, EXTERNAL
// @Accept json
// @Produce json
// @Param form body usecases.UNLFileSyncInputDTO true "sync params"
// @Success 200 {object} usecases.UNLFileSyncResponse
// @Security ApiKeyAuth
// @Router /v1/unl-file/sync [post]
func UNLFileSync(c *fiber.Ctx) error {
	dto := usecases.UNLFileSyncInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.UNLFileSyncUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

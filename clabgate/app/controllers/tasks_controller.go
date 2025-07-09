package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/clabgate/app/di"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// TasksList func for view of list tasks.
// @Description List tasks. Roles: [admin, instructor]
// @Summary list tasks
// @Tags Tasks
// @Accept json
// @Produce json
// @Param form body usecases.TasksListInputDTO true "tasks list info"
// @Success 200 {object} usecases.TasksListResponse
// @Security ApiKeyAuth
// @Router /v1/tasks/list [post]
func TasksList(c *fiber.Ctx) error {
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.TasksListInputDTO{}
	err := utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc := container.TasksListUC()
	output, err := uc.Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

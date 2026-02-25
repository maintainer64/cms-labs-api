package controllers

import (
	"gitlab.com/a10869/api-modules/clabgate/app/di"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// TaskList func for view of list tasks.
// @Description List tasks. Roles: [admin, instructor]
// @Summary list tasks
// @Tags task
// @Accept json
// @Produce json
// @Param object body usecases.TaskListRequest true "tasks list info"
// @Success 200 {object} usecases.TaskListResponse
// @Security ApiKeyAuth
// @Router /clabgate/api/v1/rpc/task.list [post]
func TaskList(c *jsonrpc.Ctx) (interface{}, error) {
	if _, err := auth.ExtractTokenMetadata(
		c,
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
	); err != nil {
		return nil, err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.TaskListInputDTO{}
	err := jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc := container.TasksListUC()
	output, err := uc.Execute(dto)
	return output, err
}

package controllers

import (
	"gitlab.com/a10869/api-modules/clabgate/app/di"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// ContainerGet func for device describe by name.
// @Description Describe device topology
// @Summary describe device topology
// @Tags Device
// @Accept json
// @Produce json
// @Param form body usecases.ContainerGetRequest true "topology namespace"
// @Success 200 {object} usecases.ContainersGetResponse
// @Security ApiKeyAuth
// @Router /clabgate/api/v1/rpc/container.get [post]
func ContainerGet(c *jsonrpc.Ctx) (interface{}, error) {
	user, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return nil, err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.ContainerGetInputDTO{}
	err = jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc, err := container.ContainersGetUC()
	if err != nil {
		return nil, err
	}
	output, err := uc.SetContext(user).Execute(dto)
	return output, err
}

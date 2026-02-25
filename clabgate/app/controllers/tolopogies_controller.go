package controllers

import (
	"gitlab.com/a10869/api-modules/clabgate/app/di"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/auth"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gitlab.com/a10869/api-modules/shared/logs"
)

// TopologyGet func for view of list tasks.
// @Description List tasks. Roles: [student, admin, instructor]
// @Summary list tasks
// @Tags Topology
// @Accept json
// @Produce json
// @Param form body usecases.TopologiesGetRequest true "topology namespace"
// @Success 200 {object} usecases.TopologiesGetResponse
// @Security ApiKeyAuth
// @Router /clabgate/api/v1/rpc/topology.get [post]
func TopologyGet(c *jsonrpc.Ctx) (interface{}, error) {
	user, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return nil, err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.TopologiesGetInputDTO{}
	err = jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc, err := container.TopologiesGetUC()
	if err != nil {
		return nil, err
	}
	output, err := uc.SetContext(user).Execute(dto)
	return output, err
}

// TopologyCreate func for create personal topologies.
// @Description create personal topology. Roles: [student, admin, instructor]
// @Summary create personal topologies
// @Tags Topology
// @Accept json
// @Produce json
// @Param form body usecases.TopologyCreateRequest true "create params"
// @Success 200 {object} usecases.TopologyCreateResponse
// @Security ApiKeyAuth
// @Router /clabgate/api/v1/rpc/topology.create [post]
func TopologyCreate(c *jsonrpc.Ctx) (interface{}, error) {
	user, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return nil, err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.TopologyCreateInputDTO{}
	err = jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc, err := container.TopologiesCreateUC()
	if err != nil {
		return nil, err
	}
	output, err := uc.SetContext(user).Execute(dto)
	return output, err
}

// TopologyDelete func for delete personal topologies.
// @Description delete personal topology. Roles: [student, admin, instructor]
// @Summary delete personal topologies
// @Tags Topology
// @Accept json
// @Produce json
// @Param form body usecases.TopologyDeleteRequest true "delete params"
// @Success 200 {object} usecases.TopologyDeleteResponse
// @Security ApiKeyAuth
// @Router /clabgate/api/v1/rpc/topology.delete [post]
func TopologyDelete(c *jsonrpc.Ctx) (interface{}, error) {
	user, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return nil, err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.TopologyDeleteInputDTO{}
	err = jsonrpc.ValidatorBase(c, &dto)
	if err != nil {
		return nil, err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return nil, err
	}
	defer container.Close()
	uc, err := container.TopologiesDeleteUC()
	if err != nil {
		return nil, err
	}
	output, err := uc.SetContext(user).Execute(dto)
	return output, err
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/clabgate/app/di"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/auth"
	"gitlab.com/a10869/api-modules/clabgate/pkg/configs"
	"gitlab.com/a10869/api-modules/shared/logs"
	"gitlab.com/a10869/api-modules/shared/utils"
)

// TokenAccessJson func for get json of user cluster.
// @Description Token cluster. Roles: [student, admin, instructor]
// @Summary token cluster json
// @Tags Token
// @Accept json
// @Produce json
// @Param form body usecases.TokenAccessGetInputDTO true "params"
// @Success 200 {object} usecases.TokenAccessGetResponse
// @Security ApiKeyAuth
// @Router /v1/tokens/json [get]
func TokenAccessJson(c *fiber.Ctx) error {
	user, err := auth.ExtractTokenMetadata(c, []string{})
	if err != nil {
		return err
	}
	diLoggerConf := logs.NewZeroLoggerConf(c)
	dto := usecases.TokenAccessGetInputDTO{}
	err = utils.FiberValidatorBase(c, &dto)
	if err != nil {
		return err
	}
	container, err := di.NewDIContainer(diLoggerConf)
	if err != nil {
		return err
	}
	defer container.Close()
	uc, err := container.TokenAccessGetUC()
	if err != nil {
		return err
	}
	output, err := uc.SetContext(user).Execute(dto)
	if err != nil {
		return err
	}
	return utils.FiberSuccessResponse{Result: output}
}

// TokenAccessYaml func for get yaml of user cluster.
// @Description Token cluster. Roles: any
// @Summary token cluster yaml
// @Tags Token
// @Accept json
// @Produce json
// @Param form body usecases.TokenFileYAMLGetInputDTO true "params"
// @Success 200 {object} usecases.TokenFileYAMLGetResponse
// @Router /v1/tokens/yaml [get]
func TokenAccessYaml(c *fiber.Ctx) error {
	fileContent := configs.AppConfig.K8S.KrewConfigYaml
	filename := "kubeconfig.yaml"

	// Set headers and send file
	c.Set(fiber.HeaderContentDisposition, `attachment; filename="`+filename+`"`)
	c.Set(fiber.HeaderContentType, "application/octet-stream")
	return c.Send(configs.AppConfig.K8S.KrewConfigYaml)
}

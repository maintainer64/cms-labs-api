package controllers

import (
	fiber "github.com/gofiber/fiber/v2"
	"gitlab.com/a10869/api-modules/clabgate/pkg/configs"
)

// TokenAccessYaml func for get yaml of user cluster.
// @Description Token cluster. Roles: any
// @Summary token cluster yaml
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /clabgate/api/v1/tokens/yaml [get]
func TokenAccessYaml(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentDisposition, `attachment; filename="kubeconfig.yaml"`)
	c.Set(fiber.HeaderContentType, "application/octet-stream")
	return c.Send([]byte(configs.AppConfig.K8S.KrewConfigYaml))
}

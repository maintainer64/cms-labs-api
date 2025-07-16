package usecases

import (
	"context"
	"errors"
	"fmt"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/clabgate/app/queries/topology"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/response"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/utils"
)

type TopologiesGetUC struct {
	*zerolog.Logger
	KubernetesAdminQuery *queries.KubernetesAdminQuery
	user                 *cms_client.SSOTokenPublicData
}

type TopologiesGetInputDTO struct {
	Namespace string `json:"namespace"`
}

type TopologiesGetOutputDTO struct {
	Topology *topology.Topology `json:"topology"`
	WebUrl   string             `json:"web_url"`
}

type TopologiesGetResponse = response.Response[TopologiesGetOutputDTO]

func (u *TopologiesGetUC) SetContext(user *cms_client.SSOTokenPublicData) *TopologiesGetUC {
	u.user = user
	return u
}

func (u *TopologiesGetUC) Execute(dto TopologiesGetInputDTO) (TopologiesGetOutputDTO, error) {
	if u.user == nil {
		return TopologiesGetOutputDTO{}, errors.New("not logged in")
	}
	username := u.KubernetesAdminQuery.NormalizeEntityName(u.user.Username)
	namespace := u.KubernetesAdminQuery.NormalizeEntityName(dto.Namespace)
	if !cms_client.SSOHasIntersection(
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
		u.user.Roles,
	) && !strings.HasPrefix(namespace, username+"-") {
		return TopologiesGetOutputDTO{}, utils.FiberValidationException{
			Status:    fiber.StatusForbidden,
			Exception: errors.New("User not allow current namespace"),
		}
	}
	ctx := context.Background()
	u.Logger.Error().Msg(fmt.Sprintf("TopologiesGetUC: GetTopologyYAML namespace: %v", dto.Namespace))
	yamlContent, err := u.KubernetesAdminQuery.GetTopologyYAML(ctx, dto.Namespace)
	if err != nil {
		u.Logger.Error().Msg(fmt.Sprintf("TopologiesGetUC: GetTopologyYAML error: %v", err))
		return TopologiesGetOutputDTO{}, err
	}
	u.Logger.Debug().Msg(fmt.Sprintf("TopologiesGetUC: GetTopologyYAML content: %v", string(yamlContent)))
	u.Logger.Info().Msg(fmt.Sprintf("TopologiesGetUC: GetTopologyYAML content: %v", string(yamlContent)[:100]))
	topologyContent, err := topology.Convert(yamlContent)
	if err != nil {
		u.Logger.Error().Msg(fmt.Sprintf("TopologiesGetUC: Unmarshal topology error: %v by namespace: %v", err, dto.Namespace))
		return TopologiesGetOutputDTO{}, err
	}
	webUrl, _ := u.KubernetesAdminQuery.GetSecretByName(ctx, namespace, queries.GitlabWebUrlDeploy)
	return TopologiesGetOutputDTO{
		Topology: topologyContent,
		WebUrl:   webUrl,
	}, err
}

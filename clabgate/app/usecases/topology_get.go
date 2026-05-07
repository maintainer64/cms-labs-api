package usecases

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
	"gopkg.in/yaml.v3"
)

type TopologiesGetUC struct {
	*zerolog.Logger
	KubernetesAdminQuery *queries.KubernetesAdminQuery
	user                 *cms_client.SSOTokenPublicData
}

type TopologiesGetInputDTO struct {
	Namespace string `json:"namespace"`
}

type TopologiesGetRequest struct {
	JSONRPC string                `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string                `json:"method" default:"topology.get" required:"true"`
	Params  TopologiesGetInputDTO `json:"params,omitempty"`
	ID      string                `json:"id,omitempty" default:"1" required:"true"`
}

type TopologiesGetOutputDTO struct {
	Topology    string                   `json:"topology"`
	Deployments []queries.DeploymentInfo `json:"deployments,omitempty"`
	Services    []queries.ServiceInfo    `json:"services,omitempty"`
}

type TopologiesGetResponse struct {
	JSONRPC string                 `json:"jsonrpc" default:"2.0" required:"true"`
	Result  TopologiesGetOutputDTO `json:"result,omitempty"`
	Error   interface{}            `json:"error,omitempty"`
	ID      string                 `json:"id,omitempty" default:"1" required:"true"`
}

func (u *TopologiesGetUC) SetContext(user *cms_client.SSOTokenPublicData) *TopologiesGetUC {
	u.user = user
	return u
}

func (u *TopologiesGetUC) Execute(dto TopologiesGetInputDTO) (TopologiesGetOutputDTO, error) {
	if u.user == nil {
		return TopologiesGetOutputDTO{}, errors.New("not logged in")
	}
	if !cms_client.SSOHasIntersection(
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
		u.user.Roles,
	) && !strings.HasPrefix(dto.Namespace, "jup-"+u.user.Username+"-") {
		return TopologiesGetOutputDTO{}, jsonrpc.NewRpcError("user_not_allow_topology", "user not allow connect topology")
	}
	ctx := context.Background()
	u.Logger.Info().Msg(fmt.Sprintf("TopologiesGetUC: GetTopologyYAML namespace: %v", dto.Namespace))
	yamlContent, err := u.KubernetesAdminQuery.GetTopologyYAML(ctx, dto.Namespace)

	var containerlabContent string
	if err == nil && len(yamlContent) > 0 {
		var topoData map[string]interface{}
		if err := yaml.Unmarshal(yamlContent, &topoData); err == nil {
			if spec, ok := topoData["spec"].(map[string]interface{}); ok {
				if definition, ok := spec["definition"].(map[string]interface{}); ok {
					if clab, ok := definition["containerlab"].(string); ok {
						containerlabContent = clab
					}
				}
			}
		}
	} else {
		u.Logger.Error().Msg(fmt.Sprintf("TopologiesGetUC: GetTopologyYAML error: %v", err))
	}

	u.Logger.Debug().Msg(fmt.Sprintf("TopologiesGetUC: containerlab content length: %d", len(containerlabContent)))

	deployments, _ := u.KubernetesAdminQuery.GetDeploymentsInfo(ctx, dto.Namespace)
	services, _ := u.KubernetesAdminQuery.GetServicesInfo(ctx, dto.Namespace)

	return TopologiesGetOutputDTO{
		Topology:    containerlabContent,
		Deployments: deployments,
		Services:    services,
	}, nil
}

package usecases

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/clabgate/app/queries/topology"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
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
	Topology *topology.Topology `json:"topology"`
	WebUrl   string             `json:"web_url"`
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
	username := u.KubernetesAdminQuery.NormalizeEntityName(u.user.Username)
	namespace := u.KubernetesAdminQuery.NormalizeEntityName(dto.Namespace)
	if !cms_client.SSOHasIntersection(
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
		u.user.Roles,
	) && !strings.HasPrefix(namespace, username+"-") {
		return TopologiesGetOutputDTO{}, jsonrpc.NewRpcError("user_not_allow_topology", "user not allow connect topology")
	}
	ctx := context.Background()
	u.Logger.Error().Msg(fmt.Sprintf("TopologiesGetUC: GetTopologyYAML namespace: %v", dto.Namespace))
	yamlContent, err := u.KubernetesAdminQuery.GetTopologyYAML(ctx, dto.Namespace)
	if err != nil {
		u.Logger.Error().Msg(fmt.Sprintf("TopologiesGetUC: GetTopologyYAML error: %v", err))
	}
	u.Logger.Debug().Msg(fmt.Sprintf("TopologiesGetUC: GetTopologyYAML content: %v", string(yamlContent)))
	u.Logger.Info().Msg("TopologiesGetUC: GetTopologyYAML content")
	topologyContent, err := topology.Convert(yamlContent)
	if err != nil {
		u.Logger.Error().Msg(fmt.Sprintf("TopologiesGetUC: Unmarshal topology error: %v by namespace: %v", err, dto.Namespace))
	}
	webUrl, _ := u.KubernetesAdminQuery.GetSecretByName(ctx, namespace, queries.GitlabWebUrlDeploy)
	return TopologiesGetOutputDTO{
		Topology: topologyContent,
		WebUrl:   webUrl,
	}, nil
}

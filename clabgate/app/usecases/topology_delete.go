package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
)

type TopologyDeleteUC struct {
	*zerolog.Logger
	KubernetesAdminQuery *queries.KubernetesAdminQuery
	user                 *cms_client.SSOTokenPublicData
}

type TopologyDeleteInputDTO struct {
	Namespaces []string `json:"namespaces"`
}

type TopologyDeleteRequest struct {
	JSONRPC string                 `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string                 `json:"method" default:"topology.delete" required:"true"`
	Params  TopologyDeleteInputDTO `json:"params,omitempty"`
	ID      string                 `json:"id,omitempty" default:"1" required:"true"`
}

type TopologyDeleteOutputDTO struct {
	Namespaces []string `json:"namespaces"`
}

type TopologyDeleteResponse struct {
	JSONRPC string                  `json:"jsonrpc" default:"2.0" required:"true"`
	Result  TopologyDeleteOutputDTO `json:"result,omitempty"`
	Error   interface{}             `json:"error,omitempty"`
	ID      string                  `json:"id,omitempty" default:"1" required:"true"`
}

func (u *TopologyDeleteUC) SetContext(user *cms_client.SSOTokenPublicData) *TopologyDeleteUC {
	u.user = user
	return u
}

func (u *TopologyDeleteUC) Execute(dto TopologyDeleteInputDTO) (TopologyDeleteOutputDTO, error) {
	output := TopologyDeleteOutputDTO{}
	if u.user == nil {
		return output, errors.New("not logged in")
	}
	username := u.KubernetesAdminQuery.NormalizeEntityName(u.user.Username)
	if !cms_client.SSOHasIntersection(
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
		u.user.Roles,
	) {
		return output, jsonrpc.NewRpcError("user_not_allow_topology", "user not allow connect topology")
	}
	ctx := context.Background()
	for _, namespace := range dto.Namespaces {
		ns, err := u.KubernetesAdminQuery.GetNamespaceByName(
			ctx,
			u.KubernetesAdminQuery.NormalizeEntityName(namespace),
		)
		if err != nil {
			continue
		}
		owner, ok := ns.Labels["owner"]
		if !ok || owner == "" {
			u.Logger.Info().Msg(fmt.Sprintf("Namespace %s is not owned", ns.Name))
			continue
		}
		if owner != username && !cms_client.SSOHasIntersection(
			[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
			u.user.Roles,
		) {
			u.Logger.Info().Msg(fmt.Sprintf("Namespace %s is not deleted by %s", ns.Name, username))
			continue
		}
		_ = u.KubernetesAdminQuery.DeleteNamespaceByName(ctx, ns.Name)
		output.Namespaces = append(output.Namespaces, ns.Name)
	}
	return output, nil
}

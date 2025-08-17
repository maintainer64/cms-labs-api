package usecases

import (
	"context"
	"errors"
	"strings"

	"gitlab.com/a10869/api-modules/clabgate/app/queries/topology"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/response"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/utils"
)

type ContainersGetUC struct {
	*zerolog.Logger
	KubernetesAdminQuery     *queries.KubernetesAdminQuery
	KubeDashboardClientQuery *queries.KubeDashboardClientQuery
	user                     *cms_client.SSOTokenPublicData
}

type ContainersGetInputDTO struct {
	Namespace  string `json:"namespace"`
	Deployment string `json:"deployment"`
}

type ContainersGetItem struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Pod   string `json:"pod"`
	// Status enum: running,waiting,terminated,unknown
	Status       string `json:"status"`
	Namespace    string `json:"namespace"`
	Ready        bool   `json:"ready"`
	RestartCount int32  `json:"restart_count"`
	SessionId    string `json:"session_id"`
	ConnectUrl   string `json:"connect_url"`
	// Type enum: containerlab,default
	Type    string `json:"type"`
	Startup string `json:"startup"`
}

type ContainersGetOutputDTO struct {
	Containers []ContainersGetItem `json:"containers"`
}

type ContainersGetResponse = response.Response[ContainersGetOutputDTO]

func (u *ContainersGetUC) SetContext(user *cms_client.SSOTokenPublicData) *ContainersGetUC {
	u.user = user
	return u
}

func (u *ContainersGetUC) Execute(dto ContainersGetInputDTO) (ContainersGetOutputDTO, error) {
	output := ContainersGetOutputDTO{}
	if u.user == nil {
		return output, errors.New("not logged in")
	}
	username := u.KubernetesAdminQuery.NormalizeEntityName(u.user.Username)
	namespace := u.KubernetesAdminQuery.NormalizeEntityName(dto.Namespace)
	if !cms_client.SSOHasIntersection(
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
		u.user.Roles,
	) && !strings.HasPrefix(namespace, username+"-") {
		return output, utils.FiberValidationException{
			Status:    fiber.StatusForbidden,
			Exception: errors.New("User not allow current namespace"),
		}
	}
	ctx := context.Background()
	yamlContent, _ := u.KubernetesAdminQuery.GetTopologyYAML(ctx, namespace)
	topologyContent, _ := topology.Convert(yamlContent)
	containers, err := u.KubernetesAdminQuery.GetPodByDeploymentName(ctx, namespace, dto.Deployment)
	if err != nil {
		return output, errors.New("not logged in")
	}
	for _, container := range containers {
		tokenWS, _ := u.KubeDashboardClientQuery.Shell(
			u.KubernetesAdminQuery.KubernetesAdminConst.AdminBearerToken,
			container.Namespace,
			container.Pod,
			container.Name,
		)
		var nodeLabel string
		typeNode := "default"
		startup := ""
		var node *topology.TopologiesNode
		if topologyContent != nil {
			node = topologyContent.GetNodeByID(dto.Deployment)
		}
		if node != nil {
			nodeLabel = node.Label
			typeNode = "containerlab"
			startup = node.Data.Startup
		}
		output.Containers = append(
			output.Containers,
			ContainersGetItem{
				Name:         container.Name,
				Label:        nodeLabel,
				Pod:          container.Pod,
				Status:       container.Status,
				Namespace:    container.Namespace,
				Ready:        container.Ready,
				RestartCount: container.RestartCount,
				SessionId:    tokenWS.ID,
				ConnectUrl:   u.KubeDashboardClientQuery.GetUrlByShell(),
				Type:         typeNode,
				Startup:      startup,
			},
		)
	}
	return output, nil
}

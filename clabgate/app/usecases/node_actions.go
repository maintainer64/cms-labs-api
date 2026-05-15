package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

type NodeActionsUC struct {
	*zerolog.Logger
	KubernetesAdminQuery *queries.KubernetesAdminQuery
	user                 *cms_client.SSOTokenPublicData
}

type NodeActionItem struct {
	Nodes string `json:"node"`
	// enum: wipe,restart
	Action string `json:"action"`
}

type NodeActionInputDTO struct {
	Actions       []NodeActionItem `json:"actions"`
	Username      string           `json:"username"`
	AttemptNumber string           `json:"attempt_number"`
}

type NodeActionRequest struct {
	JSONRPC string             `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string             `json:"method" default:"nodes.action" required:"true"`
	Params  NodeActionInputDTO `json:"params,omitempty"`
	ID      string             `json:"id,omitempty" default:"1" required:"true"`
}

type NodeActionOutputDTO struct {
	Count int `json:"count"`
}

type NodeActionResponse struct {
	JSONRPC string              `json:"jsonrpc" default:"2.0" required:"true"`
	Result  NodeActionOutputDTO `json:"result,omitempty"`
	Error   interface{}         `json:"error,omitempty"`
	ID      string              `json:"id,omitempty" default:"1" required:"true"`
}

func (u *NodeActionsUC) SetContext(user *cms_client.SSOTokenPublicData) *NodeActionsUC {
	u.user = user
	return u
}

func (u *NodeActionsUC) Execute(dto NodeActionInputDTO) (NodeActionOutputDTO, error) {
	if u.user == nil {
		return NodeActionOutputDTO{}, errors.New("not logged in")
	}
	if !cms_client.SSOHasIntersection(
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
		u.user.Roles,
	) && dto.Username != u.user.Username {
		return NodeActionOutputDTO{}, errors.New("user not allow connect topology")
	}
	ctx := context.Background()
	u.Logger.Info().Msg(fmt.Sprintf("NodeActionsUC: action by attempt: %v", dto.AttemptNumber))
	namespace := fmt.Sprintf("jup-%s-%s", dto.Username, dto.AttemptNumber)

	var count int
	for _, action := range dto.Actions {
		switch action.Action {
		case "restart":
			if err := u.KubernetesAdminQuery.RestartPod(ctx, namespace, action.Nodes); err != nil {
				u.Logger.Error().Err(err).Msg(
					fmt.Sprintf("failed to restart pod %s with namespace %s", action.Nodes, namespace),
				)
			} else {
				count++
			}
		case "wipe":
			if err := u.KubernetesAdminQuery.DeletePod(ctx, namespace, action.Nodes); err != nil {
				u.Logger.Error().Err(err).Msg(
					fmt.Sprintf("failed to wipe pod %s with namespace %s", action.Nodes, namespace),
				)
			} else {
				count++
			}
		default:
			u.Logger.Warn().Msg(fmt.Sprintf("unknown action: %s", action.Action))
		}
	}

	return NodeActionOutputDTO{
		Count: count,
	}, nil
}

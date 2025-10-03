package usecases

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/jsonrpc"
)

type TopologyCreateUC struct {
	*zerolog.Logger
	KubernetesAdminQuery *queries.KubernetesAdminQuery
	GitClient            queries.GitCodeRegistry
	user                 *cms_client.SSOTokenPublicData
}

type TopologyCreateInputDTO struct {
	Username string `json:"username"`
	TaskID   string `json:"task_id"`
	ReDeploy bool   `json:"redeploy"`
}

type TopologyCreateRequest struct {
	JSONRPC string                 `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string                 `json:"method" default:"topology.create" required:"true"`
	Params  TopologyCreateInputDTO `json:"params,omitempty"`
	ID      string                 `json:"id,omitempty" default:"1" required:"true"`
}
type TopologyCreateOutputDTO struct {
	Namespace        string `json:"namespace"`
	NamespaceCreated bool   `json:"namespace_created"`
	UserCreated      bool   `json:"user_created"`
	DeployCreated    bool   `json:"deploy_created"`
}

type TopologyCreateResponse struct {
	JSONRPC string                  `json:"jsonrpc" default:"2.0" required:"true"`
	Result  TopologyCreateOutputDTO `json:"result,omitempty"`
	Error   interface{}             `json:"error,omitempty"`
	ID      string                  `json:"id,omitempty" default:"1" required:"true"`
}

func (u *TopologyCreateUC) SetContext(user *cms_client.SSOTokenPublicData) *TopologyCreateUC {
	u.user = user
	return u
}

func (u *TopologyCreateUC) Execute(dto TopologyCreateInputDTO) (TopologyCreateOutputDTO, error) {
	output := TopologyCreateOutputDTO{}
	if u.user == nil {
		return output, errors.New("not logged in")
	}
	// Это пользователь, который просит доступ до неймспейса (токен)
	usernameConnected := u.KubernetesAdminQuery.NormalizeEntityName(u.user.Username)
	// Это пользователь, который владелец неймспейса (username)
	usernameOwner := u.KubernetesAdminQuery.NormalizeEntityName(dto.Username)
	if !cms_client.SSOHasIntersection(
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
		u.user.Roles,
	) && usernameOwner != usernameConnected {
		return output, jsonrpc.NewRpcError("user_not_allow_topology", "user not allow connect topology")
	}
	u.Logger.Info().Msg(fmt.Sprintf("TopologyCreateUC: FindTaskById id: %s", dto.TaskID))
	task, err := u.FindTaskById(dto.TaskID)
	if err != nil {
		return output, err
	}
	namespace := u.KubernetesAdminQuery.NormalizeEntityName(
		fmt.Sprintf("%s-%s", usernameOwner, task.NamespaceSuffix),
	)
	u.Logger.Info().Msg(fmt.Sprintf("TopologyCreateUC: Namespace is: %s", namespace))
	output.Namespace = namespace
	ctx := context.Background()
	if usernameConnected != usernameOwner {
		_, err = u.KubernetesAdminQuery.GetNamespaceByName(ctx, namespace)
		if err != nil {
			return output, err
		}
	}
	_, namespaceCreated, err := u.KubernetesAdminQuery.CreateNamespace(ctx, usernameOwner, namespace)
	output.NamespaceCreated = namespaceCreated
	if err != nil {
		return output, err
	}
	users := slices.Compact([]string{usernameOwner, usernameConnected})
	u.Logger.Info().Msg(fmt.Sprintf("TopologyCreateUC: GrantAccessNamespacesList users: %d", len(users)))
	err = u.KubernetesAdminQuery.GrantAccessNamespacesList(ctx, users)
	if err != nil {
		return output, err
	}
	u.Logger.Info().Msg(fmt.Sprintf("TopologyCreateUC: GrantAccessUserNamespace users: %d", len(users)))
	err = u.KubernetesAdminQuery.GrantAccessUserNamespace(ctx, namespace, users)
	if err != nil {
		return output, err
	}
	if !(dto.ReDeploy || namespaceCreated) {
		u.Logger.Info().Msg(fmt.Sprintf("TopologyCreateUC: namespace not deployed: %s", namespace))
		return output, err
	}
	for _, user := range users {
		namespaces, _ := u.KubernetesAdminQuery.GetUserNamespacesOwner(ctx, user)
		for _, namespace := range namespaces {
			u.Logger.Info().Msg(fmt.Sprintf("TopologyCreateUC: grant %s to namespace %s", user, namespace))
			_ = u.KubernetesAdminQuery.GrantAccessUserNamespace(ctx, namespace, []string{user})
		}
	}
	u.Logger.Info().Msg(fmt.Sprintf("TopologyCreateUC: namespace deployed: %s", namespace))
	output.DeployCreated = true
	webUrl, err := u.GitClient.DeployTopology(task.FullPath, namespace)
	if err != nil {
		u.Logger.Error().Msg(fmt.Sprintf("TopologyCreateUC: DeployTopology failed: %s", err))
	}
	u.Logger.Info().Msg(fmt.Sprintf("TopologyCreateUC: Create deploy url: %s", webUrl))
	_ = u.KubernetesAdminQuery.SetSecretByName(ctx, namespace, queries.GitlabWebUrlDeploy, webUrl)
	return output, nil
}

func (u *TopologyCreateUC) FindTaskById(taskID string) (*queries.TaskCodeRegistryItem, error) {
	tasks, err := u.GitClient.TasksList()
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		if task.Id == taskID {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

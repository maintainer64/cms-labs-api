package usecases

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/response"
	"gitlab.com/a10869/api-modules/shared/cms_client"
	"gitlab.com/a10869/api-modules/shared/utils"
	"slices"
)

type TopologiesCreateUC struct {
	*zerolog.Logger
	KubernetesAdminQuery *queries.KubernetesAdminQuery
	GitClient            queries.GitCodeRegistry
	user                 *cms_client.SSOTokenPublicData
}

type TopologiesCreateInputDTO struct {
	UserEmail string `json:"user_email"`
	TaskID    string `json:"task_id"`
	ReDeploy  bool   `json:"redeploy"`
}

type TopologiesCreateOutputDTO struct {
	Namespace        string `json:"namespace"`
	NamespaceCreated bool   `json:"namespace_created"`
	UserCreated      bool   `json:"user_created"`
	DeployCreated    bool   `json:"deploy_created"`
}

type TopologiesCreateResponse = response.Response[TopologiesCreateOutputDTO]

func (u *TopologiesCreateUC) SetContext(user *cms_client.SSOTokenPublicData) *TopologiesCreateUC {
	u.user = user
	return u
}

func (u *TopologiesCreateUC) Execute(dto TopologiesCreateInputDTO) (TopologiesCreateOutputDTO, error) {
	output := TopologiesCreateOutputDTO{}
	if u.user == nil {
		return output, errors.New("not logged in")
	}
	// Это пользователь, который просит доступ до неймспейса (токен)
	usernameConnected := u.KubernetesAdminQuery.NormalizeEntityName(u.user.Username())
	// Это пользователь, который владелец неймспейса (email)
	usernameOwner := u.KubernetesAdminQuery.NormalizeEntityName(cms_client.UsernameByEmail(dto.UserEmail))
	if !cms_client.SSOHasIntersection(
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
		u.user.Roles,
	) && usernameOwner != usernameConnected {
		return output, utils.FiberValidationException{
			Status:    fiber.StatusForbidden,
			Exception: errors.New("User not allow connect topology"),
		}
	}
	u.Logger.Info().Msg(fmt.Sprintf("TopologiesCreateUC: FindTaskById id: %s", dto.TaskID))
	task, err := u.FindTaskById(dto.TaskID)
	if err != nil {
		return output, err
	}
	namespace := u.KubernetesAdminQuery.NormalizeEntityName(
		fmt.Sprintf("%s-%s", usernameOwner, task.NamespaceSuffix),
	)
	u.Logger.Info().Msg(fmt.Sprintf("TopologiesCreateUC: Namespace is: %s", namespace))
	output.Namespace = namespace
	ctx := context.Background()
	_, usernameOwnerCreate, err := u.KubernetesAdminQuery.CreateUser(ctx, usernameOwner)
	output.UserCreated = usernameOwnerCreate
	if err != nil {
		return output, err
	}
	_, _, err = u.KubernetesAdminQuery.CreateUser(ctx, usernameOwner)
	if err != nil {
		return output, err
	}
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
	u.Logger.Info().Msg(fmt.Sprintf("TopologiesCreateUC: GrantAccessNamespacesList users: %d", len(users)))
	err = u.KubernetesAdminQuery.GrantAccessNamespacesList(ctx, users)
	if err != nil {
		return output, err
	}
	u.Logger.Info().Msg(fmt.Sprintf("TopologiesCreateUC: GrantAccessUserNamespace users: %d", len(users)))
	err = u.KubernetesAdminQuery.GrantAccessUserNamespace(ctx, namespace, users)
	if err != nil {
		return output, err
	}
	if !(dto.ReDeploy || namespaceCreated) {
		u.Logger.Info().Msg(fmt.Sprintf("TopologiesCreateUC: namespace not deployed: %s", namespace))
		return output, err
	}
	for _, user := range users {
		namespaces, _ := u.KubernetesAdminQuery.GetUserNamespacesOwner(ctx, user)
		for _, namespace := range namespaces {
			u.Logger.Info().Msg(fmt.Sprintf("TopologiesCreateUC: grant %s to namespace %s", user, namespace))
			_ = u.KubernetesAdminQuery.GrantAccessUserNamespace(ctx, namespace, []string{user})
		}
	}
	u.Logger.Info().Msg(fmt.Sprintf("TopologiesCreateUC: namespace deployed: %s", namespace))
	output.DeployCreated = true
	return output, err
}

func (u *TopologiesCreateUC) FindTaskById(taskID string) (*queries.TaskCodeRegistryItem, error) {
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

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
)

type TopologiesDeleteUC struct {
	*zerolog.Logger
	KubernetesAdminQuery *queries.KubernetesAdminQuery
	user                 *cms_client.SSOTokenPublicData
}

type TopologiesDeleteInputDTO struct {
	Namespaces []string `json:"namespaces"`
}

type TopologiesDeleteOutputDTO struct {
	Namespaces []string `json:"namespaces"`
}

type TopologiesDeleteResponse = response.Response[TopologiesDeleteOutputDTO]

func (u *TopologiesDeleteUC) SetContext(user *cms_client.SSOTokenPublicData) *TopologiesDeleteUC {
	u.user = user
	return u
}

func (u *TopologiesDeleteUC) Execute(dto TopologiesDeleteInputDTO) (TopologiesDeleteOutputDTO, error) {
	output := TopologiesDeleteOutputDTO{}
	if u.user == nil {
		return output, errors.New("not logged in")
	}
	username := u.KubernetesAdminQuery.NormalizeEntityName(u.user.Username())
	if !cms_client.SSOHasIntersection(
		[]string{cms_client.SSOUsersRoleAdmin, cms_client.SSOUsersRoleInstructor},
		u.user.Roles,
	) {
		return output, utils.FiberValidationException{
			Status:    fiber.StatusForbidden,
			Exception: errors.New("User not allow connect topology"),
		}
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

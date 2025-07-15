package usecases

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/response"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

type TokenAccessGetUC struct {
	*zerolog.Logger
	KubernetesAdminQuery *queries.KubernetesAdminQuery
	user                 *cms_client.SSOTokenPublicData
}

type TokenAccessGetInputDTO struct {
}

type TokenAccessGetOutputDTO struct {
	Token string `json:"token"`
}

type TokenAccessGetResponse = response.Response[TokenAccessGetOutputDTO]

func (u *TokenAccessGetUC) SetContext(user *cms_client.SSOTokenPublicData) *TokenAccessGetUC {
	u.user = user
	return u
}

func (u *TokenAccessGetUC) Execute(dto TokenAccessGetInputDTO) (TokenAccessGetOutputDTO, error) {
	output := TokenAccessGetOutputDTO{}
	if u.user == nil {
		return output, errors.New("not logged in")
	}
	username := u.KubernetesAdminQuery.NormalizeEntityName(u.user.Username)
	ctx := context.Background()
	token, err := u.KubernetesAdminQuery.GetUserToken(ctx, username)
	if err != nil {
		return output, err
	}
	output.Token = token
	return output, nil
}

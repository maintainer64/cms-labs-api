package usecases

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/response"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

type TokenFileYAMLGetUC struct {
	*zerolog.Logger
	KubernetesAdminQuery *queries.KubernetesAdminQuery
	user                 *cms_client.SSOTokenPublicData
}

type TokenFileYAMLGetInputDTO struct {
}

type TokenFileYAMLGetOutputDTO struct {
	Content  string `json:"content"`
	Filename string `json:"filename"`
}

type TokenFileYAMLGetResponse = response.Response[TokenFileYAMLGetOutputDTO]

func (u *TokenFileYAMLGetUC) SetContext(user *cms_client.SSOTokenPublicData) *TokenFileYAMLGetUC {
	u.user = user
	return u
}

func (u *TokenFileYAMLGetUC) Execute(dto TokenFileYAMLGetInputDTO) (TokenFileYAMLGetOutputDTO, error) {
	output := TokenFileYAMLGetOutputDTO{}
	if u.user == nil {
		return output, errors.New("not logged in")
	}
	username := u.KubernetesAdminQuery.NormalizeEntityName(u.user.Username)
	ctx := context.Background()
	token, err := u.KubernetesAdminQuery.GetUserToken(ctx, username)
	if err != nil {
		return output, err
	}
	output.Content = token
	return output, nil
}

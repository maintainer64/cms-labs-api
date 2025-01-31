package usecases

import (
	"strings"

	"gitlab.com/a10869/api-modules/backend/app/usecases/response"

	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries/lti_query"
)

type LTIFormGetUC struct {
	LTIFormQueries *lti_query.LTIFormQueries
}

type LTIFormGetInputDTO struct {
	ID uint `json:"id" required:"true"`
}

type LTIFormGetOutputDTO struct {
	Model models.LTIForm `json:"model" required:"true"`
}

type LTIFormGetResponse = response.Response[LTIFormGetOutputDTO]

func (u *LTIFormGetUC) ReplacePublicKey(publicKey string) string {
	return strings.Replace(
		strings.Replace(
			publicKey,
			"BEGIN RSA PUBLIC KEY",
			"BEGIN PUBLIC KEY",
			1,
		),
		"END RSA PUBLIC KEY",
		"END PUBLIC KEY",
		1,
	)
}
func (u *LTIFormGetUC) Execute(dto LTIFormGetInputDTO) (LTIFormGetOutputDTO, error) {
	form, err := u.LTIFormQueries.Get(dto.ID)
	form.PublicKey = u.ReplacePublicKey(form.PublicKey)
	return LTIFormGetOutputDTO{
		Model: form,
	}, err
}

package queries

import (
	"crypto/tls"
	"errors"
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/shared/connection"
)

type KubeDashboardClientQuery struct {
	Config *connection.KubeDashboardConfig
	Client *resty.Client
	Logger *zerolog.Logger
}

type ShellID struct {
	ID string `json:"id"`
}

func (g *KubeDashboardClientQuery) Shell(
	token,
	namespace,
	pod,
	container string,
) (*ShellID, error) {
	if g.Config.NoVerifySSL {
		g.Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) // #nosec G402
	}
	uri := fmt.Sprintf(
		"%s/api/v1/pod/%s/%s/shell/%s",
		g.Config.BaseUrl, namespace, pod, container,
	)
	g.Logger.Info().Msg(fmt.Sprintf("Shell kubedashboard url: %s", uri))
	var responseJson ShellID
	response, _ := g.Client.R().SetHeaders(map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", token),
	}).SetResult(&responseJson).Get(uri)
	if response == nil {
		g.Logger.Info().Msg("Shell: kubedashboard not response")
		return &responseJson, errors.New("kubedashboard server not create response")
	}
	g.Logger.Debug().Msg(fmt.Sprintf("Shell: kubedashboard body: %s", string(response.Body())))
	if response.IsError() {
		g.Logger.Info().Msg(fmt.Sprintf("Shell: kubedashboard statusCode: %d", response.StatusCode()))
		return &responseJson, fmt.Errorf("kubedashboard server error statusCode: %d", response.StatusCode())
	}
	g.Logger.Info().Msg(fmt.Sprintf("Shell: kubedashboard id is: %s", responseJson.ID))
	return &responseJson, nil
}

func (g *KubeDashboardClientQuery) GetUrlByShell() string {
	return fmt.Sprintf("%s/api/sockjs", g.Config.BaseUrl)
}

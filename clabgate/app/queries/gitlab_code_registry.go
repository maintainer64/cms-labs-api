package queries

import (
	"encoding/base64"
	"errors"
	"fmt"

	resty "github.com/go-resty/resty/v2"
	json "github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/shared/connection"
)

type GitlabCodeRegistryQuery struct {
	Config *connection.GitConfig
	Client *resty.Client
	Logger *zerolog.Logger
}

type GitlabCodeRegistryFileResult struct {
	FileName        string `json:"file_name"`
	FilePath        string `json:"file_path"`
	Size            int64  `json:"size"`
	Encoding        string `json:"encoding"`
	ContentSha256   string `json:"content_sha256"`
	Ref             string `json:"ref"`
	BlobId          string `json:"blob_id"`
	CommitId        string `json:"commit_id"`
	LastCommitId    string `json:"last_commit_id"`
	ExecuteFileMode bool   `json:"execute_filemode"`
	Content         string `json:"content"`
}

type GitlabPipelineTrigger struct {
	WebUrl string `json:"web_url"`
}

func (g *GitlabCodeRegistryQuery) TasksList() ([]TaskCodeRegistryItem, error) {
	uri := fmt.Sprintf(
		"%s/api/v4/projects/%s/repository/files/%s?ref=%s",
		g.Config.BaseUrl, g.Config.RepoId, GitClabGateTasksConfig, g.Config.Branch,
	)
	g.Logger.Debug().Msg(fmt.Sprintf("TasksList url: %s", uri))
	var responseJson GitlabCodeRegistryFileResult
	response, _ := g.Client.R().SetHeaders(map[string]string{
		"PRIVATE-TOKEN": g.Config.AccessToken,
	}).SetResult(&responseJson).Get(uri)
	if response == nil {
		g.Logger.Info().Msg(fmt.Sprintf("TasksList: url is incorrect: %s", uri))
		return nil, errors.New("gitlab server not create response")
	}
	if response.IsError() {
		g.Logger.Info().Msg(fmt.Sprintf("TasksList: statusCode: %d", response.StatusCode()))
		return nil, fmt.Errorf("gitlab server error statusCode: %d", response.StatusCode())
	}
	if responseJson.Content == "" || responseJson.Encoding != "base64" {
		g.Logger.Info().Msg("TasksList: gitlab file not encoded base64")
		return nil, errors.New("gitlab file not encoded base64")
	}
	content, err := base64.StdEncoding.DecodeString(responseJson.Content)
	if err != nil {
		g.Logger.Info().Msg("TasksList: content is not decoded base64")
		return nil, err
	}
	var result TaskCodeRegistry
	err = json.Unmarshal(content, &result)
	if err != nil {
		g.Logger.Info().Msg(fmt.Sprintf("TasksList: content is not decoded json %s", content))
		return nil, err
	}
	return result.Labs, nil
}

func (g *GitlabCodeRegistryQuery) DeployTopology(pathFile string, namespace string) (string, error) {
	uri := fmt.Sprintf(
		"%s/api/v4/projects/%s/trigger/pipeline",
		g.Config.BaseUrl, g.Config.RepoId,
	)
	g.Logger.Debug().Msg(fmt.Sprintf("DeployTopology url: %s", uri))
	var responseJson GitlabPipelineTrigger
	response, _ := g.Client.R().SetFormData(map[string]string{
		"token": g.Config.TriggerToken,
		"ref":   g.Config.Branch,
		"variables[CLABERNETES_TOPOLOGY_NAMESPACE]": namespace,
		"variables[CLABERNETES_TOPOLOGY_PATH]":      pathFile,
	}).SetResult(&responseJson).Post(uri)
	if response == nil {
		g.Logger.Info().Msg(fmt.Sprintf("DeployTopology: url is incorrect: %s", uri))
		return "", errors.New("gitlab server not create response")
	}
	if response.IsError() {
		g.Logger.Info().Msg(fmt.Sprintf("TasksList: statusCode: %d", response.StatusCode()))
		return "", fmt.Errorf("gitlab server error statusCode: %d", response.StatusCode())
	}
	return responseJson.WebUrl, nil
}

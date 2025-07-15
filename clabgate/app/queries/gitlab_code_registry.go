package queries

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/shared/connection"
)

type GitlabCodeRegistryQuery struct {
	Config *connection.GitConfig
	Client *resty.Client
	*zerolog.Logger
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

func (g *GitlabCodeRegistryQuery) TasksList() ([]TaskCodeRegistryItem, error) {
	uri := fmt.Sprintf(
		"%s/api/v4/projects/%s/repository/files/%s?ref=%s",
		g.Config.BaseUrl, g.Config.RepoId, GitClabGateTasksConfig, g.Config.Branch,
	)
	g.Logger.Debug().Msg(fmt.Sprintf("TasksList url: %s", uri))
	var responseJson GitlabCodeRegistryFileResult
	response, err := g.Client.R().SetHeaders(map[string]string{
		"PRIVATE-TOKEN": g.Config.Token,
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
		g.Logger.Info().Msg(fmt.Sprintf("TasksList: content is not decoded base64"))
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

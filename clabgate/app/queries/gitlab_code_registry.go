package queries

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"gitlab.com/a10869/api-modules/shared/connection"
)

type GitlabCodeRegistryQuery struct {
	Config *connection.GitConfig
	Client *resty.Client
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
	var responseJson GitlabCodeRegistryFileResult
	response, err := g.Client.R().SetHeaders(map[string]string{
		"Content-Type":  "application/json",
		"PRIVATE-TOKEN": g.Config.Token,
	}).SetResult(&responseJson).Get(uri)
	if response == nil {
		return nil, errors.New("gitlab server not create response")
	}
	if response.IsError() {
		return nil, fmt.Errorf("gitlab server error statusCode: %d", response.StatusCode())
	}
	if responseJson.Content == "" || responseJson.Encoding != "base64" {
		return nil, errors.New("gitlab file not encoded base64")
	}
	content, err := base64.StdEncoding.DecodeString(responseJson.Content)
	if err != nil {
		return nil, err
	}
	var results []TaskCodeRegistryItem
	err = json.Unmarshal(content, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

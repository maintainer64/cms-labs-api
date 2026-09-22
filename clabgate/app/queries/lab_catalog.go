package queries

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strconv"
	"strings"

	resty "github.com/go-resty/resty/v2"
)

const gitLabPageSize = 100

type repositoryProvider string

const (
	providerGitLab repositoryProvider = "gitlab"
	providerGitHub repositoryProvider = "github"
)

type repositoryLocation struct {
	provider repositoryProvider
	project  string
	apiBase  string
}

type LabCatalog struct {
	projectURL string
	branch     string
	token      string
	client     *resty.Client
}

type LabBundle struct {
	Manifest   string
	Revision   string
	Files      []string
	ProjectURL string
}

type gitLabTreeEntry struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

type gitLabCommit struct {
	ID string `json:"id"`
}

type gitHubCommit struct {
	SHA string `json:"sha"`
}

type gitHubTree struct {
	Tree      []gitLabTreeEntry `json:"tree"`
	Truncated bool              `json:"truncated"`
}

func NewLabCatalog(projectURL, branch, token string, client *resty.Client) *LabCatalog {
	return &LabCatalog{
		projectURL: strings.TrimRight(projectURL, "/"),
		branch:     branch,
		token:      token,
		client:     client,
	}
}

// Bundle resolves one directory in a GitLab or GitHub project and downloads every YAML
// file below it in deterministic order. Kubernetes object validation happens
// later, before any document is applied.
func (c *LabCatalog) Bundle(ctx context.Context, labsPath string) (LabBundle, error) {
	return c.BundleAt(ctx, labsPath, "")
}

// BundleAt uses an already resolved commit when retrying a partially created
// session, so a moving branch cannot alter that session between retries.
func (c *LabCatalog) BundleAt(ctx context.Context, labsPath, pinnedRevision string) (LabBundle, error) {
	if c.projectURL == "" {
		return LabBundle{}, fmt.Errorf("CMS_TASK_URL is not configured")
	}
	repository, err := parseRepositoryURL(c.projectURL)
	if err != nil {
		return LabBundle{}, err
	}
	cleaned, err := cleanRepositoryPath(labsPath)
	if err != nil {
		return LabBundle{}, err
	}

	revision := strings.TrimSpace(pinnedRevision)
	if revision == "" {
		revision, err = c.resolveRevision(ctx, repository)
		if err != nil {
			return LabBundle{}, err
		}
	}
	files, err := c.listYAMLFiles(ctx, repository, cleaned, revision)
	if err != nil {
		return LabBundle{}, err
	}

	documents := make([]string, 0, len(files))
	for _, file := range files {
		response, requestErr := c.downloadFile(ctx, repository, file, revision)
		if requestErr != nil {
			return LabBundle{}, fmt.Errorf("download task manifest %q: %w", file, requestErr)
		}
		if response.StatusCode() < http.StatusOK || response.StatusCode() >= http.StatusMultipleChoices {
			return LabBundle{}, fmt.Errorf("download task manifest %q: HTTP %d", file, response.StatusCode())
		}
		if value := strings.TrimSpace(response.String()); value != "" {
			documents = append(documents, value)
		}
	}

	return LabBundle{
		Manifest:   strings.Join(documents, "\n---\n"),
		Revision:   revision,
		Files:      files,
		ProjectURL: c.projectURL,
	}, nil
}

func (c *LabCatalog) resolveRevision(ctx context.Context, repository repositoryLocation) (string, error) {
	branch := strings.TrimSpace(c.branch)
	if branch == "" {
		branch = "master"
	}
	if repository.provider == providerGitHub {
		commit := gitHubCommit{}
		response, err := c.request(ctx, repository.provider).SetResult(&commit).Get(
			repository.apiBase + "/repos/" + repository.project + "/commits/" + url.PathEscape(branch),
		)
		if err != nil {
			return "", fmt.Errorf("resolve GitHub ref %q: %w", branch, err)
		}
		if response.StatusCode() < http.StatusOK || response.StatusCode() >= http.StatusMultipleChoices {
			return "", fmt.Errorf("resolve GitHub ref %q: HTTP %d", branch, response.StatusCode())
		}
		if commit.SHA == "" {
			return "", fmt.Errorf("resolve GitHub ref %q: empty commit SHA", branch)
		}
		return commit.SHA, nil
	}

	commit := gitLabCommit{}
	response, err := c.request(ctx, repository.provider).SetResult(&commit).Get(
		repository.apiBase + "/projects/" + url.PathEscape(repository.project) + "/repository/commits/" + url.PathEscape(branch),
	)
	if err != nil {
		return "", fmt.Errorf("resolve GitLab ref %q: %w", branch, err)
	}
	if response.StatusCode() < http.StatusOK || response.StatusCode() >= http.StatusMultipleChoices {
		return "", fmt.Errorf("resolve GitLab ref %q: HTTP %d", branch, response.StatusCode())
	}
	if commit.ID == "" {
		return "", fmt.Errorf("resolve GitLab ref %q: empty commit ID", branch)
	}
	return commit.ID, nil
}

func (c *LabCatalog) listYAMLFiles(ctx context.Context, repository repositoryLocation, labsPath, revision string) ([]string, error) {
	if repository.provider == providerGitHub {
		tree := gitHubTree{}
		response, err := c.request(ctx, repository.provider).
			SetResult(&tree).
			SetQueryParam("recursive", "1").
			Get(repository.apiBase + "/repos/" + repository.project + "/git/trees/" + url.PathEscape(revision))
		if err != nil {
			return nil, fmt.Errorf("list task manifests in %q: %w", labsPath, err)
		}
		if response.StatusCode() == http.StatusNotFound {
			return nil, fmt.Errorf("lab directory %q was not found in GitHub", labsPath)
		}
		if response.StatusCode() < http.StatusOK || response.StatusCode() >= http.StatusMultipleChoices {
			return nil, fmt.Errorf("list task manifests in %q: HTTP %d", labsPath, response.StatusCode())
		}
		if tree.Truncated {
			return nil, fmt.Errorf("GitHub repository tree is truncated; split the task catalog or use GitLab")
		}
		files := filterYAMLFiles(tree.Tree, labsPath)
		if len(files) == 0 {
			return nil, fmt.Errorf("lab directory %q contains no YAML manifests", labsPath)
		}
		return files, nil
	}

	projectAPI := repository.apiBase + "/projects/" + url.PathEscape(repository.project)
	files := make([]string, 0)
	for pageNumber := 1; ; pageNumber++ {
		entries := []gitLabTreeEntry{}
		response, err := c.request(ctx, repository.provider).
			SetResult(&entries).
			SetQueryParams(map[string]string{
				"path": labsPath, "ref": revision, "recursive": "true",
				"per_page": strconv.Itoa(gitLabPageSize), "page": strconv.Itoa(pageNumber),
			}).
			Get(projectAPI + "/repository/tree")
		if err != nil {
			return nil, fmt.Errorf("list task manifests in %q: %w", labsPath, err)
		}
		if response.StatusCode() == http.StatusNotFound {
			return nil, fmt.Errorf("lab directory %q was not found in GitLab", labsPath)
		}
		if response.StatusCode() < http.StatusOK || response.StatusCode() >= http.StatusMultipleChoices {
			return nil, fmt.Errorf("list task manifests in %q: HTTP %d", labsPath, response.StatusCode())
		}
		for _, entry := range entries {
			extension := strings.ToLower(path.Ext(entry.Path))
			if entry.Type == "blob" && (extension == ".yaml" || extension == ".yml") {
				files = append(files, entry.Path)
			}
		}
		if response.Header().Get("X-Next-Page") == "" {
			break
		}
	}
	sort.Strings(files)
	return files, nil
}

func filterYAMLFiles(entries []gitLabTreeEntry, labsPath string) []string {
	prefix := strings.TrimSuffix(labsPath, "/") + "/"
	files := make([]string, 0)
	for _, entry := range entries {
		extension := strings.ToLower(path.Ext(entry.Path))
		if entry.Type == "blob" && strings.HasPrefix(entry.Path, prefix) && (extension == ".yaml" || extension == ".yml") {
			files = append(files, entry.Path)
		}
	}
	sort.Strings(files)
	return files
}

func (c *LabCatalog) downloadFile(ctx context.Context, repository repositoryLocation, file, revision string) (*resty.Response, error) {
	if repository.provider == providerGitHub {
		return c.request(ctx, repository.provider).
			SetHeader("Accept", "application/vnd.github.raw+json").
			SetQueryParam("ref", revision).
			Get(repository.apiBase + "/repos/" + repository.project + "/contents/" + escapeRepositoryPath(file))
	}
	return c.request(ctx, repository.provider).
		SetQueryParam("ref", revision).
		Get(repository.apiBase + "/projects/" + url.PathEscape(repository.project) + "/repository/files/" + url.PathEscape(file) + "/raw")
}

func (c *LabCatalog) request(ctx context.Context, provider repositoryProvider) *resty.Request {
	request := c.client.R().SetContext(ctx)
	if c.token != "" {
		if provider == providerGitHub {
			request.SetAuthToken(c.token)
		} else {
			request.SetHeader("PRIVATE-TOKEN", c.token)
		}
	}
	if provider == providerGitHub {
		request.SetHeader("X-GitHub-Api-Version", "2022-11-28")
	}
	return request
}

func parseRepositoryURL(raw string) (repositoryLocation, error) {
	parsed, err := url.Parse(raw)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return repositoryLocation{}, fmt.Errorf("CMS_TASK_URL must be a full GitLab or GitHub project URL")
	}
	if parsed.Scheme != "https" && parsed.Scheme != "http" {
		return repositoryLocation{}, fmt.Errorf("CMS_TASK_URL must use HTTP or HTTPS")
	}
	if parsed.RawQuery != "" || parsed.Fragment != "" || parsed.User != nil {
		return repositoryLocation{}, fmt.Errorf("CMS_TASK_URL must not contain credentials, query or fragment")
	}
	project := strings.Trim(strings.TrimSuffix(parsed.EscapedPath(), ".git"), "/")
	decoded, decodeErr := url.PathUnescape(project)
	if decodeErr != nil || decoded == "" || !strings.Contains(decoded, "/") {
		return repositoryLocation{}, fmt.Errorf("CMS_TASK_URL must include repository owner and project")
	}
	if strings.EqualFold(parsed.Hostname(), "github.com") {
		if len(strings.Split(decoded, "/")) != 2 {
			return repositoryLocation{}, fmt.Errorf("GitHub CMS_TASK_URL must include owner and repository")
		}
		return repositoryLocation{provider: providerGitHub, project: escapeRepositoryPath(decoded), apiBase: "https://api.github.com"}, nil
	}
	return repositoryLocation{provider: providerGitLab, project: decoded, apiBase: parsed.Scheme + "://" + parsed.Host + "/api/v4"}, nil
}

func parseGitLabProjectURL(raw string) (project, apiBase string, err error) {
	repository, err := parseRepositoryURL(raw)
	if err != nil {
		return "", "", err
	}
	if repository.provider != providerGitLab {
		return "", "", fmt.Errorf("CMS_TASK_URL is not a GitLab project URL")
	}
	return repository.project, repository.apiBase, nil
}

func escapeRepositoryPath(value string) string {
	parts := strings.Split(value, "/")
	for index := range parts {
		parts[index] = url.PathEscape(parts[index])
	}
	return strings.Join(parts, "/")
}

func cleanRepositoryPath(value string) (string, error) {
	if strings.TrimSpace(value) == "" {
		return "", fmt.Errorf("labs_path is required")
	}
	cleaned := strings.Trim(path.Clean("/"+value), "/")
	if cleaned == "" || cleaned == "." || strings.Contains(value, "\\") {
		return "", fmt.Errorf("invalid labs_path")
	}
	return cleaned, nil
}

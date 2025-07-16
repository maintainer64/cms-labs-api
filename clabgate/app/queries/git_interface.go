package queries

type TaskCodeRegistryItem struct {
	Id              string `json:"id"`
	NamespaceSuffix string `json:"namespace_suffix"`
	FullPath        string `json:"full_path"`
	Title           string `json:"title"`
}

type TaskCodeRegistry struct {
	Labs []TaskCodeRegistryItem `json:"labs"`
}

var (
	GitClabGateTasksConfig string = "clabgate.json"
)

type GitCodeRegistry interface {
	TasksList() ([]TaskCodeRegistryItem, error)
	DeployTopology(pathFile string, namespace string) (string, error)
}

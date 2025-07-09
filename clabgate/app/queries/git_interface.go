package queries

type TaskCodeRegistryItem struct {
	Id       string `json:"id"`
	Filename string `json:"filename"`
	Name     string `json:"name"`
}

var (
	GitClabGateTasksConfig string = "clabgate.json"
)

type GitCodeRegistry interface {
	TasksList() ([]TaskCodeRegistryItem, error)
}

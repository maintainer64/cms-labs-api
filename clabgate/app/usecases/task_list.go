package usecases

import (
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
)

type TaskListUC struct {
	*zerolog.Logger
	GitClient queries.GitCodeRegistry
}

type TaskListRequest struct {
	JSONRPC string           `json:"jsonrpc" default:"2.0" required:"true"`
	Method  string           `json:"method" default:"tasks.list" required:"true"`
	Params  TaskListInputDTO `json:"params,omitempty"`
	ID      string           `json:"id,omitempty" default:"1" required:"true"`
}

type TaskListInputDTO struct {
}

type TaskListOutputDTO struct {
	Model []queries.TaskCodeRegistryItem `json:"model" validate:"required"`
}

type TaskListResponse struct {
	JSONRPC string            `json:"jsonrpc" default:"2.0" required:"true"`
	Result  TaskListOutputDTO `json:"result,omitempty"`
	Error   interface{}       `json:"error,omitempty"`
	ID      string            `json:"id,omitempty" default:"1" required:"true"`
}

func (u *TaskListUC) Execute(dto TaskListInputDTO) (TaskListOutputDTO, error) {
	entities, err := u.GitClient.TasksList()
	return TaskListOutputDTO{
		Model: entities,
	}, err
}

package usecases

import (
	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/clabgate/app/queries"
	"gitlab.com/a10869/api-modules/clabgate/app/usecases/response"
)

type TasksListUC struct {
	*zerolog.Logger
	GitClient queries.GitCodeRegistry
}

type TasksListInputDTO struct {
}

type TasksListOutputDTO struct {
	Model []queries.TaskCodeRegistryItem `json:"model" validate:"required"`
}

type TasksListResponse = response.Response[TasksListOutputDTO]

func (u *TasksListUC) Execute(dto TasksListInputDTO) (TasksListOutputDTO, error) {
	entities, err := u.GitClient.TasksList()
	return TasksListOutputDTO{
		Model: entities,
	}, err
}

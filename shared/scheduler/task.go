package scheduler

// Task интерфейс для задач, которые будут выполняться
type Task interface {
	Execute() error
}

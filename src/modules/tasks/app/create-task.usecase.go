package usecases

import (
	"app/src/modules/tasks/domain/entities"
	"app/src/modules/tasks/domain/repositories"
	error "app/src/shared/errors"
)

type CreateTaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *CreateTaskUseCase) Execute(createProduct entities.CreateTask, userId uint) (entities.Task, error.Error) {
	tasks, err := l.TaskRepository.Create(createProduct, userId)
	if err.StatusCode != 0 {
		return entities.Task{}, err
	}
	return tasks, error.Error{}
}

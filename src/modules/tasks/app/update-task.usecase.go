package usecases

import (
	"app/src/modules/tasks/domain/entities"
	"app/src/modules/tasks/domain/repositories"
	error "app/src/shared/errors"
)

type UpdateTaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *UpdateTaskUseCase) Execute(updateProduct entities.UpdateTask, id int) (entities.Task, error.Error) {
	tasks, err := l.TaskRepository.Update(updateProduct, id)
	if err.StatusCode != 0 {
		return entities.Task{}, err
	}
	return tasks, error.Error{}
}

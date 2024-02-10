package usecases

import (
	"app/src/modules/tasks/domain/entities"
	"app/src/modules/tasks/domain/repositories"
	error "app/src/shared/errors"
)

type ListTasksUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *ListTasksUseCase) Execute(userId uint) (entities.ListTask, error.Error) {
	tasks, err := l.TaskRepository.FindAll(userId)
	if err.StatusCode != 0 {
		return entities.ListTask{}, err
	}
	return tasks, error.Error{}
}

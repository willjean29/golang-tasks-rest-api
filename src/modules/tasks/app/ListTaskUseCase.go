package usecases

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/domain/repositories"
	error "app/src/shared/errors"
)

type ListTasksUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *ListTasksUseCase) Execute() (models.IListTask, error.Error) {
	tasks, err := l.TaskRepository.FindAll()
	if err.StatusCode != 0 {
		return models.IListTask{}, err
	}
	return tasks, error.Error{}
}

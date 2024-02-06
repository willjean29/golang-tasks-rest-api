package usecases

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/domain/repositories"
	error "app/src/shared/errors"
)

type GetTaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *GetTaskUseCase) Execute(id int) (models.ITask, error.Error) {
	tasks, err := l.TaskRepository.FindById(id)
	if err.StatusCode != 0 {
		return models.ITask{}, err
	}
	return tasks, error.Error{}
}

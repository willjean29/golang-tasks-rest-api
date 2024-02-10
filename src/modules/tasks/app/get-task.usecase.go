package usecases

import (
	"app/src/modules/tasks/domain/entities"
	"app/src/modules/tasks/domain/repositories"
	error "app/src/shared/errors"
)

type GetTaskUseCase struct {
	TaskRepository repositories.TaskRepository
}

func (l *GetTaskUseCase) Execute(id int) (entities.Task, error.Error) {
	tasks, err := l.TaskRepository.FindById(id)
	if err.StatusCode != 0 {
		return entities.Task{}, err
	}
	return tasks, error.Error{}
}

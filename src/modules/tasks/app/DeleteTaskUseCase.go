package usecases

import (
	"app/src/modules/tasks/domain/repositories"
	error "app/src/shared/errors"
)

type DeleteTaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *DeleteTaskUseCase) Execute(id int) error.Error {
	err := l.TaskRepository.Delete(id)
	if err.StatusCode != 0 {
		return err
	}
	return error.Error{}
}

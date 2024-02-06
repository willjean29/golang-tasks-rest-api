package usecases

import (
	"app/src/modules/tasks/domain/repositories"
)

type DeleteTaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *DeleteTaskUseCase) Execute(id int) error {
	err := l.TaskRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

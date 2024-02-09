package usecases

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/domain/repositories"
	error "app/src/shared/errors"
)

type SaveTaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (s *SaveTaskUseCase) Execute(task models.ITask) error.Error {
	err := s.TaskRepository.Save(task)
	if err.StatusCode != 0 {
		return err
	}
	return error.Error{}
}

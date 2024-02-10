package usecases

import (
	"app/src/modules/tasks/domain/entities"
	"app/src/modules/tasks/domain/repositories"
	error "app/src/shared/errors"
)

type SaveTaskUseCase struct {
	TaskRepository repositories.TaskRepository
}

func (s *SaveTaskUseCase) Execute(task entities.Task) error.Error {
	err := s.TaskRepository.Save(task)
	if err.StatusCode != 0 {
		return err
	}
	return error.Error{}
}

package usecases

import (
	"app/src/modules/users/domain/entities"
	"app/src/modules/users/domain/repositories"
	error "app/src/shared/errors"
)

type SaveUserUseCase struct {
	UserRepository repositories.UserRepository
}

func (s *SaveUserUseCase) Execute(user entities.User) error.Error {
	err := s.UserRepository.Save(user)
	if err.StatusCode != 0 {
		return err
	}
	return error.Error{}
}

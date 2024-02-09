package usecases

import (
	"app/src/modules/users/domain/models"
	"app/src/modules/users/domain/repositories"
	error "app/src/shared/errors"
)

type SaveUserUseCase struct {
	UserRepository repositories.IUserRepository
}

func (s *SaveUserUseCase) Execute(user models.IUser) error.Error {
	err := s.UserRepository.Save(user)
	if err.StatusCode != 0 {
		return err
	}
	return error.Error{}
}

package usecases

import (
	"app/src/modules/users/domain/entities"
	"app/src/modules/users/domain/repositories"
	error "app/src/shared/errors"
)

type ListUserUseCase struct {
	UserRepository repositories.UserRepository
}

func (l *ListUserUseCase) Execute() (entities.ListUsers, error.Error) {
	users, err := l.UserRepository.FindAll()
	if err.StatusCode != 0 {
		return entities.ListUsers{}, err
	}
	return users, error.Error{}
}

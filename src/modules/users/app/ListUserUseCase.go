package usecases

import (
	"app/src/modules/users/domain/models"
	"app/src/modules/users/domain/repositories"
	error "app/src/shared/errors"
)

type ListUserUseCase struct {
	UserRepository repositories.IUserRepository
}

func (l *ListUserUseCase) Execute() (models.IListUsers, error.Error) {
	users, err := l.UserRepository.FindAll()
	if err.StatusCode != 0 {
		return models.IListUsers{}, err
	}
	return users, error.Error{}
}

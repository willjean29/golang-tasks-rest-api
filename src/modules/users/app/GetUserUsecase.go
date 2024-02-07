package usecases

import (
	"app/src/modules/users/domain/models"
	"app/src/modules/users/domain/repositories"
	error "app/src/shared/errors"
)

type GetUserUseCase struct {
	UserRepository repositories.IUserRepository
}

func (g *GetUserUseCase) Execute(id int) (models.IUser, error.Error) {
	user, err := g.UserRepository.FindById(id)
	if err.StatusCode != 0 {
		return models.IUser{}, err
	}
	return user, error.Error{}
}

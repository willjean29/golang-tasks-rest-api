package usecases

import (
	"app/src/modules/users/domain/entities"
	"app/src/modules/users/domain/repositories"
	error "app/src/shared/errors"
)

type GetUserUseCase struct {
	UserRepository repositories.UserRepository
}

func (g *GetUserUseCase) Execute(id int) (entities.User, error.Error) {
	user, err := g.UserRepository.FindById(id)
	if err.StatusCode != 0 {
		return entities.User{}, err
	}
	return user, error.Error{}
}

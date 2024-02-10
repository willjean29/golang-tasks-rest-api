package usecases

import (
	"app/src/modules/users/domain/entities"
	"app/src/modules/users/domain/repositories"
	error "app/src/shared/errors"
	"fmt"
)

type CreateUserUseCase struct {
	UserRepository repositories.UserRepository
}

func (c *CreateUserUseCase) Execute(createUser entities.CreateUser) (entities.User, string, error.Error) {
	hashPassword, err := hashAdapter.HashPassword(createUser.Password)
	if err != nil {
		return entities.User{}, "", *error.New("Internal server error", 500, err)
	}
	createUser.Password = hashPassword

	user, errorApp := c.UserRepository.Create(createUser)
	if errorApp.StatusCode != 0 {
		return entities.User{}, "", errorApp
	}

	tokenString, err := tokenAdapter.GenerateToken("userId", fmt.Sprintf("%d", user.ID))
	if err != nil {
		return entities.User{}, "", *error.New("Internal server error", 500, err)
	}
	return user, tokenString, error.Error{}
}

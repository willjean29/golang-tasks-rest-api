package usecases

import (
	"app/src/modules/users/domain/models"
	"app/src/modules/users/domain/repositories"
	error "app/src/shared/errors"
	"fmt"
)

type CreateUserUseCase struct {
	UserRepository repositories.IUserRepository
}

func (c *CreateUserUseCase) Execute(createUser models.ICreateUser) (models.IUser, string, error.Error) {
	user, errorApp := c.UserRepository.Create(createUser)
	if errorApp.StatusCode != 0 {
		return models.IUser{}, "", errorApp
	}

	tokenString, err := tokenProvider.GenerateToken("userId", fmt.Sprintf("%d", user.ID))
	if err != nil {
		return models.IUser{}, "", *error.New("Internal server error", 500, err)
	}
	return user, tokenString, error.Error{}
}

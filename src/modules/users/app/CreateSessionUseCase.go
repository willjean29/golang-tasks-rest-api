package usecases

import (
	hash "app/providers/HashProvider"
	token "app/providers/TokenProvider"
	"app/src/modules/users/domain/models"
	"app/src/modules/users/domain/repositories"
	error "app/src/shared/errors"
)

var hashProvider hash.HashProvider = &hash.BcryptProvider{}
var tokenProvider token.TokenProvider = token.NewJwtProvider()

type CreateSessionUseCase struct {
	UserRepository repositories.IUserRepository
}

func (c *CreateSessionUseCase) Execute(createSession models.ICreateSession) (models.IUser, error.Error) {
	user, errorApp := c.UserRepository.FindByEmail(createSession.Email)
	if errorApp.StatusCode != 0 {
		return models.IUser{}, errorApp
	}
	err := hashProvider.ComparePasswords(user.Password, createSession.Password)
	if err != nil {
		return models.IUser{}, *error.New("Invalid data (password)", 401, err)
	}
	return user, error.Error{}
}

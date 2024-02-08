package usecases

import (
	"app/src/modules/users/domain/models"
	"app/src/modules/users/domain/repositories"
	hash "app/src/shared/adapters/hash"
	token "app/src/shared/adapters/token"
	error "app/src/shared/errors"
	"fmt"
)

var hashAdapter hash.HashAdapter = &hash.BcryptAdapter{}
var tokenAdapter token.TokenAdapter = token.NewJwtAdapter()

type CreateSessionUseCase struct {
	UserRepository repositories.IUserRepository
}

func (c *CreateSessionUseCase) Execute(createSession models.ICreateSession) (models.IUser, string, error.Error) {
	user, errorApp := c.UserRepository.FindByEmail(createSession.Email)
	if errorApp.StatusCode != 0 {
		return models.IUser{}, "", errorApp
	}
	err := hashAdapter.ComparePasswords(user.Password, createSession.Password)
	if err != nil {
		return models.IUser{}, "", *error.New("Invalid data (password)", 401, err)
	}
	tokenString, err := tokenAdapter.GenerateToken("userId", fmt.Sprintf("%d", user.ID))
	if err != nil {
		return models.IUser{}, "", *error.New("Internal server error", 500, err)
	}
	return user, tokenString, error.Error{}
}

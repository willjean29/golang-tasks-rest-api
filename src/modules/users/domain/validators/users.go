package validators

import (
	"app/src/modules/users/domain/models"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type UserValidator struct {
}

func NewUserValidator() UserValidator {
	validate = validator.New(validator.WithRequiredStructEnabled())
	return UserValidator{}
}

func (u *UserValidator) ValidateCreateUser(createUser models.ICreateUser) error {
	var err error
	err = validate.Struct(createUser)
	return err
}

func (u *UserValidator) ValidateLogin(createSession models.ICreateSession) error {
	var err error
	err = validate.Struct(createSession)
	return err
}

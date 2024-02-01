package validators

import (
	"app/dtos"

	"github.com/go-playground/validator/v10"
)

type UserValidator struct{}

func NewUserValidator() UserValidator {
	validate = validator.New(validator.WithRequiredStructEnabled())
	return UserValidator{}
}

func (u *UserValidator) ValidateCreateUser(createUserDto dtos.CreateUserDto) error {
	var err error
	err = validate.Struct(createUserDto)
	return err
}

func (u *UserValidator) ValidateLogin(loginDto dtos.LoginDto) error {
	var err error
	err = validate.Struct(loginDto)
	return err
}

package repositories

import (
	"app/src/modules/users/domain/entities"
	error "app/src/shared/errors"
)

type UserRepository interface {
	FindAll() (entities.ListUsers, error.Error)
	FindById(id int) (entities.User, error.Error)
	FindByEmail(email string) (entities.User, error.Error)
	Create(task entities.CreateUser) (entities.User, error.Error)
	Save(task entities.User) error.Error
}

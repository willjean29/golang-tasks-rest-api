package repositories

import (
	"app/src/modules/users/domain/models"
	error "app/src/shared/errors"
)

type IUserRepository interface {
	FindAll() (models.IListUsers, error.Error)
	FindById(id int) (models.IUser, error.Error)
	FindByEmail(email string) (models.IUser, error.Error)
	Create(task models.ICreateUser) (models.IUser, error.Error)
	Save(task models.IUser) error.Error
	// Update(task models.IUser, id int) (models.IUser, error.Error)
	// Delete(id int) error.Error
}

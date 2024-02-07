package repositories

import (
	"app/src/modules/users/domain/models"
	error "app/src/shared/errors"
)

type IUserRepository interface {
	FindAll() (models.IListUsers, error.Error)
	FindById(id int) (models.IUser, error.Error)
	// GetUserByEmail(email string) (models.IUser, error.Error)
	// Create(task models.IUser) (models.IUser, error.Error)
	// // Save(task models.ITask) error
	// Update(task models.IUser, id int) (models.IUser, error.Error)
	// Delete(id int) error.Error
}

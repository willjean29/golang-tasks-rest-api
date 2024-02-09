package repositories

import (
	"app/src/modules/users/domain/datasource"
	"app/src/modules/users/domain/models"
	error "app/src/shared/errors"
)

type UsersRepository struct {
	Datasource datasource.IUserDatasource
}

func (u *UsersRepository) FindAll() (models.IListUsers, error.Error) {
	return u.Datasource.FindAll()
}

func (u *UsersRepository) FindById(id int) (models.IUser, error.Error) {
	return u.Datasource.FindById(id)
}

func (u *UsersRepository) FindByEmail(email string) (models.IUser, error.Error) {
	return u.Datasource.FindByEmail(email)
}

func (u *UsersRepository) Create(createUser models.ICreateUser) (models.IUser, error.Error) {
	return u.Datasource.Create(createUser)
}

func (u *UsersRepository) Save(user models.IUser) error.Error {
	return u.Datasource.Save(user)
}

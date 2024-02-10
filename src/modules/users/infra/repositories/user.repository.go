package repositories

import (
	"app/src/modules/users/domain/datasource"
	"app/src/modules/users/domain/entities"
	error "app/src/shared/errors"
)

type UserRepository struct {
	Datasource datasource.UserDatasource
}

func (u *UserRepository) FindAll() (entities.ListUsers, error.Error) {
	return u.Datasource.FindAll()
}

func (u *UserRepository) FindById(id int) (entities.User, error.Error) {
	return u.Datasource.FindById(id)
}

func (u *UserRepository) FindByEmail(email string) (entities.User, error.Error) {
	return u.Datasource.FindByEmail(email)
}

func (u *UserRepository) Create(createUser entities.CreateUser) (entities.User, error.Error) {
	return u.Datasource.Create(createUser)
}

func (u *UserRepository) Save(user entities.User) error.Error {
	return u.Datasource.Save(user)
}

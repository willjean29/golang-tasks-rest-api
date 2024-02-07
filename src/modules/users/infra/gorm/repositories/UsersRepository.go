package repositories

import (
	"app/src/modules/users/domain/models"
	"app/src/modules/users/infra/gorm/entities"
	"app/src/modules/users/infra/utils"
	error "app/src/shared/errors"
	"errors"

	"gorm.io/gorm"
)

type UsersRepository struct {
	Repository *gorm.DB
}

func (u *UsersRepository) FindAll() (models.IListUsers, error.Error) {
	var listUsers entities.ListUsers
	var users models.IListUsers

	query := u.Repository.Preload("Tasks").Find(&listUsers)
	err := query.Error

	if err != nil {
		return models.IListUsers{}, *error.New("Data of users not found", 404, err)
	}

	users = utils.MapperToUsers(listUsers)
	return users, error.Error{}
}

func (u *UsersRepository) FindById(id int) (models.IUser, error.Error) {
	var user entities.User
	var userModel models.IUser

	query := u.Repository.Preload("Tasks").First(&user, id)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.IUser{}, *error.New("User not found", 404, errors.New(err.Error()))
		} else {
			return models.IUser{}, *error.New("Error get user", 500, errors.New(err.Error()))
		}
	}
	if query.RowsAffected == 0 {
		return models.IUser{}, *error.New("User not found", 404, errors.New("User not found"))
	}

	userModel = utils.MapperToUser(user)
	return userModel, error.Error{}
}

func (u *UsersRepository) FindByEmail(email string) (models.IUser, error.Error) {
	var user entities.User
	var userModel models.IUser

	query := u.Repository.Where("email = ?", email).First(&user)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.IUser{}, *error.New("User not found", 404, errors.New(err.Error()))
		} else {
			return models.IUser{}, *error.New("Error get user", 500, errors.New(err.Error()))
		}
	}
	if query.RowsAffected == 0 {
		return models.IUser{}, *error.New("User not found", 404, errors.New("User not found"))
	}

	userModel = utils.MapperToUser(user)
	return userModel, error.Error{}
}

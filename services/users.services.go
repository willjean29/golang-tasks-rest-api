package services

import (
	"app/db"
	"app/error"
	"app/models"
	"errors"

	"gorm.io/gorm"
)

type UserService struct{}

func (u *UserService) CreateUser(user *models.User) error.Error {
	query := db.DB.Create(&user)
	err := query.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *error.New("User not craeted", 400, errors.New(err.Error()))
		} else {
			return *error.New("Error create user", 500, errors.New(err.Error()))
		}
	}

	if query.RowsAffected == 0 {
		return *error.New("User not created", 400, errors.New("User not created"))
	}

	return error.Error{}
}

func (u *UserService) GetUserByEmail(email string) (models.User, error.Error) {
	var user models.User
	query := db.DB.Where("email = ?", email).First(&user)
	err := query.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, *error.New("User not found", 404, errors.New(err.Error()))
		} else {
			return models.User{}, *error.New("Error get user", 500, errors.New(err.Error()))
		}
	}
	if query.RowsAffected == 0 {
		return models.User{}, *error.New("User not created", 400, errors.New("User not created"))
	}
	return user, error.Error{}
}

func (s *UserService) GetUsers() (models.ListUsers, error.Error) {
	var users models.ListUsers
	query := db.DB.Preload("Tasks").Find(&users)
	err := query.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ListUsers{}, *error.New("User not found", 404, errors.New(err.Error()))
		} else {
			return models.ListUsers{}, *error.New("Error get users", 500, errors.New(err.Error()))
		}
	}
	if query.RowsAffected == 0 {
		return models.ListUsers{}, *error.New("Users not found", 404, errors.New("Users not found"))
	}
	return users, error.Error{}
}

func (s *UserService) GetUser(id int) (models.User, error.Error) {
	var user models.User
	query := db.DB.Preload("Tasks").First(&user, id)
	err := query.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, *error.New("User not found", 404, errors.New(err.Error()))
		} else {
			return models.User{}, *error.New("Error get user", 500, errors.New(err.Error()))
		}
	}
	if query.RowsAffected == 0 {
		return models.User{}, *error.New("User not found", 404, errors.New("User not found"))
	}
	return user, error.Error{}
}

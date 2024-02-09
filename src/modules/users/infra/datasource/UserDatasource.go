package datasource

import (
	"app/src/modules/users/domain/models"
	"app/src/modules/users/infra/data/gorm/entities"
	"app/src/modules/users/infra/data/gorm/utils"
	db "app/src/shared/data/gorm"
	error "app/src/shared/errors"
	"errors"
	"log"

	"gorm.io/gorm"
)

type UserDatasource struct{}

func (u *UserDatasource) FindAll() (models.IListUsers, error.Error) {
	var listUsers entities.ListUsers
	var users models.IListUsers

	query := db.DB.Preload("Tasks").Find(&listUsers)
	err := query.Error

	if err != nil {
		return models.IListUsers{}, *error.New("Data of users not found", 404, err)
	}

	users = utils.MapperToUsers(listUsers)
	return users, error.Error{}
}

func (u *UserDatasource) FindById(id int) (models.IUser, error.Error) {
	var user entities.User
	var userModel models.IUser

	query := db.DB.Preload("Tasks").First(&user, id)
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

func (u *UserDatasource) FindByEmail(email string) (models.IUser, error.Error) {
	var user entities.User
	var userModel models.IUser

	query := db.DB.Where("email = ?", email).First(&user)
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

func (u *UserDatasource) Create(createUser models.ICreateUser) (models.IUser, error.Error) {
	user := entities.User{
		Name:     createUser.Name,
		Email:    createUser.Email,
		Password: createUser.Password,
	}

	query := db.DB.Create(&user)
	if query.RowsAffected == 0 {
		return models.IUser{}, *error.New("User not created", 400, errors.New("User not created"))
	}

	err := query.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.IUser{}, *error.New("User not created", 400, errors.New(err.Error()))
		} else {
			return models.IUser{}, *error.New("Error create user", 500, errors.New(err.Error()))
		}
	}
	userModel := utils.MapperToUser(user)
	return userModel, error.Error{}
}

func (u *UserDatasource) Save(user models.IUser) error.Error {
	log.Println(user)
	userEntity := utils.MapperToUserEntity(user)
	log.Println(userEntity)

	query := db.DB.Where("id = ?", user.ID).Updates(&userEntity)
	err := query.Error
	if query.RowsAffected == 0 {
		return *error.New("User not created", 400, errors.New("User not created"))
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *error.New("User not updated", 400, errors.New(err.Error()))
		} else {
			return *error.New("Error update user", 500, errors.New(err.Error()))
		}
	}
	return error.Error{}
}

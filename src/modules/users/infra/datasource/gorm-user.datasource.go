package datasource

import (
	"app/src/modules/users/domain/entities"
	"app/src/modules/users/infra/data/gorm/models"
	"app/src/modules/users/infra/data/gorm/utils"
	db "app/src/shared/data/gorm"
	error "app/src/shared/errors"
	"errors"
	"log"

	"gorm.io/gorm"
)

type GormUserDatasource struct{}

func (g *GormUserDatasource) FindAll() (entities.ListUsers, error.Error) {
	var listUsers models.ListUsers
	var users entities.ListUsers

	query := db.DB.Preload("Tasks").Find(&listUsers)
	err := query.Error

	if err != nil {
		return entities.ListUsers{}, *error.New("Data of users not found", 404, err)
	}

	users = utils.MapperToUsers(listUsers)
	return users, error.Error{}
}

func (g *GormUserDatasource) FindById(id int) (entities.User, error.Error) {
	var user models.User
	var userEntity entities.User

	query := db.DB.Preload("Tasks").First(&user, id)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, *error.New("User not found", 404, errors.New(err.Error()))
		} else {
			return entities.User{}, *error.New("Error get user", 500, errors.New(err.Error()))
		}
	}
	if query.RowsAffected == 0 {
		return entities.User{}, *error.New("User not found", 404, errors.New("User not found"))
	}

	userEntity = utils.MapperToUser(user)
	return userEntity, error.Error{}
}

func (g *GormUserDatasource) FindByEmail(email string) (entities.User, error.Error) {
	var user models.User
	var userEntity entities.User

	query := db.DB.Where("email = ?", email).First(&user)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, *error.New("User not found", 404, errors.New(err.Error()))
		} else {
			return entities.User{}, *error.New("Error get user", 500, errors.New(err.Error()))
		}
	}
	if query.RowsAffected == 0 {
		return entities.User{}, *error.New("User not found", 404, errors.New("User not found"))
	}

	userEntity = utils.MapperToUser(user)
	return userEntity, error.Error{}
}

func (g *GormUserDatasource) Create(createUser entities.CreateUser) (entities.User, error.Error) {
	user := models.User{
		Name:     createUser.Name,
		Email:    createUser.Email,
		Password: createUser.Password,
	}

	query := db.DB.Create(&user)
	if query.RowsAffected == 0 {
		return entities.User{}, *error.New("User not created", 400, errors.New("User not created"))
	}

	err := query.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, *error.New("User not created", 400, errors.New(err.Error()))
		} else {
			return entities.User{}, *error.New("Error create user", 500, errors.New(err.Error()))
		}
	}
	userEntity := utils.MapperToUser(user)
	return userEntity, error.Error{}
}

func (g *GormUserDatasource) Save(user entities.User) error.Error {
	log.Println(user)
	userModel := utils.MapperToUserModel(user)
	log.Println(userModel)

	query := db.DB.Where("id = ?", user.ID).Updates(&userModel)
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

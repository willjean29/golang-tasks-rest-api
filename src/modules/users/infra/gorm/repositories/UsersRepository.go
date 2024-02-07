package repositories

import (
	"app/src/modules/users/domain/models"
	"app/src/modules/users/infra/gorm/entities"
	"app/src/modules/users/infra/utils"
	error "app/src/shared/errors"

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

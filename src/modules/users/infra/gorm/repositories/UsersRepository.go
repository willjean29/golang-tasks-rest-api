package repositories

import (
	"app/src/modules/users/domain/models"
	"app/src/modules/users/infra/gorm/entities"

	error "app/src/shared/errors"

	"gorm.io/gorm"
)

type UsersRepository struct {
	Repository *gorm.DB
}

func (u *UsersRepository) FindAll() (models.IListUsers, error.Error) {
	var listUsers entities.ListUsers
	var users models.IListUsers

	query := u.Repository.Find(&listUsers)
	err := query.Error

	if err != nil {
		return models.IListUsers{}, *error.New("Data of users not found", 404, err)
	}

	for _, value := range listUsers {
		user := models.IUser{
			ID:        value.ID,
			Name:      value.Name,
			Email:     value.Email,
			Password:  value.Password,
			Image:     value.Image,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		users = append(users, user)
	}
	return users, error.Error{}
}

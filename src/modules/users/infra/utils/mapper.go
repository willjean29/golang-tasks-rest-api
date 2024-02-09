package utils

import (
	taskModels "app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/infra/data/gorm/utils"
	"app/src/modules/users/domain/models"
	"app/src/modules/users/infra/gorm/entities"
)

func MapperToUserEntity(mapper models.IUser) entities.User {
	return entities.User{
		Name:     mapper.Name,
		Email:    mapper.Email,
		Password: mapper.Password,
		Image:    mapper.Image,
	}
}

func MapperToUser(user entities.User) models.IUser {
	tasks := []taskModels.ITask{}
	if len(user.Tasks) > 0 {
		tasks = utils.MapperToTasks(user.Tasks)
	}

	return models.IUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Tasks:     tasks,
	}
}

func MapperToUsers(list entities.ListUsers) models.IListUsers {
	var listUser models.IListUsers
	for _, value := range list {
		user := MapperToUser(value)
		listUser = append(listUser, user)
	}
	return listUser
}

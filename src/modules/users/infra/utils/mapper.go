package utils

import (
	taskModels "app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/infra/utils"
	"app/src/modules/users/domain/models"
	"app/src/modules/users/infra/gorm/entities"
)

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

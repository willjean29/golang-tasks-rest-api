package utils

import (
	taskEntities "app/src/modules/tasks/domain/entities"
	"app/src/modules/tasks/infra/data/gorm/utils"
	"app/src/modules/users/domain/entities"
	"app/src/modules/users/infra/data/gorm/models"
)

func MapperToUserModel(mapper entities.User) models.User {
	return models.User{
		Name:     mapper.Name,
		Email:    mapper.Email,
		Password: mapper.Password,
		Image:    mapper.Image,
	}
}

func MapperToUser(user models.User) entities.User {
	tasks := []taskEntities.Task{}
	if len(user.Tasks) > 0 {
		tasks = utils.MapperToTasks(user.Tasks)
	}

	return entities.User{
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

func MapperToUsers(list models.ListUsers) entities.ListUsers {
	var listUser entities.ListUsers
	for _, value := range list {
		user := MapperToUser(value)
		listUser = append(listUser, user)
	}
	return listUser
}

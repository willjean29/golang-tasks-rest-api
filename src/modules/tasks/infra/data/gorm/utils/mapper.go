package utils

import (
	"app/src/modules/tasks/domain/entities"
	"app/src/modules/tasks/infra/data/gorm/models"
)

func MapperToTaskModel(task entities.Task) models.Task {
	return models.Task{
		Name:    task.Name,
		Content: task.Content,
		Image:   task.Image,
	}
}

func MapperToTask(task models.Task) entities.Task {
	return entities.Task{
		ID:        task.ID,
		Name:      task.Name,
		Content:   task.Content,
		Image:     task.Image,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

func MapperToTasks(list models.ListTask) entities.ListTask {
	var listTask entities.ListTask
	for _, value := range list {
		task := MapperToTask(value)
		listTask = append(listTask, task)
	}
	return listTask
}

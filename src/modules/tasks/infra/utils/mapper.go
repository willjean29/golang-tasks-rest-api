package utils

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/infra/gorm/entities"
)

func MapperToTaskEntity(task models.ITask) entities.Task {
	return entities.Task{
		Name:    task.Name,
		Content: task.Content,
		Image:   task.Image,
	}
}

func MapperToTask(task entities.Task) models.ITask {
	return models.ITask{
		ID:        task.ID,
		Name:      task.Name,
		Content:   task.Content,
		Image:     task.Image,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

func MapperToTasks(list entities.ListTask) models.IListTask {
	var listTask models.IListTask
	for _, value := range list {
		task := MapperToTask(value)
		listTask = append(listTask, task)
	}
	return listTask
}

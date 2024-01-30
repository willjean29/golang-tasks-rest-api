package services

import (
	"app/db"
	"app/models"
	"errors"
)

type TaskService struct{}

func (t *TaskService) GetTasks() (models.ListTask, error) {
	var tasks models.ListTask
	query := db.DB.Find(&tasks)
	err := query.Error
	if err != nil {
		return models.ListTask{}, errors.New("Data of tasks not found")
	}
	return tasks, nil
}

func (t *TaskService) GetTask(id int) (models.Task, error) {
	var task models.Task
	query := db.DB.First(&task, id)
	err := query.Error
	if err != nil {
		return models.Task{}, errors.New("Task not found")
	}
	return task, nil
}

func (t *TaskService) CreateTask(task models.Task) (models.Task, error) {
	query := db.DB.Create(&task)
	err := query.Error
	if err != nil {
		return models.Task{}, errors.New("Task not created")
	}
	return task, nil
}

func (t *TaskService) UpdateTask(task models.Task, id int) (models.Task, error) {
	query := db.DB.Where("id = ?", id).Updates(&task)
	err := query.RowsAffected
	if err == 0 {
		return models.Task{}, errors.New("Task not updated")
	}
	return task, nil
}

func (t *TaskService) DeleteTask(id int) error {
	var task models.Task
	query := db.DB.Delete(&task, id)
	err := query.RowsAffected
	if err == 0 {
		return errors.New("Task not deleted")
	}
	return nil
}

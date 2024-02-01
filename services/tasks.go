package services

import (
	"app/db"
	"app/error"
	"app/models"
	"errors"

	"gorm.io/gorm"
)

type TaskService struct{}

func (t *TaskService) GetTasks() (models.ListTask, error.Error) {
	var tasks models.ListTask
	query := db.DB.Find(&tasks)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ListTask{}, *error.New("Data of tasks not found", 404, errors.New(err.Error()))
		} else {
			return models.ListTask{}, *error.New("Error get list task", 500, errors.New(err.Error()))
		}
	}
	return tasks, error.Error{}
}

func (t *TaskService) GetTask(id int) (models.Task, error.Error) {
	var task models.Task
	query := db.DB.First(&task, id)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Task{}, *error.New("Task not found", 404, errors.New(err.Error()))
		} else {
			return models.Task{}, *error.New("Error get task", 500, errors.New(err.Error()))
		}
	}

	return task, error.Error{}
}

func (t *TaskService) CreateTask(task models.Task) (models.Task, error.Error) {
	query := db.DB.Create(&task)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Task{}, *error.New("Task not created", 400, errors.New(err.Error()))
		} else {
			return models.Task{}, *error.New("Error create task", 500, errors.New(err.Error()))
		}
	}
	return task, error.Error{}
}

func (t *TaskService) UpdateTask(task models.Task, id int) (models.Task, error.Error) {
	query := db.DB.Where("id = ?", id).Updates(&task)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Task{}, *error.New("Task not found", 404, errors.New(err.Error()))
		} else {
			return models.Task{}, *error.New("Error update task", 500, errors.New(err.Error()))
		}
	}
	if query.RowsAffected == 0 {
		return models.Task{}, *error.New("Task not found", 404, errors.New("No task was updated"))
	}
	return task, error.Error{}
}

func (t *TaskService) DeleteTask(id int) error.Error {
	var task models.Task
	query := db.DB.Delete(&task, id)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *error.New("Task not found", 404, errors.New(err.Error()))
		} else {
			return *error.New("Error get task", 500, errors.New(err.Error()))
		}
	}

	if query.RowsAffected == 0 {
		return *error.New("Task not found", 404, errors.New("No task was deleted"))
	}
	return error.Error{}
}

func (t *TaskService) SaveTask(task *models.Task) error.Error {
	query := db.DB.Save(&task)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *error.New("Task not found", 404, errors.New(err.Error()))
		} else {
			return *error.New("Error get task", 500, errors.New(err.Error()))
		}
	}

	if query.RowsAffected == 0 {
		return *error.New("Task not found", 404, errors.New("No task was deleted"))
	}
	return error.Error{}
}

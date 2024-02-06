package repositories

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/infra/gorm/entities"
	error "app/src/shared/errors"
	"errors"

	"gorm.io/gorm"
)

type TasksRepository struct {
	Repository *gorm.DB
}

func (t *TasksRepository) FindAll() (models.IListTask, error.Error) {
	var listTask models.IListTask
	var list entities.ListTask

	query := t.Repository.Find(&list)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.IListTask{}, *error.New("Data of tasks not found", 404, errors.New(err.Error()))
		} else {
			return models.IListTask{}, *error.New("Error get list task", 500, errors.New(err.Error()))
		}
	}

	for _, value := range list {
		task := models.ITask{
			ID:        value.ID,
			Name:      value.Name,
			Content:   value.Content,
			Image:     value.Image,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		listTask = append(listTask, task)
	}
	return listTask, error.Error{}
}

func (t *TasksRepository) FindById(id int) (models.ITask, error.Error) {

	var task entities.Task

	query := t.Repository.First(&task, id)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ITask{}, *error.New("Task not found", 404, errors.New(err.Error()))
		} else {
			return models.ITask{}, *error.New("Error get task", 500, errors.New(err.Error()))
		}
	}

	taskModel := models.ITask{
		ID:        task.ID,
		Name:      task.Name,
		Content:   task.Content,
		Image:     task.Image,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return taskModel, error.Error{}
}

func (t *TasksRepository) Create(createTask models.ICreateTask) (models.ITask, error.Error) {
	task := entities.Task{
		Name:    createTask.Name,
		Content: createTask.Content,
		UserID:  4,
	}

	query := t.Repository.Create(&task)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ITask{}, *error.New("Task not created", 400, errors.New(err.Error()))
		} else {
			return models.ITask{}, *error.New("Error create task", 500, errors.New(err.Error()))
		}
	}
	taskModel := models.ITask{
		ID:        task.ID,
		Name:      task.Name,
		Content:   task.Content,
		Image:     task.Image,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return taskModel, error.Error{}
}

func (t *TasksRepository) Delete(id int) error.Error {
	var task entities.Task
	query := t.Repository.Delete(&task, id)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *error.New("Task not found", 404, errors.New(err.Error()))
		} else {
			return *error.New("Error delete task", 500, errors.New(err.Error()))
		}
	}

	if query.RowsAffected == 0 {
		return *error.New("Task not found", 404, errors.New("No task was updated"))
	}

	return error.Error{}
}

func (t *TasksRepository) Update(updateTask models.IUpdateTask, id int) (models.ITask, error.Error) {
	task := entities.Task{
		Name:    updateTask.Name,
		Content: updateTask.Content,
	}

	query := t.Repository.Where("id = ?", id).Updates(&task)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ITask{}, *error.New("Task not found", 404, errors.New(err.Error()))
		} else {
			return models.ITask{}, *error.New("Error update task", 500, errors.New(err.Error()))
		}
	}
	if query.RowsAffected == 0 {
		return models.ITask{}, *error.New("Task not found", 404, errors.New("No task was updated"))
	}
	taskModel := models.ITask{
		ID:        task.ID,
		Name:      task.Name,
		Content:   task.Content,
		Image:     task.Image,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return taskModel, error.Error{}
}

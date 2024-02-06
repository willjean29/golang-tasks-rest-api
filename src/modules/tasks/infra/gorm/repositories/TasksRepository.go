package repositories

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/infra/gorm/entities"
	"errors"

	"gorm.io/gorm"
)

type TasksRepository struct {
	Repository *gorm.DB
}

func (t *TasksRepository) FindAll() (models.IListTask, error) {
	var listTask models.IListTask
	var list entities.ListTask

	query := t.Repository.Find(&list)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.IListTask{}, errors.New("Data of tasks not found")
		} else {
			return models.IListTask{}, errors.New("Error get list task")
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
	return listTask, nil
}

func (t *TasksRepository) FindById(id int) (models.ITask, error) {

	var task entities.Task

	query := t.Repository.First(&task, id)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ITask{}, errors.New("Data of task not found")
		} else {
			return models.ITask{}, errors.New("Error get task")
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

	return taskModel, nil
}

func (t *TasksRepository) Create(createTask models.ICreateTask) (models.ITask, error) {
	task := entities.Task{
		Name:    createTask.Name,
		Content: createTask.Content,
		UserID:  4,
	}

	query := t.Repository.Create(&task)
	err := query.Error

	if err != nil {
		return models.ITask{}, errors.New("Error create task")
	}
	taskModel := models.ITask{
		ID:        task.ID,
		Name:      task.Name,
		Content:   task.Content,
		Image:     task.Image,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return taskModel, nil
}

func (t *TasksRepository) Delete(id int) error {
	var task entities.Task
	query := t.Repository.Delete(&task, id)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Task not found")
		} else {
			return errors.New("Error delete task")
		}
	}

	return nil
}

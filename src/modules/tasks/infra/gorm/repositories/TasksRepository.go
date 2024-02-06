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

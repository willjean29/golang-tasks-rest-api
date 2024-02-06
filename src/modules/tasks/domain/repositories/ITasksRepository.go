package repositories

import (
	"app/src/modules/tasks/domain/models"
	error "app/src/shared/errors"
)

type ITaskRepository interface {
	FindAll() (models.IListTask, error.Error)
	FindById(id int) (models.ITask, error.Error)
	Create(task models.ICreateTask) (models.ITask, error.Error)
	// Save(task models.ITask) error
	Update(task models.IUpdateTask, id int) (models.ITask, error.Error)
	Delete(id int) error.Error
}

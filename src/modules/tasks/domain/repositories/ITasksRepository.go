package repositories

import "app/src/modules/tasks/domain/models"

type ITaskRepository interface {
	FindAll() (models.IListTask, error)
	FindById(id int) (models.ITask, error)
	Create(task models.ICreateTask) (models.ITask, error)
	// Save(task models.ITask) error
	// Update(task models.IUpdateTask) (models.ITask, error)
	Delete(id int) error
}

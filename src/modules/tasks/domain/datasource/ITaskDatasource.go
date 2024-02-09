package datasource

import (
	"app/src/modules/tasks/domain/models"
	error "app/src/shared/errors"
)

type ITaskDatasource interface {
	FindAll(userId uint) (models.IListTask, error.Error)
	FindById(id int) (models.ITask, error.Error)
	Create(task models.ICreateTask, userId uint) (models.ITask, error.Error)
	Save(task models.ITask) error.Error
	Update(task models.IUpdateTask, id int) (models.ITask, error.Error)
	Delete(id int) error.Error
}

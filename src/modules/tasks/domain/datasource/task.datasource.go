package datasource

import (
	"app/src/modules/tasks/domain/entities"
	error "app/src/shared/errors"
)

type ITaskDatasource interface {
	FindAll(userId uint) (entities.ListTask, error.Error)
	FindById(id int) (entities.Task, error.Error)
	Create(task entities.CreateTask, userId uint) (entities.Task, error.Error)
	Save(task entities.Task) error.Error
	Update(task entities.UpdateTask, id int) (entities.Task, error.Error)
	Delete(id int) error.Error
}

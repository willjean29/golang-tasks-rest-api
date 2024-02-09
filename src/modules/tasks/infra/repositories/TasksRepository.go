package repositories

import (
	"app/src/modules/tasks/domain/datasource"
	"app/src/modules/tasks/domain/models"
	error "app/src/shared/errors"
)

type TasksRepository struct {
	Datasource datasource.ITaskDatasource
}

func (t *TasksRepository) FindAll(userId uint) (models.IListTask, error.Error) {
	return t.Datasource.FindAll(userId)
}

func (t *TasksRepository) FindById(id int) (models.ITask, error.Error) {
	return t.Datasource.FindById(id)
}

func (t *TasksRepository) Create(createTask models.ICreateTask, userId uint) (models.ITask, error.Error) {
	return t.Datasource.Create(createTask, userId)
}

func (t *TasksRepository) Delete(id int) error.Error {
	return t.Datasource.Delete(id)
}

func (t *TasksRepository) Update(updateTask models.IUpdateTask, id int) (models.ITask, error.Error) {
	return t.Datasource.Update(updateTask, id)
}

func (t *TasksRepository) Save(task models.ITask) error.Error {
	return t.Datasource.Save(task)
}

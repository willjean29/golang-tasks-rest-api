package repositories

import (
	"app/src/modules/tasks/domain/datasource"
	"app/src/modules/tasks/domain/entities"
	error "app/src/shared/errors"
)

type TasksRepository struct {
	Datasource datasource.TaskDatasource
}

func (t *TasksRepository) FindAll(userId uint) (entities.ListTask, error.Error) {
	return t.Datasource.FindAll(userId)
}

func (t *TasksRepository) FindById(id int) (entities.Task, error.Error) {
	return t.Datasource.FindById(id)
}

func (t *TasksRepository) Create(createTask entities.CreateTask, userId uint) (entities.Task, error.Error) {
	return t.Datasource.Create(createTask, userId)
}

func (t *TasksRepository) Delete(id int) error.Error {
	return t.Datasource.Delete(id)
}

func (t *TasksRepository) Update(updateTask entities.UpdateTask, id int) (entities.Task, error.Error) {
	return t.Datasource.Update(updateTask, id)
}

func (t *TasksRepository) Save(task entities.Task) error.Error {
	return t.Datasource.Save(task)
}

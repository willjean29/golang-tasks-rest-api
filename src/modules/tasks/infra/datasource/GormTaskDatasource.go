package datasource

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/infra/gorm/entities"
	"app/src/modules/tasks/infra/utils"
	error "app/src/shared/errors"
	db "app/src/shared/infra/gorm"
	"errors"

	"gorm.io/gorm"
)

type GormTaskDatasource struct{}

func (g *GormTaskDatasource) FindAll(userId uint) (models.IListTask, error.Error) {
	var listTask models.IListTask
	var list entities.ListTask

	query := db.DB.Where("user_id = ?", userId).Find(&list)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.IListTask{}, *error.New("Data of tasks not found", 404, errors.New(err.Error()))
		} else {
			return models.IListTask{}, *error.New("Error get list task", 500, errors.New(err.Error()))
		}
	}

	listTask = utils.MapperToTasks(list)
	return listTask, error.Error{}
}

func (g *GormTaskDatasource) FindById(id int) (models.ITask, error.Error) {

	var task entities.Task

	query := db.DB.First(&task, id)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ITask{}, *error.New("Task not found", 404, errors.New(err.Error()))
		} else {
			return models.ITask{}, *error.New("Error get task", 500, errors.New(err.Error()))
		}
	}

	taskModel := utils.MapperToTask(task)

	return taskModel, error.Error{}
}

func (g *GormTaskDatasource) Create(createTask models.ICreateTask, userId uint) (models.ITask, error.Error) {
	task := entities.Task{
		Name:    createTask.Name,
		Content: createTask.Content,
		UserID:  userId,
	}

	query := db.DB.Create(&task)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ITask{}, *error.New("Task not created", 400, errors.New(err.Error()))
		} else {
			return models.ITask{}, *error.New("Error create task", 500, errors.New(err.Error()))
		}
	}
	taskModel := utils.MapperToTask(task)
	return taskModel, error.Error{}
}

func (g *GormTaskDatasource) Delete(id int) error.Error {
	var task entities.Task
	query := db.DB.Delete(&task, id)
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

func (g *GormTaskDatasource) Update(updateTask models.IUpdateTask, id int) (models.ITask, error.Error) {
	task := entities.Task{
		Name:    updateTask.Name,
		Content: updateTask.Content,
	}

	query := db.DB.Where("id = ?", id).Updates(&task)
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
	taskModel := utils.MapperToTask(task)
	return taskModel, error.Error{}
}

func (g *GormTaskDatasource) Save(task models.ITask) error.Error {
	taskEntity := utils.MapperToTaskEntity(task)
	query := db.DB.Where("id = ?", task.ID).Updates(&taskEntity)
	err := query.Error
	if query.RowsAffected == 0 {
		return *error.New("Task not created", 400, errors.New("Task not created"))
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *error.New("Task not updated", 400, errors.New(err.Error()))
		} else {
			return *error.New("Error update task", 500, errors.New(err.Error()))
		}
	}
	return error.Error{}
}

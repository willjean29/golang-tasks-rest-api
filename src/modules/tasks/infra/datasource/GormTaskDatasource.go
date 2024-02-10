package datasource

import (
	"app/src/modules/tasks/domain/entities"
	"app/src/modules/tasks/infra/data/gorm/models"
	"app/src/modules/tasks/infra/data/gorm/utils"
	db "app/src/shared/data/gorm"
	error "app/src/shared/errors"
	"errors"

	"gorm.io/gorm"
)

type GormTaskDatasource struct{}

func (g *GormTaskDatasource) FindAll(userId uint) (entities.ListTask, error.Error) {
	var listTask entities.ListTask
	var list models.ListTask

	query := db.DB.Where("user_id = ?", userId).Find(&list)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.ListTask{}, *error.New("Data of tasks not found", 404, errors.New(err.Error()))
		} else {
			return entities.ListTask{}, *error.New("Error get list task", 500, errors.New(err.Error()))
		}
	}

	listTask = utils.MapperToTasks(list)
	return listTask, error.Error{}
}

func (g *GormTaskDatasource) FindById(id int) (entities.Task, error.Error) {

	var task models.Task

	query := db.DB.First(&task, id)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Task{}, *error.New("Task not found", 404, errors.New(err.Error()))
		} else {
			return entities.Task{}, *error.New("Error get task", 500, errors.New(err.Error()))
		}
	}

	taskModel := utils.MapperToTask(task)

	return taskModel, error.Error{}
}

func (g *GormTaskDatasource) Create(createTask entities.CreateTask, userId uint) (entities.Task, error.Error) {
	task := models.Task{
		Name:    createTask.Name,
		Content: createTask.Content,
		UserID:  userId,
	}

	query := db.DB.Create(&task)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Task{}, *error.New("Task not created", 400, errors.New(err.Error()))
		} else {
			return entities.Task{}, *error.New("Error create task", 500, errors.New(err.Error()))
		}
	}
	taskModel := utils.MapperToTask(task)
	return taskModel, error.Error{}
}

func (g *GormTaskDatasource) Delete(id int) error.Error {
	var task models.Task
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

func (g *GormTaskDatasource) Update(updateTask entities.UpdateTask, id int) (entities.Task, error.Error) {
	task := models.Task{
		Name:    updateTask.Name,
		Content: updateTask.Content,
	}

	query := db.DB.Where("id = ?", id).Updates(&task)
	err := query.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Task{}, *error.New("Task not found", 404, errors.New(err.Error()))
		} else {
			return entities.Task{}, *error.New("Error update task", 500, errors.New(err.Error()))
		}
	}
	if query.RowsAffected == 0 {
		return entities.Task{}, *error.New("Task not found", 404, errors.New("No task was updated"))
	}
	taskModel := utils.MapperToTask(task)
	return taskModel, error.Error{}
}

func (g *GormTaskDatasource) Save(task entities.Task) error.Error {
	taskEntity := utils.MapperToTaskModel(task)
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

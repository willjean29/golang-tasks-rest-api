package validators

import (
	"app/dtos"

	"github.com/go-playground/validator/v10"
)

const (
	CreateTask = "CreateTask"
	UpdateTask = "UpdateTask"
)

var validate *validator.Validate

type TaskValidator struct{}

func NewTaskValidator() TaskValidator {
	validate = validator.New(validator.WithRequiredStructEnabled())
	return TaskValidator{}
}

func (t *TaskValidator) ValidateCreateTask(createTaskDto dtos.CreateTaskDto) error {
	var err error
	err = validate.Struct(createTaskDto)
	return err
}

func (t *TaskValidator) ValidateUpdateTask(updateTaskDto dtos.UpdateTaskDto) error {
	var err error
	err = validate.Struct(updateTaskDto)
	return err
}

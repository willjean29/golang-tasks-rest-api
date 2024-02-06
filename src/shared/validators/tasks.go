package validators

import (
	"app/src/modules/tasks/domain/models"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type TaskValidator struct{}

func NewTaskValidator() TaskValidator {
	validate = validator.New(validator.WithRequiredStructEnabled())
	return TaskValidator{}
}

func (t *TaskValidator) ValidateCreateTask(createTaskDto models.ICreateTask) error {
	var err error
	err = validate.Struct(createTaskDto)
	return err
}

func (t *TaskValidator) ValidateUpdateTask(updateTaskDto models.IUpdateTask) error {
	var err error
	err = validate.Struct(updateTaskDto)
	return err
}

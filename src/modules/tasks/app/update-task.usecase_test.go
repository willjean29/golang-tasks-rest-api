package usecases_test

import (
	usecases "app/src/modules/tasks/app"
	"app/src/modules/tasks/domain/entities"
	error "app/src/shared/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateTaskUseCase_Execute(t *testing.T) {
	mock := new(MockTaskRepository)
	useCase := usecases.UpdateTaskUseCase{TaskRepository: mock}
	taskId := 1
	updateTask := entities.UpdateTask{Name: "Task 2", Content: "Description 2"}
	expectedTask := entities.Task{ID: uint(taskId), Name: "Task 2", Content: "Description 2"}
	mock.On("Update", updateTask, taskId).Return(expectedTask, error.Error{})
	task, err := useCase.Execute(updateTask, taskId)
	assert.Equal(t, error.Error{}, err)
	assert.Equal(t, expectedTask, task)
}

func TestUpdateTaskUseCase_Execute_Error(t *testing.T) {
	mock := new(MockTaskRepository)
	useCase := usecases.UpdateTaskUseCase{TaskRepository: mock}
	taskId := 1
	updateTask := entities.UpdateTask{Name: "Task 2", Content: "Description 2"}
	expectedError := error.Error{
		StatusCode: 404,
		Message:    "task not found",
	}
	mock.On("Update", updateTask, taskId).Return(entities.Task{}, expectedError)
	task, err := useCase.Execute(updateTask, taskId)
	assert.Equal(t, entities.Task{}, task)
	assert.Equal(t, expectedError, err)
}

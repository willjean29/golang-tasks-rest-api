package usecases_test

import (
	usecases "app/src/modules/tasks/app"
	"app/src/modules/tasks/domain/entities"
	error "app/src/shared/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTaskUsecase_Execute(t *testing.T) {
	mock := new(MockTaskRepository)
	useCase := usecases.GetTaskUseCase{TaskRepository: mock}

	expectedTask := entities.Task{ID: 1, Name: "Task 1", Content: "Description 1"}

	mock.On("FindById", 1).Return(expectedTask, error.Error{})
	task, err := useCase.Execute(1)
	assert.Equal(t, expectedTask, task)
	assert.Equal(t, error.Error{}, err)
}

func TestGetTaskUsecase_Execute_Error(t *testing.T) {
	mock := new(MockTaskRepository)
	useCase := usecases.GetTaskUseCase{TaskRepository: mock}

	expectedError := error.Error{
		StatusCode: 404,
		Message:    "task not found",
	}

	mock.On("FindById", 1).Return(entities.Task{}, expectedError)
	task, err := useCase.Execute(1)
	assert.Equal(t, entities.Task{}, task)
	assert.Equal(t, expectedError, err)
}

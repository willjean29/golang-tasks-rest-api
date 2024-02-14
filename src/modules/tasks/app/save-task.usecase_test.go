package usecases_test

import (
	usecases "app/src/modules/tasks/app"
	"app/src/modules/tasks/domain/entities"
	error "app/src/shared/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskUseCase(t *testing.T) {
	mock := new(MockTaskRepository)
	useCase := usecases.SaveTaskUseCase{TaskRepository: mock}

	expectedTask := entities.Task{ID: 1, Name: "Task 1", Content: "Description 1"}

	mock.On("Save", expectedTask).Return(error.Error{})
	err := useCase.Execute(expectedTask)
	assert.Equal(t, error.Error{}, err)
}

func TestTaskUseCase_Error(t *testing.T) {
	mock := new(MockTaskRepository)
	useCase := usecases.SaveTaskUseCase{TaskRepository: mock}

	expectedTask := entities.Task{ID: 1, Name: "Task 1", Content: "Description 1"}
	expectedError := error.Error{
		StatusCode: 500,
		Message:    "Internal server error",
	}

	mock.On("Save", expectedTask).Return(expectedError)
	err := useCase.Execute(expectedTask)
	assert.Equal(t, expectedError, err)
}

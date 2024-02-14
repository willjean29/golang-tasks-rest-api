package usecases_test

import (
	usecases "app/src/modules/tasks/app"
	error "app/src/shared/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteTaskUseCase_Execute(t *testing.T) {
	mock := new(MockTaskRepository)
	useCase := usecases.DeleteTaskUseCase{TaskRepository: mock}
	taskId := 1
	expectedError := error.Error{}
	mock.On("Delete", taskId).Return(expectedError)
	err := useCase.Execute(taskId)
	assert.Equal(t, expectedError, err)
}

func TestDeleteTaskUseCase_Execute_Error(t *testing.T) {
	mock := new(MockTaskRepository)
	useCase := usecases.DeleteTaskUseCase{TaskRepository: mock}
	taskId := 1
	expectedError := error.Error{
		StatusCode: 404,
		Message:    "task not found",
	}
	mock.On("Delete", taskId).Return(expectedError)
	err := useCase.Execute(taskId)
	assert.Equal(t, expectedError, err)
}

package usecases_test

import (
	usecases "app/src/modules/tasks/app"
	"app/src/modules/tasks/domain/entities"
	error "app/src/shared/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListTaskUseCase_Execute(t *testing.T) {
	// Configurar el mock del repositorio de tareas
	mockRepo := new(MockTaskRepository)

	// Configurar el caso de uso con el repositorio mock
	useCase := usecases.ListTasksUseCase{TaskRepository: mockRepo}

	// Definir los datos de prueba
	userId := uint(1)
	expectedTasks := entities.ListTask{
		{ID: 1, Name: "Task 1", Content: "Description 1"},
		{ID: 2, Name: "Task 2", Content: "Description 2"},
	}

	// Configurar el comportamiento esperado del mock
	mockRepo.On("FindAll", userId).Return(expectedTasks, error.Error{})

	// Ejecutar el caso de uso
	task, err := useCase.Execute(userId)

	// Validar el resultado
	assert.Equal(t, expectedTasks, task)
	assert.Equal(t, error.Error{}, err)
}

func TestListTaskUseCase_Execute_Error(t *testing.T) {
	// Configurar el mock del repositorio de tareas
	mockRepo := new(MockTaskRepository)

	// Configurar el caso de uso con el repositorio mock
	useCase := usecases.ListTasksUseCase{TaskRepository: mockRepo}

	// Definir los datos de prueba
	userId := uint(1)
	expectedError := error.Error{
		StatusCode: 500,
		Message:    "Internal server error",
	}

	// Configurar el comportamiento esperado del mock
	mockRepo.On("FindAll", userId).Return(entities.ListTask{}, expectedError)

	// Ejecutar el caso de uso
	task, err := useCase.Execute(userId)

	// Validar el resultado
	assert.Equal(t, entities.ListTask{}, task)
	assert.Equal(t, expectedError, err)
}

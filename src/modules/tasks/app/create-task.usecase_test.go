package usecases_test

import (
	usecases "app/src/modules/tasks/app"
	"app/src/modules/tasks/domain/entities"
	error "app/src/shared/errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Definir un mock del repositorio de tareas
type MockTaskRepository struct {
	mock.Mock
}

// Implementar el método Create del repositorio de tareas
func (m *MockTaskRepository) Create(createProduct entities.CreateTask, userId uint) (entities.Task, error.Error) {
	args := m.Called(createProduct, userId)
	return args.Get(0).(entities.Task), args.Get(1).(error.Error)
}

func (m *MockTaskRepository) FindAll(userId uint) (entities.ListTask, error.Error) {
	args := m.Called(userId)
	return args.Get(0).(entities.ListTask), args.Get(1).(error.Error)
}

func (m *MockTaskRepository) FindById(id int) (entities.Task, error.Error) {
	args := m.Called(id)
	return args.Get(0).(entities.Task), args.Get(1).(error.Error)
}

func (m *MockTaskRepository) Save(task entities.Task) error.Error {
	args := m.Called(task)
	return args.Get(0).(error.Error)
}

func (m *MockTaskRepository) Update(task entities.UpdateTask, id int) (entities.Task, error.Error) {
	args := m.Called(task, id)
	return args.Get(0).(entities.Task), args.Get(1).(error.Error)
}

func (m *MockTaskRepository) Delete(id int) error.Error {
	args := m.Called(id)
	return args.Get(0).(error.Error)
}

func TestCreateTaskUseCase_Execute(t *testing.T) {
	// Configurar el mock del repositorio de tareas
	mockRepo := new(MockTaskRepository)

	// Configurar el caso de uso con el repositorio mock
	useCase := usecases.CreateTaskUseCase{TaskRepository: mockRepo}

	// Definir los datos de prueba
	createProduct := entities.CreateTask{Name: "Task 1", Content: "Description 1"}
	userId := uint(1)
	expectedTask := entities.Task{ID: 1, Name: "Task 1", Content: "Description 1"}

	// Configurar expectativas en el mock del repositorio
	mockRepo.On("Create", createProduct, userId).Return(expectedTask, error.Error{})

	// Ejecutar la función del caso de uso
	task, err := useCase.Execute(createProduct, userId)

	// Verificar que se llamó al método Create del repositorio mock con los argumentos adecuados
	mockRepo.AssertCalled(t, "Create", createProduct, userId)

	// Verificar que se obtuvo el resultado esperado
	assert.Equal(t, err.StatusCode, 0)  // Verificar que no haya error
	assert.Equal(t, expectedTask, task) // Verificar que el resultado coincida con el esperado
}

func TestCreateTaskUseCase_Execute_Error(t *testing.T) {
	// Configurar el mock del repositorio de tareas
	mockRepo := new(MockTaskRepository)

	// Configurar el caso de uso con el repositorio mock
	useCase := usecases.CreateTaskUseCase{TaskRepository: mockRepo}

	// Definir los datos de prueba
	createProduct := entities.CreateTask{Name: "Task 1", Content: "Description 1"}
	userId := uint(1)
	expectedError := error.Error{StatusCode: 400, Message: "Error"}

	// Configurar expectativas en el mock del repositorio
	mockRepo.On("Create", createProduct, userId).Return(entities.Task{}, expectedError)

	// Ejecutar la función del caso de uso
	task, err := useCase.Execute(createProduct, userId)

	// Verificar que se llamó al método Create del repositorio mock con los argumentos adecuados
	mockRepo.AssertCalled(t, "Create", createProduct, userId)

	// Verificar que se obtuvo el resultado esperado
	assert.Equal(t, expectedError, err)    // Verificar que se obtuvo el error esperado
	assert.Equal(t, entities.Task{}, task) // Verificar que no se obtuvo una tarea
}

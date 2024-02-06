package usecases

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/domain/repositories"
	error "app/src/shared/errors"
)

type CreateTaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *CreateTaskUseCase) Execute(createProduct models.ICreateTask) (models.ITask, error.Error) {
	tasks, err := l.TaskRepository.Create(createProduct)
	if err.StatusCode != 0 {
		return models.ITask{}, err
	}
	return tasks, error.Error{}
}

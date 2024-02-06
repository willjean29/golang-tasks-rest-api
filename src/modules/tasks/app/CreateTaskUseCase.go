package usecases

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/domain/repositories"
)

type CreateTaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *CreateTaskUseCase) Execute(createProduct models.ICreateTask) (models.ITask, error) {
	tasks, err := l.TaskRepository.Create(createProduct)
	if err != nil {
		return models.ITask{}, err
	}
	return tasks, nil
}

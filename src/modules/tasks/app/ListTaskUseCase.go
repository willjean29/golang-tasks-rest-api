package usecases

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/domain/repositories"
)

type ListTasksUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *ListTasksUseCase) Execute() (models.IListTask, error) {
	tasks, err := l.TaskRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

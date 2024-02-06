package usecases

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/domain/repositories"
)

type GetTaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *GetTaskUseCase) Execute(id int) (models.ITask, error) {
	tasks, err := l.TaskRepository.FindById(id)
	if err != nil {
		return models.ITask{}, err
	}
	return tasks, nil
}

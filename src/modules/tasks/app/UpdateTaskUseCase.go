package usecases

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/domain/repositories"
)

type UpdateTaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *UpdateTaskUseCase) Execute(updateProduct models.IUpdateTask, id int) (models.ITask, error) {
	tasks, err := l.TaskRepository.Update(updateProduct, id)
	if err != nil {
		return models.ITask{}, err
	}
	return tasks, nil
}

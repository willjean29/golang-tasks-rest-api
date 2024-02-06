package usecases

import (
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/domain/repositories"
	error "app/src/shared/errors"
)

type UpdateTaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func (l *UpdateTaskUseCase) Execute(updateProduct models.IUpdateTask, id int) (models.ITask, error.Error) {
	tasks, err := l.TaskRepository.Update(updateProduct, id)
	if err.StatusCode != 0 {
		return models.ITask{}, err
	}
	return tasks, error.Error{}
}

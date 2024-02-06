package controllers

import (
	"app/db"
	usecases "app/src/modules/tasks/app"
	"app/src/modules/tasks/infra/gorm/repositories"
	"encoding/json"
	"net/http"
)

type TasksController struct{}

func (t *TasksController) List(w http.ResponseWriter, r *http.Request) {

	usecase := usecases.ListTasksUseCase{
		TaskRepository: &repositories.TasksRepository{
			Repository: db.DB,
		},
	}

	tasks, err := usecase.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

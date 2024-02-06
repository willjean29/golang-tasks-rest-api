package controllers

import (
	"app/db"
	usecases "app/src/modules/tasks/app"
	"app/src/modules/tasks/infra/gorm/repositories"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (t *TasksController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("Invalid ID"))
		return
	}
	usecase := usecases.GetTaskUseCase{
		TaskRepository: &repositories.TasksRepository{
			Repository: db.DB,
		},
	}

	tasks, err := usecase.Execute(taskID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

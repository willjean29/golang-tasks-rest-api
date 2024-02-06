package controllers

import (
	"app/db"
	usecases "app/src/modules/tasks/app"
	"app/src/modules/tasks/domain/models"
	"app/src/modules/tasks/infra/gorm/repositories"
	"app/utils"
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

func (t *TasksController) Create(w http.ResponseWriter, r *http.Request) {
	var createTask models.ICreateTask

	err := utils.TransformBody(r.Body, &createTask)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("Invalid data"))
		return
	}
	usecase := usecases.CreateTaskUseCase{
		TaskRepository: &repositories.TasksRepository{
			Repository: db.DB,
		},
	}

	task, err := usecase.Execute(createTask)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (t *TasksController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("Invalid ID"))
		return
	}
	usecase := usecases.DeleteTaskUseCase{
		TaskRepository: &repositories.TasksRepository{
			Repository: db.DB,
		},
	}

	err = usecase.Execute(taskID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "The task with ID " + strconv.Itoa(taskID) + " has been remove successfully",
	})
}

func (t *TasksController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("Invalid ID"))
		return
	}
	var updateTask models.IUpdateTask

	err = utils.TransformBody(r.Body, &updateTask)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("Please enter valid data"))
		return
	}
	usecase := usecases.UpdateTaskUseCase{
		TaskRepository: &repositories.TasksRepository{
			Repository: db.DB,
		},
	}

	_, err = usecase.Execute(updateTask, taskID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "The task with ID " + strconv.Itoa(taskID) + " has been update successfully",
	})
}
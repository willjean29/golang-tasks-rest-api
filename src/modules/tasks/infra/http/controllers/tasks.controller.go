package controllers

import (
	usecases "app/src/modules/tasks/app"
	"app/src/modules/tasks/domain/entities"
	"app/src/modules/tasks/domain/validators"
	"app/src/modules/tasks/infra/datasource"
	"app/src/modules/tasks/infra/repositories"
	error "app/src/shared/errors"

	"app/src/shared/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var taskValidator validators.TaskValidator = validators.NewTaskValidator()
var taskRepository = &repositories.TasksRepository{
	Datasource: &datasource.GormTaskDatasource{},
}

type TasksController struct{}

func (t *TasksController) List(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int)
	usecase := usecases.ListTasksUseCase{
		TaskRepository: taskRepository,
	}

	tasks, errorApp := usecase.Execute(uint(userId))

	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
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
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}
	usecase := usecases.GetTaskUseCase{
		TaskRepository: taskRepository,
	}

	tasks, errorApp := usecase.Execute(taskID)

	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (t *TasksController) Create(w http.ResponseWriter, r *http.Request) {
	var createTask entities.CreateTask

	err := utils.TransformBody(r.Body, &createTask)
	userId := r.Context().Value("userId").(int)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}
	usecase := usecases.CreateTaskUseCase{
		TaskRepository: taskRepository,
	}
	err = taskValidator.ValidateCreateTask(createTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	task, errorApp := usecase.Execute(createTask, uint(userId))

	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
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
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}
	usecase := usecases.DeleteTaskUseCase{
		TaskRepository: taskRepository,
	}

	errorApp := usecase.Execute(taskID)

	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
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
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}
	var updateTask entities.UpdateTask

	err = utils.TransformBody(r.Body, &updateTask)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Please enter valid data", http.StatusBadRequest, err))
		return
	}
	err = taskValidator.ValidateUpdateTask(updateTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}
	usecase := usecases.UpdateTaskUseCase{
		TaskRepository: taskRepository,
	}

	_, errorApp := usecase.Execute(updateTask, taskID)

	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "The task with ID " + strconv.Itoa(taskID) + " has been update successfully",
	})
}

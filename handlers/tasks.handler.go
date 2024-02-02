package handlers

import (
	"app/dtos"
	"app/error"
	"app/models"
	"app/services"
	"app/utils"
	"app/validators"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var TaskService services.TaskService = services.TaskService{}
var taskValidator validators.TaskValidator = validators.NewTaskValidator()

type TaskHandler struct{}

func (t *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	userID := r.Context().Value("userId").(int)
	log.Println("userId", userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	task, errorApp := TaskService.GetTask(taskID)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	taskJson, _ := task.MarshalJSON()
	w.Write(taskJson)
}

func (t *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, errorApp := TaskService.GetTasks()
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	var createTaskDto dtos.CreateTaskDto
	var errorApp error.Error
	err := utils.TransformBody(r.Body, &createTaskDto, &newTask)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	err = taskValidator.ValidateCreateTask(createTaskDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	newTask, errorApp = TaskService.CreateTask(newTask)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	taskJson, _ := newTask.MarshalJSON()

	w.WriteHeader(http.StatusCreated)
	w.Write(taskJson)
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	errorApp := TaskService.DeleteTask(taskID)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "The task with ID " + strconv.Itoa(taskID) + " has been remove successfully",
	})
}

func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	var updatedTask models.Task
	var updateTaskDto dtos.UpdateTaskDto
	var errorApp error.Error
	err = utils.TransformBody(r.Body, &updateTaskDto, &updatedTask)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Please enter valid data", http.StatusBadRequest, err))
		return
	}

	err = taskValidator.ValidateUpdateTask(updateTaskDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	updatedTask, errorApp = TaskService.UpdateTask(updatedTask, taskID)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "The task with ID " + strconv.Itoa(taskID) + " has been updated successfully",
	})
}

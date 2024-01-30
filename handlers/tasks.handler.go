package handlers

import (
	"app/dtos"
	"app/error"
	"app/models"
	"app/services"
	"app/validators"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var taskService services.TaskService = services.TaskService{}
var taskValidator validators.TaskValidator = validators.NewTaskValidator()

type TaskHandler struct{}

func (t *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	task, err := taskService.GetTask(taskID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error.New("Task not found", http.StatusNotFound, err))
		return
	}

	taskJson, _ := task.MarshalJSON()
	w.Write(taskJson)
}

func (t *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := taskService.GetTasks()
	if err != nil {
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(error.New("Data of tasks not found", http.StatusFound, err))
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTask models.Task
	var createTaskDto dtos.CreateTaskDto
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	json.Unmarshal(reqBody, &createTaskDto)

	err = taskValidator.ValidateCreateTask(createTaskDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	json.Unmarshal(reqBody, &newTask)

	newTask, err = taskService.CreateTask(newTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	taskJson, _ := newTask.MarshalJSON()

	w.WriteHeader(http.StatusCreated)
	w.Write(taskJson)
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	err = taskService.DeleteTask(taskID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error.New("Task not found", http.StatusNotFound, err))
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "The task with ID " + strconv.Itoa(taskID) + " has been remove successfully",
	})
}

func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask models.Task
	var updateTaskDto dtos.UpdateTaskDto

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Please enter valid data", http.StatusBadRequest, err))
		return
	}

	json.Unmarshal(reqBody, &updateTaskDto)

	err = taskValidator.ValidateUpdateTask(updateTaskDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	json.Unmarshal(reqBody, &updatedTask)
	updatedTask, err = taskService.UpdateTask(updatedTask, taskID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error.New("Task not found", http.StatusNotFound, err))
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "The task with ID " + strconv.Itoa(taskID) + " has been updated successfully",
	})
}

package handlers

import (
	"app/data"
	"app/error"
	"app/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TaskHandler struct{}

func (t *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest))
		return
	}
	for _, task := range data.Tasks {
		if task.ID == taskID {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(error.New("Task not found", http.StatusNotFound))
}

func (t *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data.Tasks)
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTask models.Task
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest))
		return
	}
	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(data.Tasks) + 1
	data.Tasks = append(data.Tasks, newTask)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest))
		return
	}
	for index, task := range data.Tasks {
		if task.ID == taskID {
			data.Tasks = append(data.Tasks[:index], data.Tasks[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "The task with ID " + strconv.Itoa(taskID) + " has been remove successfully",
			})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(error.New("Task not found", http.StatusNotFound))
}

func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask models.Task
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest))
		return
	}
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Please enter valid data", http.StatusBadRequest))
		return
	}

	json.Unmarshal(reqBody, &updatedTask)

	for index, task := range data.Tasks {
		if task.ID == taskID {
			data.Tasks = append(data.Tasks[:index], data.Tasks[index+1:]...)
			updatedTask.ID = task.ID
			data.Tasks = append(data.Tasks, updatedTask)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "The task with ID " + strconv.Itoa(taskID) + " has been updated successfully",
			})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(error.New("Task not found", http.StatusNotFound))
}

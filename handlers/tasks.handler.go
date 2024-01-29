package handlers

import (
	"app/db"
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
	var task models.Task
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest))
		return
	}

	findTask := db.DB.First(&task, taskID)
	err = findTask.Error

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error.New("Task not found", http.StatusNotFound))
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (t *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks models.ListTask
	findtasks := db.DB.Find(&tasks)
	err := findtasks.Error
	if err != nil {
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(error.New("Data of tasks not found", http.StatusFound))
		return
	}
	json.NewEncoder(w).Encode(tasks)
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
	createdTak := db.DB.Create(&newTask)
	err = createdTak.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest))
		return
	}

	findTask := db.DB.Delete(&task, taskID)
	rowsAffected := findTask.RowsAffected

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error.New("Task not found", http.StatusNotFound))
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

	query := db.DB.Where("id = ?", taskID).Updates(updatedTask)
	rowsAffected := query.RowsAffected
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error.New("Task not found", http.StatusNotFound))
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "The task with ID " + strconv.Itoa(taskID) + " has been updated successfully",
	})
}

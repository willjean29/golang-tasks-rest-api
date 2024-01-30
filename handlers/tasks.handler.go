package handlers

import (
	"app/db"
	"app/dtos"
	"app/error"
	"app/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
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
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	findTask := db.DB.First(&task, taskID)
	err = findTask.Error

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
	var tasks models.ListTask
	findtasks := db.DB.Find(&tasks)
	err := findtasks.Error
	if err != nil {
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(error.New("Data of tasks not found", http.StatusFound, err))
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())

	var newTask models.Task
	var createTaskDto dtos.CreateTaskDto
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	json.Unmarshal(reqBody, &createTaskDto)

	err = validate.Struct(createTaskDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	json.Unmarshal(reqBody, &newTask)
	createdTak := db.DB.Create(&newTask)
	err = createdTak.Error

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
	var task models.Task
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	findTask := db.DB.Delete(&task, taskID)
	rowsAffected := findTask.RowsAffected

	if rowsAffected == 0 {
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
	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask models.Task
	var updateTaskDto dtos.UpdateTaskDto

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error: ", err)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))

		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Please enter valid data", http.StatusBadRequest, err))
		return
	}

	json.Unmarshal(reqBody, &updatedTask)
	json.Unmarshal(reqBody, &updateTaskDto)
	err = validate.Struct(updateTaskDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	query := db.DB.Where("id = ?", taskID).Updates(updatedTask)
	rowsAffected := query.RowsAffected
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error.New("Task not found", http.StatusNotFound, err))
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "The task with ID " + strconv.Itoa(taskID) + " has been updated successfully",
	})
}

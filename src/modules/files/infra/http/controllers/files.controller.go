package controllers

import (
	taskUsecases "app/src/modules/tasks/app"
	taskDatasource "app/src/modules/tasks/infra/datasource"
	userDatasource "app/src/modules/users/infra/datasource"

	taskRepositories "app/src/modules/tasks/infra/repositories"
	userUsecases "app/src/modules/users/app"
	userRepositories "app/src/modules/users/infra/repositories"
	store "app/src/shared/adapters/storage"
	error "app/src/shared/errors"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var collections = []string{"tasks", "users"}
var storeAdapter store.StoreAdapter = &store.DiskAdapter{}
var taskRepository = &taskRepositories.TasksRepository{
	Datasource: &taskDatasource.GormTaskDatasource{},
}
var userRepository = &userRepositories.UsersRepository{
	Datasource: &userDatasource.UserDatasource{},
}

type FilesController struct{}

func (f *FilesController) UploadFile(w http.ResponseWriter, r *http.Request) {
	var isFoundCollection = false
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	for _, collection := range collections {
		if collection == vars["collection"] {
			isFoundCollection = true
			break
		}
	}

	if !isFoundCollection {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid collection", http.StatusBadRequest, err))
		return
	}
	filename := r.Context().Value("filename").(string)
	switch vars["collection"] {
	case "tasks":
		log.Println("Upload file for task with id:", id)
		message, errorApp := uploadTaskFile(id, filename)
		if errorApp.StatusCode != 0 {
			w.WriteHeader(errorApp.StatusCode)
			json.NewEncoder(w).Encode(errorApp)
			return
		}
		json.NewEncoder(w).Encode(message)
	case "users":
		log.Println("Upload file for user with id:", id)
		message, errorApp := uploadUserFile(id, filename)
		if errorApp.StatusCode != 0 {
			w.WriteHeader(errorApp.StatusCode)
			json.NewEncoder(w).Encode(errorApp)
			return
		}
		json.NewEncoder(w).Encode(message)
	}
}

func uploadTaskFile(taskId int, filename string) (map[string]string, error.Error) {
	getTaskUseCase := taskUsecases.GetTaskUseCase{
		TaskRepository: taskRepository,
	}
	task, errorApp := getTaskUseCase.Execute(taskId)
	if errorApp.StatusCode != 0 {
		return map[string]string{}, errorApp
	}
	if task.Image != "" {
		storeAdapter.DeleteFile(task.Image)
	}

	task.Image, _ = storeAdapter.SaveFile(filename)
	saveTask := taskUsecases.SaveTaskUseCase{
		TaskRepository: taskRepository,
	}
	errorApp = saveTask.Execute(task)
	if errorApp.StatusCode != 0 {

	}
	return map[string]string{"message": "File uploaded successfully"}, error.Error{}
}

func uploadUserFile(userId int, filename string) (map[string]string, error.Error) {

	getUserUseCase := userUsecases.GetUserUseCase{
		UserRepository: userRepository,
	}
	user, errorApp := getUserUseCase.Execute(userId)
	if errorApp.StatusCode != 0 {
		return map[string]string{}, errorApp
	}
	if user.Image != "" {
		storeAdapter.DeleteFile(user.Image)
	}
	user.Image, _ = storeAdapter.SaveFile(filename)
	saveUserUseCase := userUsecases.SaveUserUseCase{
		UserRepository: userRepository,
	}
	errorApp = saveUserUseCase.Execute(user)
	if errorApp.StatusCode != 0 {
		return map[string]string{}, errorApp
	}
	return map[string]string{"message": "File uploaded successfully"}, error.Error{}
}

package routes

import (
	"app/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router) {
	var taskHandler handlers.TaskHandler = handlers.TaskHandler{}
	taskRouter := router.PathPrefix("/tasks").Subrouter()
	taskRouter.HandleFunc("", taskHandler.GetTasks).Methods(http.MethodGet)
	taskRouter.HandleFunc("", taskHandler.CreateTask).Methods(http.MethodPost)
	taskRouter.HandleFunc("/{id}", taskHandler.GetTask).Methods(http.MethodGet)
	taskRouter.HandleFunc("/{id}", taskHandler.DeleteTask).Methods(http.MethodDelete)
	taskRouter.HandleFunc("/{id}", taskHandler.UpdateTask).Methods(http.MethodPut)
}

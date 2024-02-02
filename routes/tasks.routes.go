package routes

import (
	"app/handlers"
	"app/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router) {
	var taskHandler handlers.TaskHandler = handlers.TaskHandler{}
	taskRouter := router.PathPrefix("/tasks").Subrouter()
	taskRouter.HandleFunc("", taskHandler.GetTasks).Methods(http.MethodGet)
	taskRouter.HandleFunc("", taskHandler.CreateTask).Methods(http.MethodPost)
	taskRouter.Handle("/{id}", middlewares.Authenticated(http.HandlerFunc(taskHandler.GetTask))).Methods(http.MethodGet)
	taskRouter.HandleFunc("/{id}", taskHandler.DeleteTask).Methods(http.MethodDelete)
	taskRouter.HandleFunc("/{id}", taskHandler.UpdateTask).Methods(http.MethodPut)
}

package routes

import (
	"app/src/modules/tasks/infra/http/controllers"
	"app/src/shared/infra/http/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router) {
	var tasksController = controllers.TaskController{}
	taskRouter := router.PathPrefix("/tasks").Subrouter()
	taskRouter.Use(middlewares.Authenticated)
	taskRouter.HandleFunc("", tasksController.List).Methods(http.MethodGet)
	taskRouter.HandleFunc("", tasksController.Create).Methods(http.MethodPost)
	taskRouter.HandleFunc("/{id}", tasksController.Show).Methods(http.MethodGet)
	taskRouter.HandleFunc("/{id}", tasksController.Update).Methods(http.MethodPut)
	taskRouter.HandleFunc("/{id}", tasksController.Delete).Methods(http.MethodDelete)
}

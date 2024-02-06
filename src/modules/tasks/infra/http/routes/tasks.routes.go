package routes

import (
	"app/src/modules/tasks/infra/http/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router) {
	var tasksController = controllers.TasksController{}
	taskRouter := router.PathPrefix("/tareas").Subrouter()

	taskRouter.HandleFunc("", tasksController.List).Methods(http.MethodGet)
}

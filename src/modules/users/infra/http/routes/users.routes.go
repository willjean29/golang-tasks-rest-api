package routes

import (
	"app/src/modules/users/infra/http/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	var usersController = controllers.UsersController{}
	userRouter := router.PathPrefix("/usuarios").Subrouter()

	userRouter.HandleFunc("", usersController.List).Methods(http.MethodGet)
	// taskRouter.HandleFunc("", tasksController.Create).Methods(http.MethodPost)
	userRouter.HandleFunc("/{id}", usersController.Show).Methods(http.MethodGet)
	// taskRouter.HandleFunc("/{id}", tasksController.Update).Methods(http.MethodPut)
	// taskRouter.HandleFunc("/{id}", tasksController.Delete).Methods(http.MethodDelete)
	userRouter.HandleFunc("/login", usersController.Login).Methods(http.MethodPost)
	userRouter.HandleFunc("/register", usersController.Register).Methods(http.MethodPost)
}

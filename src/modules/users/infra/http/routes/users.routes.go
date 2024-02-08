package routes

import (
	"app/src/modules/users/infra/http/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	var usersController = controllers.UsersController{}
	authRouter := router.PathPrefix("/auth").Subrouter()
	userRouter := router.PathPrefix("/users").Subrouter()

	authRouter.HandleFunc("/login", usersController.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/register", usersController.Register).Methods(http.MethodPost)

	userRouter.HandleFunc("", usersController.List).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", usersController.Show).Methods(http.MethodGet)

}

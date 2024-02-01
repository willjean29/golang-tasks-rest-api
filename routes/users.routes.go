package routes

import (
	"app/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	var userHandler handlers.UserHandler = handlers.UserHandler{}
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", userHandler.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)
}

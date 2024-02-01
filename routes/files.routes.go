package routes

import (
	"app/handlers"
	"app/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func FilesRoutes(router *mux.Router) {
	var filesHandler handlers.FilesHandler = handlers.FilesHandler{}
	filesRouter := router.PathPrefix("/files").Subrouter()
	filesRouter.Handle("/{collection}/{id}", middlewares.UploadFile(http.HandlerFunc(filesHandler.UploadFile))).Methods(http.MethodPost)
}

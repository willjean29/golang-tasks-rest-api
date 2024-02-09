package routes

import (
	"app/src/modules/files/infra/http/controllers"
	"app/src/shared/infra/http/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func FileRoutes(router *mux.Router) {
	var filesController = controllers.FilesController{}
	filesRouter := router.PathPrefix("/files").Subrouter()
	filesRouter.Use(middlewares.Authenticated)
	filesRouter.Handle("/{collection}/{id}", middlewares.UploadFile(http.HandlerFunc(filesController.UploadFile))).Methods(http.MethodPost)
}

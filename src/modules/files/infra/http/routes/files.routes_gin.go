package routes

import (
	"app/src/modules/files/infra/http/controllers"
	"app/src/shared/infra/http/middlewares"

	"github.com/gin-gonic/gin"
)

func FileRoutesGin(router *gin.RouterGroup) {
	var filesController = controllers.FileControllerGin{}
	filesRoutes := router.Group("/files")
	filesRoutes.Use(middlewares.AuthenticatedGin())

	filesRoutes.POST("/:collection/:id", middlewares.UploadFileGin(), filesController.UploadFile)
}

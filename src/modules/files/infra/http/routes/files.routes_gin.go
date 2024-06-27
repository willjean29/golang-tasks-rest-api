package routes

import (
	"app/src/shared/infra/http/middlewares"

	"github.com/gin-gonic/gin"
)

func FileRoutesGin(router *gin.RouterGroup) {
	filesRoutes := router.Group("/files")
	filesRoutes.Use(middlewares.AuthenticatedGin())

	filesRoutes.POST("/:collection/:id", middlewares.UploadFileGin(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "File uploaded",
		})
	})
}

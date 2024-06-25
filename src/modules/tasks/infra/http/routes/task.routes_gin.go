package routes

import (
	"app/src/modules/tasks/infra/http/controllers"
	"app/src/shared/infra/http/middlewares"

	"github.com/gin-gonic/gin"
)

func TaskRoutesGin(router *gin.RouterGroup) {
	var tasksController = controllers.TaskControllerGin{}
	taskRoutes := router.Group("/tasks")
	taskRoutes.Use(middlewares.AuthenticatedGin())
	taskRoutes.GET("/", tasksController.List)

	taskRoutes.GET("/:id", tasksController.Show)

	taskRoutes.POST("/", tasksController.Create)

	taskRoutes.PUT("/:id", tasksController.Update)

	taskRoutes.DELETE("/:id", tasksController.Delete)
}

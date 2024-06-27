package routes

import (
	fileRoutes "app/src/modules/files/infra/http/routes"
	taskRoutes "app/src/modules/tasks/infra/http/routes"
	userRoutes "app/src/modules/users/infra/http/routes"

	"github.com/gin-gonic/gin"
)

func RoutesGin(router *gin.Engine) {
	apiRouter := router.Group("/api")
	userRoutes.UserRoutesGin(apiRouter)
	taskRoutes.TaskRoutesGin(apiRouter)
	fileRoutes.FileRoutesGin(apiRouter)
}

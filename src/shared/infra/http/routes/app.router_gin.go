package routes

import (
	userRoutes "app/src/modules/users/infra/http/routes"

	"github.com/gin-gonic/gin"
)

func RoutesGin(router *gin.Engine) {
	apiRouter := router.Group("/api")
	userRoutes.UserRoutesGin(apiRouter)
}

package routes

import (
	"app/src/modules/users/infra/http/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutesGin(router *gin.RouterGroup) {
	var usersControllerGin = controllers.UserControllerGin{}
	userRouter := router.Group("/users")
	authRouter := router.Group("/auth")

	userRouter.GET("/", usersControllerGin.List)
	userRouter.GET("/:id", usersControllerGin.Show)

	authRouter.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login user",
		})
	})

	authRouter.POST("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "register user",
		})
	})
}

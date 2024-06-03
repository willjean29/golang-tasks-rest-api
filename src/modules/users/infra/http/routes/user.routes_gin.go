package routes

import "github.com/gin-gonic/gin"

func UserRoutesGin(router *gin.RouterGroup) {
	userRouter := router.Group("/users")
	authRouter := router.Group("/auth")
	userRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get all users",
		})
	})
	userRouter.GET("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get user by id",
		})
	})

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

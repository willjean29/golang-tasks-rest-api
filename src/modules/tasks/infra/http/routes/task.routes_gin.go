package routes

import "github.com/gin-gonic/gin"

func TaskRoutesGin(router *gin.RouterGroup) {
	taskRoutes := router.Group("/tasks")

	taskRoutes.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "List all tasks",
		})
	})

	taskRoutes.GET("/:id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Show a task",
		})
	})

	taskRoutes.POST("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Create a task",
		})
	})

	taskRoutes.PUT("/:id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Update a task",
		})
	})

	taskRoutes.DELETE("/:id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Delete a task",
		})
	})
}

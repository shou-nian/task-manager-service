package routes

import (
	"github.com/gin-gonic/gin"
	"task-manager-service/controllers"
)

func Routes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		// create task
		v1.POST("/tasks", controllers.CreateTask)

		// get all tasks
		v1.GET("/tasks", nil)

		// get target task
		v1.GET("/tasks/:id", func(context *gin.Context) {
			_ = context.Param("id")
		})

		// update task
		v1.PUT("/tasks/:id", func(context *gin.Context) {
			_ = context.Param("id")
		})

		// delete task
		v1.DELETE("/tasks/:id", func(context *gin.Context) {
			_ = context.Param("id")
		})
	}

	return router
}

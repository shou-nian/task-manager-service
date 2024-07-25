package routes

import (
	"github.com/gin-gonic/gin"
	"task-manager-service/controllers"
)

func SetupRoutes(router *gin.Engine) {
	//router := gin.New()

	v1 := router.Group("/api/v1")
	{
		// create task
		v1.POST("/tasks", controllers.CreateTask)

		// get all tasks
		v1.GET("/tasks", controllers.GetAllTasks)

		// get target task
		v1.GET("/tasks/:id", controllers.GetTargetTask)

		// update task
		v1.PUT("/tasks/:id", controllers.UpdateTask)

		// delete task
		v1.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	//return router
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task-manager-service/models"
)

func CreateTask(ctx *gin.Context) {
	var json models.CreateTask

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if len(json.Title) > 256 || json.Title == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "Title must be 1 and 256 characters."})
		return
	}

	if len(json.Description) > 256 {
		ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "Description max be 256 characters."})
		return
	}

	ctx.JSONP(http.StatusOK, map[string]string{"msg": "Task created successfully."})
}

func GetAllTasks(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK,
		map[string]interface{}{
			"data": map[string]interface{}{
				"created": []string{"task1", "task2"},
				"updated": []string{"task3"},
				"deleted": []string{"task4"},
			},
		})
}

func GetTargetTask(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, map[string]interface{}{"data": id})
}

func UpdateTask(ctx *gin.Context) {

}

func DeleteTask(ctx *gin.Context) {

}

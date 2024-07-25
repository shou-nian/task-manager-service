package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"task-manager-service/middlewares"
	"task-manager-service/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery(), middlewares.Logger(), middlewares.AuthRequired())
	routes.SetupRoutes(router)

	log.Fatal(router.Run("127.0.0.1:8080"))
}

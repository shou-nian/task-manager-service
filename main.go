package main

import (
	"log"
	"task-manager-service/routes"
)

func main() {
	router := routes.Routes()

	log.Fatal(router.Run("127.0.0.1:8080"))
}

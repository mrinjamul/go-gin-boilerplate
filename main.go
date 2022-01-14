package main

import (
	"go-gin-boilerplate/api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	server := gin.Default()
	// Initialize the routes
	routes.InitRoutes(server)
	// Start and run the server
	log.Fatal(server.Run(":8080"))
}

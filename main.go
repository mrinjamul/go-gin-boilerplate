package main

import (
	"go-gin-boilerplate/api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get port from env
	port := ":8080"
	_, present := os.LookupEnv("PORT")
	if present {
		port = ":" + os.Getenv("PORT")

	}
	// Set the router as the default one shipped with Gin
	server := gin.Default()
	// Initialize the routes
	routes.InitRoutes(server)
	// Start and run the server
	log.Fatal(server.Run(port))
}

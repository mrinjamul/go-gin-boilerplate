package main

import (
	"embed"
	"go-gin-boilerplate/api/routes"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// @title go gin API
// @version 1.0
// @description This is a gin API for go
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email mrinjamul@gmail.com
// @license.name MIT License
// @license.url https://github.com/mrinjamul/go-gin-boilerplate/blob/main/LICENSE
// @BasePath /
// @schemes http https
// @securitydefinitions.apikey	APIKeyAuth
// @in header
// @name Authorization

var (
	startTime time.Time = time.Now()
)

//go:embed templates/*
var webpages embed.FS

func main() {
	// Get port from env
	port := ":3000"
	_, present := os.LookupEnv("PORT")
	if present {
		port = ":" + os.Getenv("PORT")

	}
	// Set the router as the default one shipped with Gin
	server := gin.Default()
	templ := template.Must(template.New("").ParseFS(webpages, "templates/layouts/*.html"))
	server.SetHTMLTemplate(templ)
	static, err := fs.Sub(webpages, "templates/static")
	if err != nil {
		panic(err)
	}
	server.StaticFS("/static", http.FS(static))

	// Initialize the routes
	routes.StartTime = startTime
	routes.InitRoutes(server)
	routes.BootTime = time.Since(startTime)

	// Start and run the server
	log.Fatal(server.Run(port))
}

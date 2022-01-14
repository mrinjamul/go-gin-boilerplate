package routes

import (
	"go-gin-boilerplate/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(routes *gin.Engine) {
	svc := services.NewServices()
	// Serve the frontend
	routes.LoadHTMLGlob("templates/*")
	routes.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})
	// health check
	routes.GET("/api/health", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"health": "ok",
			},
		)
	})
	// quote routes
	routes.GET("/quotes", func(c *gin.Context) {
		svc.QuoteService().GetAll(c)
	})
	routes.GET("/quotes/:id", func(c *gin.Context) {
		svc.QuoteService().Get(c)
	})
	routes.POST("/quotes", func(c *gin.Context) {
		svc.QuoteService().Add(c)
	})
	routes.PUT("/quotes/:id", func(c *gin.Context) {
		svc.QuoteService().Update(c)
	})
	routes.DELETE("/quotes/:id", func(c *gin.Context) {
		svc.QuoteService().Delete(c)
	})
}

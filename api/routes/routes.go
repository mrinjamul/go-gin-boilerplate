package routes

import (
	"go-gin-boilerplate/api/services"
	"time"

	docs "go-gin-boilerplate/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	// StartTime is the time when the server started
	StartTime time.Time
	// BootTime is the time when the server booted
	BootTime time.Duration
)

func InitRoutes(router *gin.Engine) {
	// Initialize services
	svc := services.NewServices()
	// Add CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Initialize the routes

	// Home Page
	router.GET("/", func(c *gin.Context) {
		svc.View().Index(c)
	})

	// 404 Page
	router.NoRoute(func(ctx *gin.Context) {
		svc.View().NotFound(ctx)
	})

	// Backend API
	docs.SwaggerInfo.BasePath = "/"
	// health check
	router.GET("/api/health", func(c *gin.Context) {
		svc.HealthCheckService().HealthCheck(c, StartTime, BootTime)
	})
	v1 := router.Group("/api/v1")
	{
		// quote routes
		v1.GET("/quote", func(c *gin.Context) {
			svc.QuoteService().GetAll(c)
		})
		v1.GET("/quote/:id", func(c *gin.Context) {
			svc.QuoteService().Get(c)
		})
		v1.POST("/quote", func(c *gin.Context) {
			svc.QuoteService().Add(c)
		})
		v1.PUT("/quote/:id", func(c *gin.Context) {
			svc.QuoteService().Update(c)
		})
		v1.DELETE("/quote/:id", func(c *gin.Context) {
			svc.QuoteService().Delete(c)
		})
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

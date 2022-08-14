package services

import (
	"go-gin-boilerplate/api/controllers"
	"go-gin-boilerplate/database"
	"go-gin-boilerplate/repository"
)

type Services interface {
	HealthCheckService() controllers.HealthCheck
	QuoteService() controllers.Quote
	View() controllers.Template
}

type services struct {
	healthCheck controllers.HealthCheck
	quote       controllers.Quote
	view        controllers.Template
}

func (svc *services) HealthCheckService() controllers.HealthCheck {
	return svc.healthCheck
}

func (svc *services) QuoteService() controllers.Quote {
	return svc.quote
}

func (svc *services) View() controllers.Template {
	return svc.view
}

// NewServices initializes services
func NewServices() Services {
	db := database.GetDB()
	return &services{
		healthCheck: controllers.NewHealthCheck(),
		quote: controllers.NewQuote(
			repository.NewQuoteRepo(db),
		),
		view: controllers.NewTemplate(),
	}
}

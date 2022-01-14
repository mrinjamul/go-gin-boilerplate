package services

import (
	"go-gin-boilerplate/api/controllers"
	"go-gin-boilerplate/db"
	"go-gin-boilerplate/repository"
)

type Services interface {
	QuoteService() controllers.Quote
}

type services struct {
	quote controllers.Quote
}

func (svc *services) QuoteService() controllers.Quote {
	return svc.quote
}

// NewServices initializes services
func NewServices() Services {
	db := db.GetDB()
	return &services{
		quote: controllers.NewQuote(
			repository.NewQuoteRepo(db),
		),
	}
}

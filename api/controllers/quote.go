package controllers

import (
	"encoding/json"
	"go-gin-boilerplate/models"
	"go-gin-boilerplate/repository"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Quote interface {
	Add(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type quote struct {
	quoteRepo repository.QuoteRepo
}

// Add a new quote
func (q *quote) Add(ctx *gin.Context) {
	bytes, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Fatal(err)
	}
	var quote models.Quote
	err = json.Unmarshal(bytes, &quote)
	if err != nil {
		log.Fatal(err)
	}
	q.quoteRepo.Add(ctx, &quote)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"article": quote,
	})
}

// Get a quote by id
func (q *quote) Get(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	quote := models.Quote{
		ID: uint(id),
	}

	quote, err = q.quoteRepo.Get(ctx, quote)

	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Success",
			"article": quote,
		})
	} else {

		ctx.JSON(200, gin.H{
			"message": "Fail",
			"article": "",
		})
	}
}

// GetAll all quotes
func (q *quote) GetAll(ctx *gin.Context) {
	quotes, _ := q.quoteRepo.GetAll(ctx)

	ctx.JSON(200, gin.H{
		"message":  "Success",
		"articles": quotes,
	})
}

// Update a quote
func (q *quote) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Fatal(err)
	}
	var quote models.Quote
	err = json.Unmarshal(bytes, &quote)
	if err != nil {
		log.Fatal(err)
	}
	quote.ID = uint(id)
	q.quoteRepo.Update(ctx, &quote)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"article": quote,
	})
}

// Delete a quote
func (q *quote) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	quote := models.Quote{
		ID: uint(id),
	}

	q.quoteRepo.Delete(ctx, &quote)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"article": quote,
	})
}

// NewQuote returns a new quote controller
func NewQuote(
	quoteRepo repository.QuoteRepo,
) Quote {
	return &quote{
		quoteRepo: quoteRepo,
	}
}

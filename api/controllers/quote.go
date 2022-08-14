package controllers

import (
	"encoding/json"
	"go-gin-boilerplate/models"
	"go-gin-boilerplate/repository"
	"log"
	"net/http"
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

// Add godoc
// @Summary Add a quote
// @Description Add a quote
// @ID add-quote
// @Tags quote
// @Accept  json
// @Produce  json
// @Param quote body models.Quote true "Quote"
// @Success 200 {object} models.Quote
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /quote [post]
func (q *quote) Add(ctx *gin.Context) {
	var quote models.Quote
	err := json.NewDecoder(ctx.Request.Body).Decode(&quote)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Error: models.ServiceError{
				Kind:    "BadRequest",
				Code:    "BadRequest",
				Message: "Invalid JSON",
			},
		})
		return
	}
	q.quoteRepo.Add(ctx, &quote)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"article": quote,
	})
}

// Get godoc
// @Summary Get a quote
// @Description Get a quote
// @ID get-quote
// @Tags quote
// @Accept  json
// @Produce  json
// @Param id path int true "Quote ID"
// @Success 200 {object} models.Quote
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /quote/{id} [get]
func (q *quote) Get(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	quote := models.Quote{
		Base: models.Base{
			ID: uint(id),
		},
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

// GetAll godoc
// @Summary Get all quotes
// @Description Get all quotes
// @ID get-all-quotes
// @Tags quote
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Quote
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /quote [get]
func (q *quote) GetAll(ctx *gin.Context) {
	quotes, _ := q.quoteRepo.GetAll(ctx)

	ctx.JSON(200, gin.H{
		"message":  "Success",
		"articles": quotes,
	})
}

// Update godoc
// @Summary Update a quote
// @Description Update a quote
// @ID update-quote
// @Tags quote
// @Accept  json
// @Produce  json
// @Param id path int true "Quote ID"
// @Param quote body models.Quote true "Quote"
// @Success 200 {object} models.Quote
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /quote/{id} [put]
func (q *quote) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	var quote models.Quote
	err = json.NewDecoder(ctx.Request.Body).Decode(&quote)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Error: models.ServiceError{
				Kind:    "BadRequest",
				Code:    "BadRequest",
				Message: "Invalid JSON",
			},
		})
		return
	}

	quote.ID = uint(id)
	q.quoteRepo.Update(ctx, &quote)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"article": quote,
	})
}

// Delete godoc
// @Summary Delete a quote
// @Description Delete a quote
// @ID delete-quote
// @Tags quote
// @Accept  json
// @Produce  json
// @Param id path int true "Quote ID"
// @Success 200 {object} models.Quote
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
func (q *quote) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	quote := models.Quote{
		Base: models.Base{
			ID: uint(id),
		},
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

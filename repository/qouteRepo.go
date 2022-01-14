package repository

import (
	"go-gin-boilerplate/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuoteRepo interface {
	Add(ctx *gin.Context, qoute *models.Quote) error
	Get(ctx *gin.Context, quote models.Quote) (models.Quote, error)
	GetAll(ctx *gin.Context) ([]models.Quote, error)
	Update(ctx *gin.Context, quote *models.Quote) error
	Delete(ctx *gin.Context, quote *models.Quote) error
}

type quoteRepo struct {
	db gorm.DB
}

// Add a new quote
func (repo *quoteRepo) Add(ctx *gin.Context, qoute *models.Quote) error {
	result := repo.db.Create(&qoute)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get a quote by id
func (repo *quoteRepo) Get(ctx *gin.Context, quote models.Quote) (models.Quote, error) {
	result := repo.db.First(&quote)
	if result.Error != nil {
		return quote, result.Error
	}
	return quote, nil
}

// GetAll all quotes
func (repo *quoteRepo) GetAll(ctx *gin.Context) ([]models.Quote, error) {
	var quotes []models.Quote
	result := repo.db.Find(&quotes)
	if result.Error != nil {
		return quotes, result.Error
	}
	return quotes, nil
}

// Update a quote
func (repo *quoteRepo) Update(ctx *gin.Context, quote *models.Quote) error {
	result := repo.db.Save(&quote)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete a quote
func (repo *quoteRepo) Delete(ctx *gin.Context, quote *models.Quote) error {
	result := repo.db.Delete(&quote)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewQuoteRepo(db *gorm.DB) QuoteRepo {
	return &quoteRepo{
		db: *db,
	}
}

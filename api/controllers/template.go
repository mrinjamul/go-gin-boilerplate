package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Template is a struct for go html/template
type Template interface {
	Index(c *gin.Context)
	NotFound(c *gin.Context)
}

type template struct {
}

// Index is a function for index page
func (t *template) Index(c *gin.Context) {
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
}

// NotFound is a function for not found page
func (t *template) NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"title":     "Not Found",
		"copyright": "Copyright © 2022 mrinjamul. All rights reserved.",
	})
}

// NewTemplate is a function for new template
func NewTemplate() Template {
	return &template{}
}

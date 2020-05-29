package controllers

import (
	"github.com/daisuzuki829/bookshelf/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBooks ...
func (h *Handler) GetAllBooks(c *gin.Context) {
	r := models.NewBookRepository()
	books := r.GetAll()

	c.HTML(http.StatusOK, "books.html", gin.H{
		"books": books,
	})
}
package controllers

import (
	"github.com/daisuzuki829/bookshelf/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllBooks ...
func (h *Handler) GetAllBooks(c *gin.Context) {
	r := models.NewBookRepository()
	books := r.GetAll()
	c.HTML(http.StatusOK, "books.html", gin.H{
		"books": books,
	})
}

// AddBooks ...
func (h *Handler) AddBook(c *gin.Context) {
	r := models.NewBookRepository()
	title, _     := c.GetPostForm("title")
	authors, _   := c.GetPostForm("authors")
	publisher, _ := c.GetPostForm("publisher")
	r.Add(&models.Book{Title: title, Authors: authors, Publisher:publisher})
	c.Redirect(http.StatusMovedPermanently, "/books")
}

// GetBooks ...
func (h *Handler) GetBook(c *gin.Context) {
	r := models.NewBookRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	book := r.GetOne(id)
	c.HTML(http.StatusOK, "book_edit.html", gin.H{
		"book": book,
	})
}

// EditBooks ...
func (h *Handler) EditBook(c *gin.Context) {
	r := models.NewBookRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	book := r.GetOne(id)
	book.Title, _     = c.GetPostForm("title")
	book.Authors, _   = c.GetPostForm("authors")
	book.Publisher, _ = c.GetPostForm("publisher")
	r.Edit(book)
	c.Redirect(http.StatusMovedPermanently, "/books")
}

// DeleteBooks ...
func (h *Handler) DeleteBook(c *gin.Context) {
	r := models.NewBookRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/books")
}


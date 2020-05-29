package controllers

import (
	"github.com/daisuzuki829/bookshelf/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBooks ...
func (h *Handler) GetAllUsers(c *gin.Context) {
	r := models.NewUserRepository()
	users := r.GetAll()

	c.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}
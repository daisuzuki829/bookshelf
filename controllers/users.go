package controllers

import (
	"fmt"
	"github.com/daisuzuki829/bookshelf/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUsers ...
func (h *Handler) GetAllUsers(c *gin.Context) {
	r := models.NewUserRepository()
	users := r.GetAll()

	c.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}

// AddUsers ...
func (h *Handler) AddUser(c *gin.Context) {
	r := models.NewUserRepository()
	nickname, _ := c.GetPostForm("nickname")
	password, _ := c.GetPostForm("password")
	_age, _     := c.GetPostForm("age")
	age, _      := strconv.Atoi(_age)
	role, _     := c.GetPostForm("role")

	r.Add(&models.User{Nickname: nickname, Password: password, Age: age, Role: role})

	fmt.Printf("user : ")
	spew.Dump(&models.User{Nickname: nickname, Password: password, Age: age, Role: role})
	fmt.Printf("\n")

	c.Redirect(http.StatusMovedPermanently, "/users")
}

// GetUsers ...
func (h *Handler) GetUser(c *gin.Context) {
	r := models.NewUserRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	user := r.GetOne(id)
	c.HTML(http.StatusOK, "user_edit.html", gin.H{
		"user": user,
	})
}

// EditUsers ...
func (h *Handler) EditUser(c *gin.Context) {
	r := models.NewUserRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	user := r.GetOne(id)
	user.Nickname, _ = c.GetPostForm("nickname")
	user.Password, _ = c.GetPostForm("password")
	age, _           := c.GetPostForm("age")
	user.Age, _      = strconv.Atoi(age)
	user.Role, _     = c.GetPostForm("role")
	r.Edit(user)
	c.Redirect(http.StatusMovedPermanently, "/users")
}

// DeleteUsers ...
func (h *Handler) DeleteUser(c *gin.Context) {
	r := models.NewUserRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/users")
}


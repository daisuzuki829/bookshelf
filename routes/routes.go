package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Hello...
func Hello(c *gin.Context) {
	name := c.DefaultQuery("name", "Michael")
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"name": name,
	})
}

// HelloName...
func HelloName(c *gin.Context) {
	name := c.Param("name")
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"name": name,
	})
}

// NoRoute...
func NoRoute(c *gin.Context) {
	// helloに飛ばす
	c.Redirect(http.StatusMovedPermanently, "/index")
}

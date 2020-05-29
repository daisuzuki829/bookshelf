package routes

import (
	"github.com/daisuzuki829/bookshelf/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Handler(dbConn *gorm.DB) {

	handler := controllers.Handler{
		Db: dbConn,
	}

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")

	r.GET("/books", handler.GetAllBooks)
	r.GET("/users", handler.GetAllUsers)

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "TITLE",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}


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
	rBook := r.Group("/book")
	{
		rBook.POST("/add", handler.AddBook)
		rBook.GET("/edit/:id", handler.GetBook)
		rBook.POST("/edit_ok/:id", handler.EditBook)
		rBook.GET("/delete/:id", handler.DeleteBook)
		rBook.DELETE("/delete/:id", handler.DeleteBook)
	}

	r.GET("/users", handler.GetAllUsers)
	rUser := r.Group("/user")
	{
		rUser.POST("/add", handler.AddUser)
		rUser.GET("/edit/:id", handler.GetUser)
		rUser.POST("/edit_ok/:id", handler.EditUser)
		rUser.GET("/delete/:id", handler.DeleteUser)
		rUser.DELETE("/delete/:id", handler.DeleteUser)
	}

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "TITLE",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	//spew.Dump(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}


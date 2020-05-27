package main

import (
	"github.com/daisuzuki829/bookshelf/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// テンプレートをロード
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")

	// ハンドラの指定
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "TITLE",
		})
	})


	// どのルーティングにも当てはまらなかった場合
	r.NoRoute(routes.NoRoute)
	r.Run(":8080")
}

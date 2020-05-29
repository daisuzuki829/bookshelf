package main

import (
	"github.com/daisuzuki829/bookshelf/db"
	"github.com/daisuzuki829/bookshelf/routes"
)

func main() {
	dbConn := db.Init()
	routes.Handler(dbConn)
}
package main

import (
	"github.com/daisuzuki829/bookshelf/db"
	"github.com/daisuzuki829/bookshelf/routes"
)

func main() {
	routes.Handler(db.Init())
}
package db

import (
	"github.com/daisuzuki829/bookshelf/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// dbInit...
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "bookshelf.db")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.AutoMigrate(&models.User{}, &models.Book{}, &models.RentalHistory{})
	//.AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("book_id", "books(id)", "RESTRICT", "RESTRICT")

	defer db.Close()
	return db
}

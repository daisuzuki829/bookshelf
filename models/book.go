package models

import (
	"github.com/golang-sql/civil"
	"github.com/jinzhu/gorm"
)

const (
	dbName = "bookshelf.db"
	dialect = "sqlite3"
)

type Book struct {
	gorm.Model
	Title        string  `gorm:"not null"`
	Authors      string  `gorm:"not null"`
	Publisher    string  `gorm:"not null"`
	PublishYear  civil.Date
	IgnoreMe     string  `gorm:"-"`
}

// NewBook ...
func NewBook(name string) Book {
	return Book{
		Title: name,
	}
}

// BookRepository is
type BookRepository struct {
}

// NewBookRepository ...
func NewBookRepository() BookRepository {
	return BookRepository{}
}


//DB追加
func (r *BookRepository) Add(book *Book) {
	db, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic(err)
	}
	db.Create(book)
	defer db.Close()
}

//DB更新
func (r *BookRepository) Edit(book Book) {
	db, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic(err)
	}
	db.Save(book)
	db.Close()
}

//DB全取得
func (r *BookRepository) GetAll() []Book {
	db, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic(err)
	}
	var books []Book
	db.Find(&books)
	db.Close()
	return books
}

//DB一つ取得
func (r *BookRepository) GetOne(id int) Book {
	db, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic(err)
	}
	var book Book
	db.First(&book, id)
	db.Close()
	return book
}

//DB削除
func (r *BookRepository) Delete(id int) {
	db, err := gorm.Open(dialect, dbName)
	if err != nil {
		panic(err)
	}
	var book Book
	db.First(&book, id)
	db.Delete(&book)
	db.Close()
}


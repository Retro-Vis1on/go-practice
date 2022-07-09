package models

import (
	"github.com/Retro-Vis1on/go-practice/bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
func GetBookById(id int64) (*Book, *gorm.DB) {
	var curBook Book
	db := db.Where("ID=?", id).Find(&curBook)
	return &curBook, db
}
func DeleteBook(id int64) Book {
	var curBook Book
	db.Where("ID=?", id).Delete(&curBook)
	return curBook
}

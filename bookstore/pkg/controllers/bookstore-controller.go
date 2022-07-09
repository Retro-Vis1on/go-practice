package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Retro-Vis1on/go-practice/bookstore/pkg/models"
	"github.com/Retro-Vis1on/go-practice/bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		return
	}
	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Print("Error in parsing")
	}
	DeletedBook := models.DeleteBook(id)
	res, _ := json.Marshal(DeletedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Print("Error in parsing")
	}
	bookDetails, db := models.GetBookById(id)
	if bookDetails.Name != "" {
		bookDetails.Name = updatedBook.Name
	}
	if bookDetails.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if bookDetails.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

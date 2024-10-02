package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tiltoin123/go-bookstore/pkg/models"
	"github.com/tiltoin123/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{} // Create an instance of Book
    newBooks, err := book.GetAllBooks()
    if newBooks == nil {
        http.Error(w, "Error retrieving books: ", http.StatusInternalServerError)
        return
    }

    res, err := json.Marshal(newBooks)
    if err != nil {
        http.Error(w, "Error marshaling books: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetBookById(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error while parsing",err)
	}
	book := &models.Book{}
	bookDetails, _:=book.GetBookById(ID)
	if bookDetails==nil{
		http.Error(w, "Error getting book: ", http.StatusInternalServerError)
        return
	}
	res, err := json.Marshal(bookDetails)
	if err != nil {
        http.Error(w, "Error marshaling book details: "+err.Error(), http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	book := &models.Book{}
	book.Name=vars["name"]
	book.Author=vars["author"]
	book.Publication=vars["publication"]
	CreateBook := &models.Book{}
	utils.ParseBody(r,CreateBook)
	b,err:=book.CreateBook()
	res, err := json.Marshal(b)
	if err != nil {
        http.Error(w, "Error marshaling created book: "+err.Error(), http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars :=mux.Vars(r)
	bookId := vars["bookId"]
	book := &models.Book{}
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error while parsing",err)
	}
	deletedBook,err:=book.DeleteBook(ID)
		if deletedBook==nil{
		http.Error(w, "Error deleting book: ", http.StatusInternalServerError)
        return
	}
	res, err := json.Marshal(deletedBook)
	if err != nil {
        http.Error(w, "Error marshaling deleted book: "+err.Error(), http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	bookId:= vars["id"]
	book := &models.Book{}
	book.Name=vars["name"]
	book.Author=vars["author"]
	book.Publication=vars["publication"]

    // Parse the request body into the book instance
    if err := utils.ParseBody(r, book); err != nil {
        http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Parse the book ID
    ID, err := strconv.ParseInt(bookId, 10, 64) // Use 10 for base and 64 for bit size
    if err != nil {
        http.Error(w, "Error while parsing book ID: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Call the method to update the book
    updatedBook, err := models.UpdateBook(ID, book)
    if err != nil {
        http.Error(w, "Error updating book: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Marshal the updated book into JSON
    res, err := json.Marshal(updatedBook)
    if err != nil {
        http.Error(w, "Error marshaling updated book: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Set response headers and status
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

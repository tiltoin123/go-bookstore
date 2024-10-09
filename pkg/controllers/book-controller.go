package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tiltoin123/go-bookstore/pkg/models"
	"github.com/tiltoin123/go-bookstore/pkg/utils"
)

func GetBooks(w http.ResponseWriter,r *http.Request) {
	book := &models.Book{}
	newBooks, err := book.GetAllBooks()
	if err != nil {
		fmt.Println("Error getting books", err)
	}
	if newBooks == nil {
		http.Error(w, "Error retrieving books: ", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "Error marshaling books: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing", err)
	}
	book := &models.Book{}
	bookDetails, _ := book.GetBookById(ID)
	if bookDetails == nil {
		http.Error(w, "Error getting book: ", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Error marshaling book details: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	err := utils.ParseBody(r, book)
	if err != nil {
		http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	b, err := book.CreateBook()
	if err != nil {
		http.Error(w, "Error creating book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "Error marshaling created book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.PathValue("id")
	book := &models.Book{}
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing", err)
	}
	deletedBook, err := book.DeleteBook(ID)
	if err != nil {
		fmt.Println("Error while deleting book", err)
	}
	if deletedBook == nil {
		http.Error(w, "Error deleting book: ", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(deletedBook)
	if err != nil {
		http.Error(w, "Error marshaling deleted book: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.PathValue("id")
	book := &models.Book{}
	err := utils.ParseBody(r, book)
	if err != nil {
		fmt.Println("Error parsing body", err)
	}
	// Parse the book ID from the URL
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Error parsing book ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new book instance

	// Parse the request body into the book instance
	if err := utils.ParseBody(r, &book); err != nil {
		http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call the model function to update the book in the database
	updatedBook, err := models.UpdateBook(ID, *book)
	if err != nil {
		http.Error(w, "Error updating book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal the updated book to JSON and send the response
	res, err := json.Marshal(updatedBook)
	if err != nil {
		http.Error(w, "Error marshaling updated book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

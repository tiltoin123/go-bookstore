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
    newBooks:= models.GetAllBooks() // Supondo que GetAllBooks retorne um erro
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
	bookDetails, _:=models.GetBookById(ID)
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
	CreateBook := &models.Book{}
	utils.ParseBody(r,CreateBook)
	b:=CreateBook.CreateBook()
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
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error while parsing",err)
	}
	book:=models.DeleteBook(ID)
		if book==nil{
		http.Error(w, "Error deleting book: ", http.StatusInternalServerError)
        return
	}
	res, err := json.Marshal(book)
	if err != nil {
        http.Error(w, "Error marshaling deleted book: "+err.Error(), http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	UpdateBook := &models.Book{}
	utils.ParseBody(r,UpdateBook)
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("Error while parsing",err)
	}
	book:=models.UpdateBook(ID,*UpdateBook)
	if book==nil{
		http.Error(w, "Error updating book: ", http.StatusInternalServerError)
        return
	}
	res, err := json.Marshal(book)
	if err != nil {
        http.Error(w, "Error marshaling updated book: "+err.Error(), http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
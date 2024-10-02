package models

import (
	"database/sql"
	"fmt"

	"github.com/tiltoin123/go-bookstore/pkg/config"
)

var db *sql.DB

type Book struct {
    ID          int64  `json:"id"`
    Name        string `json:"name"`
    Author      string `json:"author"`
    Publication string `json:"publication"`
    CreatedAt    string`json:"created_at"`
    UpdatedAt    string`json:"updated_at"`
    DeletedAt    string`json:"deleted_at"`
}

func init(){
	config.Connect()
	config.GetDB()
}

func (b *Book) CreateBook() (*Book, error) {
    stmt, err := db.Prepare(`INSERT INTO books (name, author, publication, created_at, updated_at) 
                             VALUES (?, ?, ?, NOW(), NOW()) 
                             RETURNING id, name, author, publication, created_at, updated_at`)
    if err != nil {
        fmt.Println("Error preparing statement:", err)
        return nil, err
    }
    defer stmt.Close()

    err = stmt.QueryRow(b.Name, b.Author, b.Publication).Scan(&b.ID, &b.Name, &b.Author, &b.Publication)
    if err != nil {
        fmt.Println("Error executing query:", err)
        return nil, err
    }

    return b, nil
}

func UpdateBook(Id int64, b Book) (*Book, error) {
    stmt, err := db.Prepare(`UPDATE books SET name = ?, author = ?, publication = ?, updated_at = NOW() WHERE id = ? RETURNING id, name, author, publication, created_at, updated_at`)
    if err != nil {
        fmt.Println("Error preparing statement:", err)
        return nil, err
    }
    defer stmt.Close()

    err = stmt.QueryRow(b.Name, b.Author, b.Publication, b.ID).Scan(&b.ID, &b.Name, &b.Author, &b.Publication, &b.CreatedAt, &b.UpdatedAt)
    if err != nil {
        fmt.Println("Error updating book:", err)
        return nil, err
    }

    return &b, nil
}

func (b *Book) GetAllBooks() ([]Book, error) {
    var books []Book

    stmt, err := db.Prepare(`SELECT id, name, author, publication, created_at, updated_at FROM books`)
    if err != nil {
        fmt.Println("Error preparing statement:", err)
        return nil, err
    }
    defer stmt.Close()

    rows, err := stmt.Query()
    if err != nil {
        fmt.Println("Error executing query:", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var book Book
        if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Publication, &book.CreatedAt, &book.UpdatedAt); err != nil {
            fmt.Println("Error scanning book:", err)
            return nil, err
        }
        books = append(books, book)
    }

    if err = rows.Err(); err != nil {
        fmt.Println("Error iterating rows:", err)
        return nil, err
    }

    return books, nil
}

func (b *Book) GetBookById(Id int64) (*Book, error) {
    var book Book
    
    stmt, err := db.Prepare(`SELECT id, name, author, publication, created_at, updated_at, deleted_at FROM books WHERE id = ?`)
    if err != nil {
        fmt.Println("Error preparing statement:", err)
        return nil, err
    }
    defer stmt.Close()

    err = stmt.QueryRow(Id).Scan(&book.ID, &book.Name, &book.Author, &book.Publication, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt)
    if err != nil {
        fmt.Println("Error executing query:", err)
        return nil, err
    }

    return &book, nil
}

func (b *Book) DeleteBook(Id int64) (*Book, error) {
    stmt, err := db.Prepare(`UPDATE books SET deleted_at = NOW() WHERE id = ? RETURNING id, name, author, publication, created_at, updated_at, deleted_at`)
    if err != nil {
        fmt.Println("Error preparing statement:", err)
        return nil, err
    }
    defer stmt.Close()

    err = stmt.QueryRow(Id).Scan(&b.ID, &b.Name, &b.Author, &b.Publication, &b.CreatedAt, &b.UpdatedAt, &b.DeletedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("No book found with the given ID")
            return nil, nil
        }
        fmt.Println("Error deleting book:", err)
        return nil, err
    }

    return b, nil
}
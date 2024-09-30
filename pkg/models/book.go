package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/tiltoin123/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `gorm:"" json:"author"`
	Publication string `gorm:"" json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	if err:=db.Error; err!=nil{
		fmt.Println("error connecting",err)
		return
	}
	db.AutoMigrate(&Book{})
	if err:=db.Error; err!=nil{
		fmt.Println("error migrating",err)
		return
	}
}

func(b *Book) CreateBook() *Book{
	db.NewRecord(b)
	if err:=db.Error; err!=nil{
		fmt.Println("error checking primary key",err)
		return nil
	}
	db.Create(&b)
	if err:=db.Error; err!=nil{
		fmt.Println("error inserting data",err)
		return nil
	}
	return b
}

func UpdateBook(Id int64, Data Book) *Book {
    var book Book
    
    if err := db.Where("ID = ?", Id).First(&book).Error; err != nil {
        fmt.Println("Error finding book:", err)
        return nil
    }
    
    if err := db.Model(&book).Updates(Data).Error; err != nil {
        fmt.Println("Error updating book:", err)
        return nil
    }
    
    return &book
}

func GetAllBooks() []Book{
	var Books []Book
	if err:=db.Find(&Books).Error; err!=nil {
		fmt.Println("Error listing books",err)
		return nil
	}
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
    var getBook Book
    db := db.Where("ID = ?", Id).Find(&getBook)
    
    if err := db.Error; err != nil {
        fmt.Println("Error finding book:", err)
        return nil, nil
    }
    
    return &getBook, db
}

func DeleteBook(Id int64) *Book {
    var book Book
    if err := db.Where("ID = ?", Id).Delete(&book).Error; err != nil {
        fmt.Println("Error deleting book:", err)
        return nil
    }
    return &book
}

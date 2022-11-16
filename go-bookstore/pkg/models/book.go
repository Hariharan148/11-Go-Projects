package main

import (
	"github.com/Hariharan148/11-Go-Projects/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)


var db *gorm.DB

type Book struct{
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author`
	Publication string `json:publication`
}


func init(){
	config.Connect()
	db = config.getDB()
	db.AutoMigrate(&Book{})
}


func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}


func GetAllBooks() []Book{
	var books []Book
	db.Find(&books)
	return books

}

func GetBookById(id int64)  (*Book, *gorm.DB){
	var GetBook Book
	db := db.Where("ID=?", id).Find(&GetBook)
	return &GetBook , db
}


func DeleteBook(id int64) Book{
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
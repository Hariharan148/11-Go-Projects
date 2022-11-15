package main 

import (
	"github.com/jinzhu/gorm"
	"github.com/Hariharan148/11-Go-Projects/go-bookstore/pkg/config"
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
	DB.AutoMigrate(&Book{})
}
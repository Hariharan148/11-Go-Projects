package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	db *gorm.DB
)


func Connect(){
	d, err:= gorm.Open("mysql", "hari: Hariharan123/hariharan12?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}


func getDB() *gorm.DB{
	return db
}

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Gorilla/mux"
	"github.com/Hariharan148/11-Go-Projects/go-bookstore/pkg/models"
)


var NewBook models.Book


func GetBook(w http.ResponseWriter, r *http.Request){
	newbook := models.GetAllBooks()
	res, _ := json.Marshal(newbook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)

}


func GetBookById(w http.ResponseWriter, r * http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Prinln("Error while parsing")
	}
	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.http.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(http.StatusOk)
	w.Write(res)

}
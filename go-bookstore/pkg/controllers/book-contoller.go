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
	w.http.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	http.WriteHeader(http.StatusOk)
	http.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book:= models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var BookUpdate = &models.Book{}
	vars := mux.Vars(r)
	utils.ParseBody(r, vars)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing ")
	}
	bookDetails, db := models.GetBookById(ID)
	if BookUpdate.Name != ""{
		bookDetails.Name = BookUpdate.Name
	}
	if BookUpdate.Author != ""{
		bookDetails.Author = BookUpdate.Author
	}
	if BookUpdate.Publication != ""{
		bookDetails.Publication = BookUpdate.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}
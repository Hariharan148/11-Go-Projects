package routes

import (
	"github.com/Hariharan148/11-Go-Projects/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router){
	router.Handlefunc("/book/", controllers.Getbook).Methods("GET")
	router.Handlefunc("/book/", controllers.CreateBook).Methods("POST")
	router.Handlefunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.Handlefunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.Handlefunc("/book/{bookId}", controllers.GetBookbyId).Methods("GET")
}
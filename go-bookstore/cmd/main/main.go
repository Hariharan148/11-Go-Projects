package main

import (
	"log"
	"net/http"

	"github.com/Hariharan148/11-Go-Projects/go-bookstore/pkg/routes"
	"github.com/Hariharan148/11-Go-Projects/pkg/routes"
	"github.com/gorilla/mux"
)


func main(){
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("Localhost:8080", r))
}




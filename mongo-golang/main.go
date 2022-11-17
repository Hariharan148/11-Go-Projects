package main

import (
	"net/http"

	"github.com/Hariharan148/11-Go-Projects/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main(){
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id",uc.Getuser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser) 
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session{
	s, err := mgo.Dial("mongodb://localhost:27001")
	if err != nil{
		panic(err)
	}
	return s
}
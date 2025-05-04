package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/manish39x/mongo-golang/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	sesion, err := mgo.Dial("mongodb+srv://Manish:manish9062@cluster0.wqew4eo.mongodb.net")
	if err != nil {
		panic(err)
	}
	return sesion
}

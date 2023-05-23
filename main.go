package main

import (
	"log"
	"net/http"

	"github.com/emorydu/mongo-go/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	log.Println("Starting server at port 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://127.0.0.1:27107")
	if err != nil {
		panic(err)
	}
	return s
}

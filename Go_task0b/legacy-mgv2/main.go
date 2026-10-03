package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"legacy/controllers"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	db := session.DB("golang_task")
	controllers.Collection = db.C("users")

	router := httprouter.New()

	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:id", controllers.GetUser)
	router.DELETE("/user/:id", controllers.DeleteUser)

	fmt.Println("Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}

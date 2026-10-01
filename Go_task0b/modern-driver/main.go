package moderndriver
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type User struct {
	ID   bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string        `bson:"name" json:"name"`
	Age  int           `bson:"age" json:"age"`
}

var client *mongo.Client

func getCollection() *mongo.Collection {
	return client.Database("taskdb").Collection("users")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u.ID = bson.NewObjectID()
	getCollection().InsertOne(ctx, u)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// URL se ID nikalne ka simple tarika: /user/ID
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var u User
	err = getCollection().FindOne(ctx, bson.M{"_id": objID}).Decode(&u)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getCollection().DeleteOne(ctx, bson.M{"_id": objID})

	w.Write([]byte("User deleted"))
}

// Single handler function requests to manage methods
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createUser(w, r)
	} else if r.Method == "GET" {
		getUser(w, r)
	} else if r.Method == "DELETE" {
		deleteUser(w, r)
	}
}

func main() {
	var err error
	client, err = mongo.Connect(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/user/", userHandler)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
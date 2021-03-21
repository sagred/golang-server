package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client *mongo.Client

type Person struct {
	ID 		  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty`
	FirstName string  `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName string `json:"lastname,omitempty" bson""lastname,omitempty"`
}

func createPerson(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("Person")
}

func getPeople(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("People")
}

func main() {
	fmt.Println("Started Now")
	ctx,_  := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("")
	r := mux.NewRouter()

	r.HandleFunc("/person", createPerson).Methods("POST")
	r.HandleFunc("/people", getPeople).Methods("GET")


	log.Fatal(http.ListenAndServe(":8000", r))
}


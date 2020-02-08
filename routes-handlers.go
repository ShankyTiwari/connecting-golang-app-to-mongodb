package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// RenderHome Rendering the Home Page
func RenderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/index.html")
}

// GetUsers This function will return the response based ono user found in Database
func GetUsers(response http.ResponseWriter, request *http.Request) {
	// Creating  struct for Records
	type User struct {
		Name    string
		Country string
	}
	// Creating Variable of type User struct to hold result from database query
	var results []*User

	// Reading data from the query params
	username := mux.Vars(request)["name"]

	// Getting the instance of the collection from MongoDB Database
	collection := Client.Database("test").Collection("users")

	// Writing query to fetch the Data from the `users` collection
	databaseCursor, err := collection.Find(context.TODO(), bson.M{"name": bson.M{"$regex": username}})

	if err != nil {
		log.Fatal(err)
	}

	// Iterating over the MongoDB Cursor to decode the results
	for databaseCursor.Next(context.TODO()) {
		var elem User
		err := databaseCursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		// log.Println(elem.Name, elem.Country)

		// Appending it into the results variable
		results = append(results, &elem)
	}

	if err := databaseCursor.Err(); err != nil {
		log.Fatal(err)
	}

	jsonResponse, jsonError := json.Marshal(results)

	if jsonError != nil {
		log.Fatal(jsonError)
		returnErrorResponse(response, request)
	}

	if jsonResponse == nil {
		returnErrorResponse(response, request)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}

}

// Helper function to handle the HTTP response
func returnErrorResponse(response http.ResponseWriter, request *http.Request) {
	jsonResponse, err := json.Marshal("It's not you it's me.")
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusInternalServerError)
	response.Write(jsonResponse)
}

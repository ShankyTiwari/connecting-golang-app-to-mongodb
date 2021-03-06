package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

// AddApproutes will add the routes for the application
func AddApproutes(route *mux.Router) {

	log.Println("Loadeding Routes...")

	setStaticFolder(route)

	route.HandleFunc("/", RenderHome)

	route.HandleFunc("/users/{name}", GetUsers).Methods("GET")

	log.Println("Routes are Loaded.")
}

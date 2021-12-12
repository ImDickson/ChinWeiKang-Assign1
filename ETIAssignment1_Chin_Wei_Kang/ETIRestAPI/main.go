package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

//initialize router
func initializeRouter() {
	r := mux.NewRouter()

	//provide route information
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
   //Not used r.HandleFunc("/users/{id}", DeleteUser).Methods("Delete")


	//passing the port information
	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	InitialMigration()
	//calling the router
	initializeRouter()
}

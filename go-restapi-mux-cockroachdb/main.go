package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func handleRequest() {
	// Init router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", NewUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal("Error while creating router", http.ListenAndServe(":8000", myRouter))
}
func main() {
	fmt.Println("Welcome to REST API in go with cockroachdb")
	dbConnection := CreateDbConnection()
	CreateTable(dbConnection)
	handleRequest()
	defer dbConnection.Close()
}

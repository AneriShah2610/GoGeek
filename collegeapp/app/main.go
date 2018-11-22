package main

import (
	"log"
	"net/http"

	"github.com/AneriShah2610/GoGeek/collegeapp/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/college", api.GetCollege).Methods("GET")
	router.HandleFunc("/college", api.CreateCollege).Methods("POST")
	router.HandleFunc("/college/{id}", api.DeleteCollege).Methods("DELETE")
	router.HandleFunc("/college/{id}", api.UpdateCollege).Methods("PUT")
	log.Fatal("error while routing : ", http.ListenAndServe(":8000", router))
}

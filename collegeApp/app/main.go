package main

import (
	"log"
	"net/http"

	"github.com/AneriShah2610/GoGeek/collegeApp/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(api.MiddleWareHandler)
	router.HandleFunc("/college", api.ShowColleges).Methods("GET")
	log.Fatal("error while routing", http.ListenAndServe(":8000", router))
}

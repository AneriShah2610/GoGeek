package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	// College model
	College struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Address *Address
	}
	// Address model
	Address struct {
		City    string `json:"city"`
		State   string `json:"state"`
		Country string `json:"country"`
		PinCode string `json:"pincode"`
	}
)

var collegeData []College

func routes() {
	router := mux.NewRouter()
	router.HandleFunc("/data", GetData).Methods("GET")
	router.HandleFunc("/data/{id}", GetCollege).Methods("GET")
	router.HandleFunc("/data", CreateCollge).Methods("POST")
	router.HandleFunc("/data/{id}", DeleteCollege).Methods("DELETE")
	log.Fatal("Error at roueter handle : ", http.ListenAndServe(":8000", router))
}
func main() {

	collegeData = append(collegeData, College{"1", "Aneri", &Address{"Ahmedabad", "Gujarat", "India", "380015"}})
	routes()
}

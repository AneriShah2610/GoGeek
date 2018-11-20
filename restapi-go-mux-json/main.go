package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	// Person model
	Person struct {
		ID        string   `json:"id"`
		FirstName string   `json:"firstname"`
		LastName  string   `json:"lastname"`
		Address   *Address `json:"address"`
	}
	// Address model
	Address struct {
		City    string `json:"city"`
		State   string `json:"state"`
		Country string `json:"country"`
	}
)

var people []Person

func main() {

	router := mux.NewRouter()
	people = append(people, Person{"1", "Aneri", "Shah", &Address{"Ahmedabad", "Gujarat", "India"}})
	people = append(people, Person{"2", "Anokhi", "Shah", &Address{"Ahmedabad", "Gujarat", "India"}})
	people = append(people, Person{"3", "Ani", "Shah", &Address{"Ahmedabad", "Gujarat", "India"}})
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal("error while routing HandleFunc ", http.ListenAndServe(":8000", router))

}

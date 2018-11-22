package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetPeople get all details of people
func GetPeople(writer http.ResponseWriter, requets *http.Request, _ httprouter.Params) {
	json.NewEncoder(writer).Encode(people)
}

// GetPerson get person detail
func GetPerson(writer http.ResponseWriter, requets *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	for _, item := range people {
		if item.ID == id { // If this condition teue than it returns item
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Person{})
}

// CreatePerson create new person
func CreatePerson(writer http.ResponseWriter, requets *http.Request, ps httprouter.Params) {
	var person Person
	_ = json.NewDecoder(requets.Body).Decode(&person)
	people = append(people, person)
	json.NewEncoder(writer).Encode(people)
}

// DeletePerson delete person by their id
func DeletePerson(writer http.ResponseWriter, requets *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(people)
}

package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	// Person model
	Person struct {
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

var people []Person

func routes() {
	router := httprouter.New()
	router.GET("/people", GetPeople)
	router.GET("/people/:id", GetPerson)
	router.POST("/people", CreatePerson)
	router.DELETE("/people/:id", DeletePerson)
	log.Fatal("Error at roueter handle : ", http.ListenAndServe(":8000", router))
}
func main() {

	people = append(people, Person{"1", "Aneri", &Address{"Ahmedabad", "Gujarat", "India", "380015"}})
	routes()
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Data struct {
		Stuff Category
	}
	Category struct {
		Fruits     Fruit
		Vegetables Vegetable
	}
	Fruit     map[string]int
	Vegetable map[string]int
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", serverRest)
	log.Fatal("Error at routin function", http.ListenAndServe(":8000", router))
}
func serverRest(writer http.ResponseWriter, request *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		log.Fatal("Error while marshling", err)
	}
	fmt.Fprintf(writer, string(response))
}
func getJsonResponse() ([]byte, error) {

	fruit := make(map[string]int)
	fruit["Apple"] = 25
	fruit["Mango"] = 1
	fruit["Orange"] = 2

	vegetable := make(map[string]int)
	vegetable["Cucmber"] = 22
	vegetable["Tomato"] = 4
	vegetable["Corn"] = 25

	c := Category{fruit, vegetable}
	d := Data{c}
	return json.MarshalIndent(d, " ", "")
}

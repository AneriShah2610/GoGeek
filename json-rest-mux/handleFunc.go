package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetData get all college details
func GetData(writer http.ResponseWriter, requets *http.Request) {
	json.NewEncoder(writer).Encode(collegeData)
}

// GetCollege get college detail
func GetCollege(writer http.ResponseWriter, requets *http.Request) {
	params := mux.Vars(requets)
	for _, item := range collegeData {
		if item.ID == params["id"] { // If this condition teue than it returns item
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&College{})
}

// CreateCollge create new college data
func CreateCollge(writer http.ResponseWriter, requets *http.Request) {
	var college College
	_ = json.NewDecoder(requets.Body).Decode(&college)
	collegeData = append(collegeData, college)
	json.NewEncoder(writer).Encode(collegeData)
}

// DeleteCollege delete college by their id
func DeleteCollege(writer http.ResponseWriter, requets *http.Request) {
	params := mux.Vars(requets)
	for index, item := range collegeData {
		if item.ID == params["id"] {
			collegeData = append(collegeData[:index], collegeData[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(collegeData)
}

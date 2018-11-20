package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("html/*.html"))
	// Init Router
	router := mux.NewRouter()
	router.HandleFunc("/", indexHTML).Methods("GET")
	router.HandleFunc("/about", aboutHTML).Methods("GET")
	router.HandleFunc("/contact", contactHTML).Methods("GET")
	http.Handle("/", router)
	log.Fatal("Error while Handle Function", http.ListenAndServe(":8000", nil))
}

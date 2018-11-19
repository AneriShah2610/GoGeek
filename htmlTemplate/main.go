package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	// Init router
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

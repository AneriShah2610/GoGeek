package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("html/*.html"))
}

func main() {
	http.HandleFunc("/", indexHTML)
	http.HandleFunc("/about", aboutHTML)
	http.HandleFunc("/contact", contactHTML)
	log.Fatal("Error while handling HandleFunction", http.ListenAndServe(":8000", nil))
}

package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	// BookData struct(model)
	BookData struct {
		BookListing []Book
	}
	// Book struct (model)
	Book struct {
		ID             int     `json:"id"`
		Isbn           string  `json:"isbn"`
		Author         string  `json:"author"`
		Title          string  `json:"title"`
		PubliationDate string  `json:"pubdate"`
		Price          float64 `json:"price"`
	}
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("server/books/*.html"))

	// Init router
	router := mux.NewRouter()
	router.HandleFunc("/books/", booksHandler)
	router.Handle("/", http.FileServer(http.Dir("client")))
	log.Fatal("Error while routing handlefunction", http.ListenAndServe(":8000", router))
}

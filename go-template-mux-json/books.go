package main

import (
	"html/template"
	"log"
	"net/http"
)

func booksHandler(w http.ResponseWriter, r *http.Request) {
	// Read book data from book.json
	bookData := &BookData{
		laodBooks("books.json"),
	}
	templates := template.Must(template.ParseFiles("server/books/index.html"))
	err := templates.Execute(w, bookData)
	if err != nil {
		log.Fatal(err)
	}
}

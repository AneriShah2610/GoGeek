package main

import (
	"html/template"
	"log"
	"net/http"
)

func booksHandler(w http.ResponseWriter, r *http.Request) {
	// Read book data from book.json
	bookData := &BookData{
		LoadBooks("books.json"),
	}
	templates := template.Must(template.ParseFiles("server/books/index.html")) //It parse index.html file
	err := templates.Execute(w, bookData)
	if err != nil {
		log.Fatal(err)
	}
}

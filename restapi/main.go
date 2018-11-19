package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type (
	// Book struct (Model)
	Book struct {
		ID     string  `json:"id"`
		Isbn   string  `json:"isbn"`
		Title  string  `json:"title"`
		Price  float64 `json:"price"`
		Author *Author `json:"author"`
	}
	// Author struct (Model)
	Author struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}
)

// Get all Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookid := mux.Vars(r)
	for _, bookData := range books {
		if bookData.ID == bookid["id"] {
			json.NewEncoder(w).Encode(bookData)
		}
	}
}

// Create a Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookid := mux.Vars(r)
	for index, bookData := range books {
		if bookData.ID == bookid["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(100))
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
		}
	}
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

// Init books variable as a slice Book struct
var books []Book

func main() {
	// Init router
	router := mux.NewRouter()

	// Mock data
	books = append(books, Book{"1", "457fg12", "Go Book", 542.50, &Author{"Abc", "Xyz"}})
	books = append(books, Book{"2", "5844ff7", "Go Action", 805.00, &Author{"Amx", "Gdf"}})
	books = append(books, Book{"3", "fs77d4g", "Go Module", 542.50, &Author{"fhf", "rge"}})
	books = append(books, Book{"4", "88fef54", "Go RESTAPI", 225.00, &Author{"aki", "deg"}})

	// router handler
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

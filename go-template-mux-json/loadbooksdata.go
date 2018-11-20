package main

import (
	"encoding/json"
	"log"
	"os"
)

func LoadBooks(fn string) []Book {
	// craete a variable of Book Struct
	var books []Book

	// Open book.json file
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the book.json file
	bookDecoder := json.NewDecoder(file)
	bookDecoder.Decode(&books)
	// return decoded book.json file
	return books
}

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID       string `json:"id"`
	Response string `json:"responce"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	books = append(books, Book{ID: "1", Response: "ti pidor"})

	r.HandleFunc("/api/books", getBooks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}

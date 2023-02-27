package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One",
		Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "847564", Title: "Book Two",
		Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers/ Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

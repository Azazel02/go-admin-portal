package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Init Router
	r := mux.NewRouter()
	openDB()

	// Mock Data - @todo implement DB
	//books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One",
	//	Author: &Author{Firstname: "John", Lastname: "Doe"}})
	//books = append(books, Book{ID: "2", Isbn: "847564", Title: "Book Two",
	//	Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers/ Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
	//  http://localhost:8000/api/books

}

var DB *sql.DB

func openDB() (*sql.DB, error) {

	const (
		DATABASE_HOST     = "localhost"
		DATABASE_PORT     = 5432
		DATABASE_USER     = "admin"
		DATABASE_PASSWORD = "admin"
		DATABASE_NAME     = "kartik"
	)

	psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s 
	password=%s dbname=%s`, DATABASE_HOST, DATABASE_PORT,
		DATABASE_USER, DATABASE_PASSWORD, DATABASE_NAME)
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("DB connection not established")
		log.Println(err)
		return nil, err
	}

	if err = DB.Ping(); err != nil {
		log.Println("DB connection not established")
		log.Println(err)
		return nil, err
	}
	log.Println("DB connection established")
	return DB, nil

}

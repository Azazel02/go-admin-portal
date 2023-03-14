package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

// Author Struct
type Author struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {

	/*	w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)*/
	//logPrefix := "BooksGet"
	books := []*Book{}

	stmt, err := DB.Prepare(`SELECT books.id,books.isbn,books.title,author.firstname,author.lastname from Books inner join author on author.id =books.Author_id`)
	if err != nil {
		//log.Println("%s failed to prepare stmt", err.Error())
		log.Printf("%s failed to prepare stmt", err.Error())
		//log.Print("%s failed to prepare stmt", err.Error())
		return
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("%s query error", err.Error())
		return
	}
	for rows.Next() {
		var data_rec Book
		data_rec.Author = &Author{}
		err = rows.Scan(&data_rec.ID, &data_rec.Isbn, &data_rec.Title, &data_rec.Author.Firstname, &data_rec.Author.Lastname)
		if err != nil {
			log.Printf("%s failed scan", err.Error())
			return
		}
		books = append(books, &data_rec)
		log.Printf("current books value : %+v", books)
	}
	response, err := json.Marshal(books)
	if err != nil {
		log.Printf("%s failed to marshal", err.Error())
		return
	}
	_, _ = w.Write(response)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			books = append(books, item)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

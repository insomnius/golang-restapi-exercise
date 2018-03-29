package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"fmt"
	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books
var books[] Book


// Get all the books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	for _,item := range books {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Create a book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(5000))
	
	books = append(books, book) 

	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"]{
			_ = json.NewDecoder(r.Body).Decode(&item)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main(){
	// Init router
	route := mux.NewRouter()

	// Mock data
	books 	= append(books, Book{ID: "1", Isbn: "123456", Title: "Title 1", Author: &Author{Firstname: "Arief", Lastname: "Rahman"}})
	books 	= append(books, Book{ID: "2", Isbn: "654321", Title: "Title 2", Author: &Author{Firstname: "Arief", Lastname: "Rahman"}})
	books 	= append(books, Book{ID: "3", Isbn: "098765", Title: "Title 3", Author: &Author{Firstname: "Arief", Lastname: "Rahman"}})
	books 	= append(books, Book{ID: "4", Isbn: "123678", Title: "Title 4", Author: &Author{Firstname: "Arief", Lastname: "Rahman"}})
	books 	= append(books, Book{ID: "5", Isbn: "098123", Title: "Title 5", Author: &Author{Firstname: "Arief", Lastname: "Rahman"}})


	// Route handler
	route.HandleFunc("/api/books", getBooks).Methods("GET")
	route.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	route.HandleFunc("/api/books/create", createBook).Methods("POST")
	route.HandleFunc("/api/books/{id}/update", updateBook).Methods("PUT")
	route.HandleFunc("/api/books/{id}/delete", deleteBook).Methods("DELETE")

	fmt.Println("Starting the server at port 8000")

	log.Fatal(http.ListenAndServe(":8000", route))
}
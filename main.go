package main

import (
	// parse data into json
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

// Book struct // (Models)
// we use the tag to create the json format
type Book struct {
	ID     string  "json:id"
	ISBN   string  "json:isbn"
	Title  string  "json:title"
	Author *Author "json:author"
	Description string "json:description"
}

// Author struct
type Author struct {
	FirstName string "json:firstname "
	LastName  string "json:lastname"
}

// Init books var as a slice Book struct
var books = make([]Book, 0)

func main() {

	// allow cors requests
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With","Access-Control-Allow-Origin"})
    allowedOrigins := handlers.AllowedOrigins([]string{"*"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	
	// init router
	r := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", ISBN: "01319", Title: "Golang For Dummies", Description: "Learn how to code by using a dummies based approach", Author: &Author{FirstName: "Musiteli", LastName: "Mubuso"}})
	books = append(books, Book{ID: "2", ISBN: "45698", Title: "Golang For Nerds", Description: "Learn how to code by using a nerds based approach", Author: &Author{FirstName: "Musiteli", LastName: "Mubuso"}})
	books = append(books, Book{ID: "3", ISBN: "45645", Title: "Golang For Dudes", Description: "Learn how to code by using a dude based approach", Author: &Author{FirstName: "Musiteli", LastName: "Mubuso"}})
	books = append(books, Book{ID: "4", ISBN: "15445", Title: "Golang For Gangsters", Description: "Learn how to code by using a gangs based approach", Author: &Author{FirstName: "Musiteli", LastName: "Mubuso"}})


	// Router Handlers	/ Endpoints
	r.HandleFunc("/api/v1/books", getBooks).Methods("GET")
	r.HandleFunc("/api/v1/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/v1/books", createBook).Methods("POST")
	r.HandleFunc("/api/v1/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/v1/books/{id}", deleteBook).Methods("DELETE")

	// We listen on port 8000 using our router r
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedHeaders,allowedMethods,allowedOrigins)(r)))
	fmt.Println("Server is running on port 8000")
}

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	// Loop through books and find with id
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			fmt.Println("inside for loop")
			return
		}
	}
	// If id is not found return default book struct
	json.NewEncoder(w).Encode(&Book{})
}

// create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	// We use the ampersand to make sure pass the book by reference and not value
	_= json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000000)) 
	books = append(books, book)
	fmt.Println(book)
	fmt.Printf("%+v\n", books)
	json.NewEncoder(w).Encode(book)
}

// delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	for key, book := range books{
		if book.ID == params["id"] {
			books = append(books[:key],books[key+1:]...)
			fmt.Println(books)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var book Book;
	params := mux.Vars(r) //Get params
	for _, item := range books {
		if item.ID == params["id"] {
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			item = book
			fmt.Println(item)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

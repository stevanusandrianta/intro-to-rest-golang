package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:title`
	Year   string `json:title`
}

var books []Book

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("get books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("get book")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("add books")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("update book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("remove book")
}

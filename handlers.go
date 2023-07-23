package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var books []Book
var nextID = 1
const invalidBookIdMessage = "Invalid Book ID"
const bookNotFoundMessage = "Book not found"


func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	id, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, invalidBookIdMessage, http.StatusBadRequest)
		return
	}
	var foundBook *Book
	for _, book := range books {
		if book.ID == id {
			foundBook = &book
			break
		}
	}
	if foundBook == nil {
		http.Error(w, bookNotFoundMessage, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(foundBook)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	newBook.ID = nextID
	nextID++
	books = append(books, newBook)
	w.WriteHeader(http.StatusCreated)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	id, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, invalidBookIdMessage, http.StatusBadRequest)
		return
	}
	var updatedBook Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	index := -1
	for i, book := range books {
		if book.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		http.Error(w, bookNotFoundMessage, http.StatusNotFound)
		return
	}
	books[index] = updatedBook
	w.WriteHeader(http.StatusNoContent)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	id, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, invalidBookIdMessage, http.StatusBadRequest)
		return
	}
	index := -1
	for i, book := range books {
		if book.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		http.Error(w, bookNotFoundMessage, http.StatusNotFound)
		return
	}
	books = append(books[:index], books[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

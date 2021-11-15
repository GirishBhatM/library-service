package controller

import (
	"encoding/json"
	"fmt"
	"library-service/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutes() {
	libraryRoutes := mux.NewRouter().StrictSlash(true)
	libraryRoutes.HandleFunc("/", indexPage).Methods("GET")
	libraryRoutes.HandleFunc("/books", getAllBooks).Methods("GET")
	libraryRoutes.HandleFunc("/books", addBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", libraryRoutes))
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, e := service.GetAllBooks(w, r)
	w.Header().Set("Content-Type", "application/json")
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error" : "Error occured"}`)
	} else {
		json.NewEncoder(w).Encode(books)
	}
}

func getBookById(w http.ResponseWriter, r *http.Request) {

}

func addBook(w http.ResponseWriter, r *http.Request) {
	book, e := service.AddBook(w, r)
	w.Header().Set("Content-Type", "application/json")
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error" : "Error occured"}`)
	} else {
		json.NewEncoder(w).Encode(book)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to index!")
}

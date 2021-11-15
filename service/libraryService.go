package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"library-service/model"
	"net/http"
)

var bookStore = make(map[int]model.Book)
var id int = 0

func GetAllBooks(w http.ResponseWriter, r *http.Request) ([]model.Book, error) {
	books := []model.Book{}
	for _, v := range bookStore {
		books = append(books, v)
	}
	return books, nil
}

func AddBook(w http.ResponseWriter, r *http.Request) (model.Book, error) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	book := model.Book{}
	error := json.Unmarshal(bodyBytes, &book)
	if error != nil {
		return model.Book{}, error
	} else {
		id++
		book.Id = id
		bookStore[book.Id] = book
		return book, nil
	}
}

func GetBook(bookId int) (model.Book, error) {
	book, exists := bookStore[bookId]
	if !exists {
		return model.Book{}, errors.New("No book exists")
	}
	return book, nil
}

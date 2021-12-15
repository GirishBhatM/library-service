package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"library-service/model"
	"library-service/repository"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) ([]model.Book, error) {
	books := repository.FindAll()
	return books, nil
}

func AddBook(w http.ResponseWriter, r *http.Request) (model.Book, error) {
	bodyBytes, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return model.Book{}, errors.New("invalid request. failed to map the structure to book object")
	}
	book := model.Book{}
	error := json.Unmarshal(bodyBytes, &book)
	if error != nil {
		return model.Book{}, error
	} else {
		book.Id = repository.Save(book)
		return book, nil
	}
}

func GetBook(bookId string) (model.Book, error) {
	book, exists := repository.GetById(bookId)
	if !exists {
		return model.Book{}, errors.New("no book exists")
	}
	return book, nil
}

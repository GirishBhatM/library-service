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
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	book := model.Book{}
	error := json.Unmarshal(bodyBytes, &book)
	if error != nil {
		return model.Book{}, error
	} else {
		book.Id = repository.Save(book)
		return book, nil
	}
}

func GetBook(bookId int) (model.Book, error) {
	book, exists := repository.GetById(bookId)
	if !exists {
		return model.Book{}, errors.New("no book exists")
	}
	return book, nil
}

package repository

import (
	"library-service/model"
	"testing"
)

func TestSaveBook(t *testing.T) {
	book := model.Book{Name: "Sample Book"}
	bookId := Save(book)
	if bookId != "random" {
		t.Errorf("expected %s found %s", "random", bookId)
	}
}

func TestGetById(t *testing.T) {
	book := model.Book{Name: "Sample Book"}
	bookId := Save(book)
	_, exists := GetById(bookId)
	if !exists {
		t.Errorf("No book found for id %d", bookId)
	}
}

func TestFindAll(t *testing.T) {
	Save(model.Book{Name: "Sample Book"})
	Save(model.Book{Name: "Sample Book2"})
	books := FindAll()
	length := len(books)
	if length != 4 {
		t.Errorf("expected book size is not matching, expected %d got %d", 4, length)
	}
}

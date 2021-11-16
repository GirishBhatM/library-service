package repository

import "library-service/model"

var bookStore = make(map[int]model.Book)
var id int = 0

func Save(book model.Book) int {
	id = id + 1
	book.Id = id
	bookStore[id] = book
	return id
}
func GetById(id int) (model.Book, bool) {
	book, exists := bookStore[id]
	return book, exists
}

func FindAll() []model.Book {
	books := []model.Book{}
	for _, book := range bookStore {
		books = append(books, book)
	}
	return books
}

package repository

import (
	"encoding/json"
	"library-service/model"

	"github.com/google/uuid"
)

var client = GetClient()

const KEY_PATTERN = "book_*"
const PAGE_SIZE = 1

func Save(book model.Book) string {
	id := uuid.New().String()
	book.Id = id
	bookJsonByte, _ := json.Marshal(book)
	client.Set("book_"+id, string(bookJsonByte), 0)
	return id
}
func GetById(id string) (model.Book, bool) {
	bookByte, err := client.Get("book_" + id).Result()
	if err != nil {
		return model.Book{}, false
	}
	book := model.Book{}
	error := json.Unmarshal([]byte(bookByte), &book)
	if error != nil {
		return model.Book{}, false
	}
	return book, true
}

func FindAll() []model.Book {
	var curs uint64
	books := []model.Book{}
	for {
		var keys []string
		var error error
		keys, curs, error = client.Scan(curs, KEY_PATTERN, PAGE_SIZE).Result()
		if error != nil {
			panic(error)
		}
		for _, key := range keys {
			bookJson, e := client.Get(key).Result()
			if e != nil {
				panic(e)
			}
			book := model.Book{}
			json.Unmarshal([]byte(bookJson), &book)
			books = append(books, book)
		}
		if curs == 0 {
			break
		}
	}
	return books
}

package main

import (
	"fmt"
	"library-service/model"
)

func main() {
	fmt.Println("hello library service")
	books := []model.Book{}
	book1 := model.Book{Id: "1", Name: "Harry Potter", Author: model.Author{Name: "Unknown"}}
	books = append(books, book1)
	books = append(books, book1)
	books = append(books, book1)
	library := model.Library{Books: books}
	fmt.Printf("%+v\n", library)
	fmt.Printf("%d", len(books))

}

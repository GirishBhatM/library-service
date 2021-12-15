package model

type Book struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Author Author `json:"author"`
}

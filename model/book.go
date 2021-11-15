package model

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author Author `json:"author"`
}

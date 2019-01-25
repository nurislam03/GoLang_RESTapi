package main

import (
	"log"
	"net/http"

	"github.com/nurislam03/GoLang_RESTapi/routes"

	"github.com/nurislam03/GoLang_RESTapi/data"
	"github.com/nurislam03/GoLang_RESTapi/model"
)

// Main function
func main() {
	// data 
	data.Books = append(data.Books, model.Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &model.Author{Firstname: "John", Lastname: "Doe"}})
	data.Books = append(data.Books, model.Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &model.Author{Firstname: "Steve", Lastname: "Smith"}})

	r := routes.Router()

	// Start server
	log.Println("Starting server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", &r))
}

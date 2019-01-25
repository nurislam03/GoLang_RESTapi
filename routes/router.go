package routes

import (
	"github.com/nurislam03/GoLang_RESTapi/api"
	"github.com/gorilla/mux"
)

func Router() mux.Router {
	// Init router
	Router := mux.NewRouter()

	// Route handles & endpoints
	Router.HandleFunc("/books", api.GetBooks).Methods("GET")
	Router.HandleFunc("/books/{id}", api.GetBook).Methods("GET")
	Router.HandleFunc("/books", api.CreateBook).Methods("POST")
	Router.HandleFunc("/books/{id}", api.UpdateBook).Methods("PUT")
	Router.HandleFunc("/books/{id}", api.DeleteBook).Methods("DELETE")

	return *Router
}

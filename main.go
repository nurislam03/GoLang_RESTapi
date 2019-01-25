package main

import (
	"log"
	"net/http"

	"github.com/nurislam03/go_restapi/routes"
)

// Main function
func main() {
	r := routes.Router()

	// Start server
	log.Println("Starting server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", &r))
}

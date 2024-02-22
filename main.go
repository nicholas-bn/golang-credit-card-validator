package main

import (
	"log"
	"net/http"

	"github.com/nicholas-bn/golang-credit-card-validator/server"
)

func main() {
	log.Println("Starting the server...")
	handler := http.HandlerFunc(server.CardValidatorServer)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

package main

import (
	"log"
	"net/http"

	"github.com/nicholas-bn/golang-credit-card-validator/server"
)

func main() {
	// if len(os.Args) > 1 {
	// 	portStr := os.Args[1]
	// 	portInt, err := strconv.Atoi(portStr)
	// 	if err != nil {
	// 		fmt.Printf("error retrieving the port input parameter: %s\n", err)
	// 		os.Exit(1)
	// 	}
	// 	api.Server(portInt)
	// } else {
	// 	api.Server(8080)
	// }

	handler := http.HandlerFunc(server.CardValidatorServer)
	log.Fatal(http.ListenAndServe(":8080", handler))

}

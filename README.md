# golang-credit-card-validator

Simple card number validator project to gain experience in Go using low-level modules.
Requires Go 1.21 and above (`slices` import).

## What does it do ?

* Sets up an http server that'll listen to incoming requests
* Will validate the input, or return a 400
* If the input schema is valid, we'll validate the card number using Luhn's algorithm.
* Return a JSON containing a response with the card number and a boolean representing the algorithm response.

## Unit tests

Run `go test ./...` to test all the module (`utilities`, `algorithm`, and `api`).
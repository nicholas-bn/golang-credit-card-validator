# Card number validator in Go

Simple card number validator project to gain experience in Go using low-level modules and TDD.
Requires Go 1.21 and above (`slices` import).

## What does it do ?

* Sets up an http server that'll listen to incoming requests
* Will validate the input, or return an appropriate HTTP code depending on the error
* If the input schema is valid, we'll validate the card number using Luhn's algorithm.
* Return a JSON containing a response with the card number and a boolean representing the algorithm response.

## Unit tests

This little module was developed fully using TDD so I could learn more about testing development in Go.

Run `go test ./...` to test all the module (`utilities`, `algorithm`, and `server`).
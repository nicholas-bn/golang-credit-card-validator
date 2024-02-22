package server

import (
	"fmt"
	"net/http"
)

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
func CardValidatorServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{\"credit_card\":\"4556737586899855\", \"is_valid\":true}")
}

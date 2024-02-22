package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestGETPlayers(t *testing.T) {
func TestGet(t *testing.T) {

	input := InputJson{
		CardNumber: "4556737586899855",
	}
	marshalled, err := json.Marshal(input)
	if err != nil {
		log.Fatalf("impossible to marshall card input: %s", err)
	}
	// fmt.Println(string(marshalled))
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api", bytes.NewReader(marshalled))
		response := httptest.NewRecorder()

		CardValidatorServer(response, request)

		got := response.Body.String()
		want := "{\"credit_card\":\"4556737586899855\", \"is_valid\":true}"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

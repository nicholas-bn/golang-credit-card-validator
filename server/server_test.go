package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestCardValidatorServerOKValidOrInvalid(t *testing.T) {

	t.Run("return card \"4556737586899855\" is valid", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api", bytes.NewReader([]byte("{\"card_number\":\"4556737586899855\"}")))
		response := httptest.NewRecorder()

		CardValidatorServer(response, request)

		outputJson := OutputJson{}
		byteList, errRead := io.ReadAll(response.Body)
		if errRead != nil {
			log.Fatalf("FAIL when reading output")
		}
		errUnmarshal := json.Unmarshal(byteList, &outputJson)
		if errUnmarshal != nil {
			log.Fatalf("FAIL when unmarshaling output")
		}

		assert.Equal(t, outputJson.CardNumber, "4556737586899855")
		assert.Equal(t, outputJson.IsValid, true)
		assertStatus(t, response.Code, http.StatusAccepted)
	})
	t.Run("return card \"2221004746855642\" is NOT valid", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api", bytes.NewReader([]byte("{\"card_number\":\"2221004746855642\"}")))
		response := httptest.NewRecorder()

		CardValidatorServer(response, request)

		outputJson := OutputJson{}
		byteList, errRead := io.ReadAll(response.Body)
		if errRead != nil {
			log.Fatalf("FAIL when reading output")
		}
		errUnmarshal := json.Unmarshal(byteList, &outputJson)
		if errUnmarshal != nil {
			log.Fatalf("FAIL when unmarshaling output")
		}

		assert.Equal(t, outputJson.CardNumber, "2221004746855642")
		assert.Equal(t, outputJson.IsValid, false)
		assertStatus(t, response.Code, http.StatusAccepted)
	})
}

func TestCardValidatorServerNOTOKWrongJson(t *testing.T) {
	t.Run("Wrong JSON format in input: missing quote", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api", bytes.NewReader([]byte("{\"card_number\":\"}")))
		response := httptest.NewRecorder()

		CardValidatorServer(response, request)

		assertResponseBody(t, response.Body.String(), "{\"error\":\"The JSON is not formatted correctly\",\"status_code\":400}")
		assertStatus(t, response.Code, http.StatusBadRequest)
	})
	t.Run("Wrong JSON format in input: card does not only contains nums", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api", bytes.NewReader([]byte("{\"card_number\":\"121323a1212\"}")))
		response := httptest.NewRecorder()

		CardValidatorServer(response, request)

		assertResponseBody(t, response.Body.String(), "{\"error\":\"The card number must contain only numbers\",\"status_code\":400}")
		assertStatus(t, response.Code, http.StatusBadRequest)
	})
}

func TestCardValidatorServerNOTOKOtherHttpVerbs(t *testing.T) {

	httpVerbsNotAllowed := []string{
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}

	for _, verb := range httpVerbsNotAllowed {
		t.Run(fmt.Sprintf("Used %s instead of GET", verb), func(t *testing.T) {
			request, _ := http.NewRequest(verb, "/api", bytes.NewReader([]byte("{\"card_number\":\"4556737586899855\"}")))
			response := httptest.NewRecorder()

			CardValidatorServer(response, request)

			assertResponseBody(t, response.Body.String(), "{\"error\":\"Method not allowed, only GET is supported\",\"status_code\":405}")
			assertStatus(t, response.Code, http.StatusMethodNotAllowed)
		})
	}

}

func assertResponseBody(t testing.TB, got string, want string) {
	t.Helper() // to notify testing package that this is an helper
	if got != want {
		t.Errorf("Response body is wrong, got \n%q \nwant %q", got, want)
	}
}

func assertStatus(t testing.TB, got int, want int) {
	t.Helper() // to notify testing package that this is an helper
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

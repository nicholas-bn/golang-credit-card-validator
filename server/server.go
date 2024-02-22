package server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/nicholas-bn/golang-credit-card-validator/algorithm"
	"github.com/nicholas-bn/golang-credit-card-validator/utilities"
)

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
func CardValidatorServer(w http.ResponseWriter, r *http.Request) {

	// Check http verb
	if r.Method != http.MethodGet {
		errorMsg, err := MapError(ErrorJson{StatusCode: http.StatusMethodNotAllowed, ErrorMessage: "Method not allowed, only GET is supported"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal server error: while mapping the error : %s", err)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, string(errorMsg))
		return
	}

	// Setup input var
	inputJson := InputJson{}

	// Retrieve input
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		errorMsg, err := MapError(ErrorJson{StatusCode: http.StatusBadRequest, ErrorMessage: "The JSON is not formatted correctly"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal server error: while mapping the error : %s", err)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, string(errorMsg))
		return
	}

	// Unmarshal the JSON
	mapErr := MapInput(body, &inputJson)
	if mapErr != nil {
		errorMsg, err := MapError(ErrorJson{StatusCode: http.StatusBadRequest, ErrorMessage: "The JSON is not formatted correctly"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal server error: while mapping the error : %s", err)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, string(errorMsg))
		return
	}

	// Cast the card number as int list and launch the validation
	cardNumAsIntList, castErr := utilities.StringToIntList(inputJson.CardNumber)
	if castErr != nil {
		errorMsg, err := MapError(ErrorJson{StatusCode: http.StatusBadRequest, ErrorMessage: "The card number must contain only numbers"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal server error: while casting the card as int list : %s", err)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, string(errorMsg))
		return
	}
	isValid := algorithm.LuhnAlgorithm(cardNumAsIntList)

	// Build the output and send it
	outputJson := OutputJson{
		CardNumber: inputJson.CardNumber,
		IsValid:    isValid,
		Date:       time.Now(),
	}
	output, mapErr := MapOutput(outputJson)
	if mapErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error: while mapping the output : %s", mapErr)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, string(output))
}

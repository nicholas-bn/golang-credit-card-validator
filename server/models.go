package server

import (
	"encoding/json"
	"fmt"
	"time"
)

type InputJson struct {
	CardNumber string `json:"card_number"`
}

type OutputJson struct {
	CardNumber string    `json:"card_number"`
	IsValid    bool      `json:"is_valid"`
	Date       time.Time `json:"date"`
}

type ErrorJson struct {
	ErrorMessage string `json:"error"`
	StatusCode   int    `json:"status_code"`
}

func MapInput(input []byte, data *InputJson) error {
	err := json.Unmarshal(input, data)
	if err != nil {
		return err
	}
	if data.CardNumber == "" {
		return fmt.Errorf("Error while doing the JSON unmarshal, no `card_number` given")
	}
	return nil
}

func MapOutput(data OutputJson) ([]byte, error) {
	return json.Marshal(data)
}

func MapError(data ErrorJson) ([]byte, error) {
	return json.Marshal(data)
}

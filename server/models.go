package server

import "time"

type InputJson struct {
	CardNumber string `json:"card_number"`
}

type OutputJson struct {
	CardNumber string    `json:"card_number"`
	IsValid    bool      `json:"is_valid"`
	Date       time.Time `json:"date"`
}

type ErrorJson struct {
	Error      error  `json:"error,omitempty"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

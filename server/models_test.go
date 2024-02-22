package server

import (
	"encoding/json"
	"log"
	"slices"
	"testing"
	"time"
)

func TestMapInputOK(t *testing.T) {

	str := []byte("{\"card_number\":\"4556737586899855\"}")
	got := InputJson{}

	err := json.Unmarshal(str, &got)
	if err != nil {
		log.Fatalf("impossible to unmarshall card input: %s", err)
	}
	want := InputJson{
		CardNumber: "4556737586899855",
	}

	if want != got {
		log.Fatalf("Mapping failed: got %s and want %s", got, want)
	}
}

func TestMapInputNOTOK(t *testing.T) {

	str := []byte("{\"WRONG_JSON\":\"121\"}")
	got := InputJson{}

	err := MapInput(str, &got)
	if err == nil {
		log.Fatalf("Should have failed as no `credit_card` parameter is given")
	}
}

func TestMapOutput(t *testing.T) {

	start := OutputJson{
		CardNumber: "4556737586899855",
		IsValid:    true,
		Date:       time.Time{},
	}
	want := []byte("{\"card_number\":\"4556737586899855\",\"is_valid\":true,\"date\":\"0001-01-01T00:00:00Z\"}")
	got, err := MapOutput(start)
	if err != nil {
		log.Fatalf("impossible to marshall card output: %s", err)
	}

	if !slices.Equal(got, want) {
		log.Fatalf("Mapping failed: got %s and want %s", got, want)
	}
}

func TestMapError(t *testing.T) {
	e := ErrorJson{
		StatusCode:   404,
		ErrorMessage: "message",
	}
	got, err := MapError(e)
	if err != nil {
		log.Fatalf("impossible to marshall card output: %s", err)
	}
	want := []byte("{\"error\":\"message\",\"status_code\":404}")

	if !slices.Equal(got, want) {
		log.Fatalf("Mapping failed: got %s and want %s", got, want)
	}

}

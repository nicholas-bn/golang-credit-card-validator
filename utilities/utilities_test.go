package utilities

import (
	"slices"
	"testing"
)

func TestStringToIntListOK12345(t *testing.T) {
	in := "12345"
	want := []int{1, 2, 3, 4, 5}
	out, err := StringToIntList(in)
	if err != nil {
		t.Fatalf(`StringToIntList("%s") produced an unexpected error: %s`, in, err)
	}
	if !slices.Equal(want, out) {
		t.Fatalf(`StringToIntList(%s) = %q, %v, want match for %#q, nil`, in, out, err, want)
	}
}

func TestStringToIntListNOTOK12_345(t *testing.T) {
	in := "12_345"
	_, err := StringToIntList(in)
	if err == nil {
		t.Fatalf(`StringToIntList("%s") dit not fail as expected: %s`, in, err)
	}
}

func TestStringToIntListNOTOK12345withSpace(t *testing.T) {
	in := "12345 "
	_, err := StringToIntList(in)
	if err == nil {
		t.Fatalf(`StringToIntList("%s") dit not fail as expected: %s`, in, err)
	}
}

func TestStringToIntListOKEmpty(t *testing.T) {
	in := ""
	want := []int{}
	out, err := StringToIntList(in)
	if err != nil {
		t.Fatalf(`StringToIntList("%s") produced an unexpected error: %s`, in, err)
	}
	if !slices.Equal(want, out) {
		t.Fatalf(`Hello(%s) = %q, %v, want match for %#q, nil`, in, out, err, want)
	}
}

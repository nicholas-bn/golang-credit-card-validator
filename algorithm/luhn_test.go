package algorithm

import (
	"fmt"
	"testing"
)

func TestLuhnAlgorithmOK(t *testing.T) {
	dataset := [][]int{
		{4, 5, 5, 6, 7, 3, 7, 5, 8, 6, 8, 9, 9, 8, 5, 5},
		{5, 1, 4, 3, 5, 6, 3, 5, 7, 8, 7, 9, 4, 5, 3, 3},
		{6, 0, 1, 1, 4, 5, 0, 4, 4, 6, 0, 1, 6, 5, 3, 8},
		{6, 0, 1, 1, 3, 5, 3, 3, 1, 9, 5, 2, 3, 0, 7, 9},
		{6, 0, 1, 1, 6, 1, 3, 7, 9, 0, 6, 2, 4, 5, 7, 5},
		{4, 5, 3, 2, 0, 6, 1, 6, 6, 1, 0, 3, 7, 3, 3, 4},
		{6, 0, 1, 1, 3, 2, 0, 8, 9, 8, 9, 9, 2, 8, 2, 7},
	}
	for _, v := range dataset {
		t.Run(fmt.Sprintf("assert that %v is valid", v), func(t *testing.T) {
			if !LuhnAlgorithm(v) {
				t.Fatalf(`Hello(%v) = false, expected a true`, v)
			}
		})
	}
}

func TestLuhnAlgorithmNOTOK(t *testing.T) {
	dataset := [][]int{
		{2, 2, 2, 1, 0, 0, 4, 7, 4, 6, 8, 5, 5, 6, 4, 2},
		{2, 2, 2, 1, 0, 0, 8, 6, 7, 4, 6, 7, 3, 0, 0, 5},
		{4, 0, 2, 4, 0, 0, 7, 1, 0, 9, 0, 2, 2, 1, 4, 3},
		{4, 9, 1, 6, 9, 0, 4, 0, 0, 6, 2, 4, 3, 5, 6, 1},
		{4, 4, 8, 5, 9, 1, 1, 7, 2, 0, 9, 7, 0, 6, 4, 6},
		{4, 9, 1, 6, 8, 0, 7, 5, 8, 8, 8, 1, 2, 2, 8, 7},
	}
	for _, v := range dataset {
		t.Run(fmt.Sprintf("assert that %v is NOT valid", v), func(t *testing.T) {
			if LuhnAlgorithm(v) {
				t.Fatalf(`Hello(%v) = true, expected a false`, v)
			}
		})
	}
}

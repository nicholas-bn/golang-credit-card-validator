package main

import (
	"fmt"

	"github.com/nicholas-bn/golang-credit-card-validator/algorithm"
)

func main() {
	fmt.Println(algorithm.LuhnAlgorithm([]int{1, 9, 4, 1, 1, 0, 6, 0, 6, 9, 0, 9, 1, 2}))
	// fmt.Println(algorithm.LuhnAlgorithm([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}))
}

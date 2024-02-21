package algorithm

func LuhnAlgorithm(cardNumber []int) bool {

	var total int = 0

	parity := len(cardNumber) % 2
	for i, num := range cardNumber {
		if parity == i%2 {
			if num <= 4 {
				total += num * 2
			} else {
				total += (num * 2) - 9
			}
		} else {
			total += num
		}
	}

	return total%10 == 0
}

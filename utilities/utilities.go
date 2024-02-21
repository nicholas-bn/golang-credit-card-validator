package utilities

import (
	"fmt"
	"strconv"
)

func StringToIntList(s string) ([]int, error) {
	var list []int
	for _, runeChar := range s {
		val, err := strconv.Atoi(string(runeChar))
		if err != nil {
			return list, fmt.Errorf("StringToIntList: error while casting %s in a int list : %s", s, err)
		}
		list = append(list, val)
	}
	return list, nil
}

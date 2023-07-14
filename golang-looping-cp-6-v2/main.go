package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	str := strconv.Itoa(numbers)
	var temp int
	var final int
	for i := 1; i < len(str); i++ {
		temp1, _ := strconv.Atoi(string(str[i-1]))
		temp2, _ := strconv.Atoi(string(str[i]))
		if temp1+temp2 > temp {
			temp = temp1 + temp2
			final, _ = strconv.Atoi(string(str[i-1]) + string(str[i]))
		}
	}
	return final // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}

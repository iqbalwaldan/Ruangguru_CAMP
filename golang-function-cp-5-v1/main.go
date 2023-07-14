package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	var numMin int
	first := true
	for _, min := range nums {
		if first {
			numMin = min
			first = false
		} else if numMin > min {
			numMin = min
		}
	}
	return numMin
}

func FindMax(nums ...int) int {
	var numMax int
	first := true
	for _, max := range nums {
		if first {
			numMax = max
			first = false
		} else if numMax < max {
			numMax = max
		}
	}
	return numMax
}

func SumMinMax(nums ...int) int {
	return FindMin(nums...) + FindMax(nums...)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// Case 1
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// Case 2
	fmt.Println(SumMinMax(333, 456, 654, 123, 111, 1000, 1500, 2000, 3000, 1250, 1111))
}

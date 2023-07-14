package main

import (
	"fmt"
	"strconv"
)

func ReverseData(arr [5]int) [5]int {
	reverse := [5]int{}
	for i := 0; i < len(arr); i++ {
		temp := ReverseNumber(arr[len(arr)-1-i])
		reverse[i] = temp
	}
	// --- Cara 2
	// for i, rev := range arr {
	// 	reverse[len(arr)-1-i] = rev
	// }
	// --- Cara 1 ---
	// for i := len(arr) - 1; i >= 0; i-- {
	// 	for j := len(arr) - 1 - i; j < len(reverse); j++ {
	// 		reverse[j] = arr[i]
	// 	}
	// }
	return reverse
}

func ReverseNumber(numbers int) int {
	number := strconv.Itoa(numbers)

	reverse := ""
	for _, rev := range number {
		reverse = string(rev) + reverse
	}
	final, _ := strconv.Atoi(reverse)
	return final
	// TODO: replace this
}

func main() {
	fmt.Println(ReverseData([5]int{123, 456, 11, 1, 2}))
	fmt.Println(ReverseData([5]int{456789, 44332, 2221, 12, 10}))
	fmt.Println(ReverseData([5]int{10, 10, 10, 10, 10}))
	fmt.Println(ReverseData([5]int{23456, 789, 123, 456, 500}))
	fmt.Println(ReverseData([5]int{0, 0, 0, 0, 0}))
}

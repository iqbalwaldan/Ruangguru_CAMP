package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {
	final := []int{}
	for _, dd1 := range date1 {
		for _, dd2 := range date2 {
			if dd2 == dd1 {
				final = append(final, dd2)
			}
		}
	}
	return final // TODO: replace this
}

func main() {
	fmt.Println(SchedulableDays([]int{1, 2, 3, 4}, []int{3, 4, 5}))
	fmt.Println(SchedulableDays([]int{11, 12, 13, 14, 15}, []int{5, 10, 12, 13, 20, 21}))
	fmt.Println(SchedulableDays([]int{2, 7, 12, 20, 21, 22}, []int{1, 3, 6, 10}))
	fmt.Println(SchedulableDays([]int{}, []int{}))
}

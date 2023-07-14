package main

import (
	"fmt"
)

func SchedulableDays(villager [][]int) []int {
	avail := make(map[int]int)
	for _, dim1 := range villager {
		for _, dim2 := range dim1 {
			avail[dim2] = avail[dim2] + 1
		}
	}
	final := []int{}
	for i, a := range avail {
		if a == len(villager) {
			final = append(final, i)
		}
	}
	return final
}

func main() {
	fmt.Println(SchedulableDays([][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}))
	fmt.Println(SchedulableDays([][]int{{1, 2, 3, 4, 5}, {2, 3, 4, 5}, {2, 3, 4, 10, 11, 12, 15}}))
	fmt.Println(SchedulableDays([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {10, 11, 12}, {21, 22, 23, 24}, {25}}))
	fmt.Println(SchedulableDays([][]int{}))

}

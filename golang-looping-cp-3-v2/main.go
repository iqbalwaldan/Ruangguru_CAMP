package main

import "fmt"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	var count int
	for i := 0; i < len(text); i++ {
		switch string(text[i]) {
		case "R", "r", "S", "s", "T", "t", "Z", "z":
			count++
		}
	}
	return count
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
	// Case 1
	fmt.Println(CountingLetter("Remaja muda yang berbakat"))
	// Case 2
	fmt.Println(CountingLetter("Zebra Zig Zag"))
}

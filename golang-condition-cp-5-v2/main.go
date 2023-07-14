package main

import (
	"fmt"
)

func TicketPlayground(height, age int) int {
	// if age > 12 {
	// 	return 100000
	// } else if age == 12 || height > 160 {
	// 	return 60000
	// } else if age >= 10 || height > 150 {
	// 	return 40000
	// } else if age >= 8 || height > 135 {
	// 	return 25000
	// } else if age >= 5 || height > 120 {
	// 	return 15000
	// } else {
	// 	return -1
	// }

	switch {
	case age > 12:
		return 100000
	case age == 12 || height > 160:
		return 60000
	case age >= 10 || height > 150:
		return 40000
	case age >= 8 || height > 135:
		return 25000
	case age >= 5 || height > 120:
		return 15000
	default:
		return -1
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
	// Case 1
	fmt.Println(TicketPlayground(160, 11))
	// Case 2
	fmt.Println(TicketPlayground(165, 10))
}

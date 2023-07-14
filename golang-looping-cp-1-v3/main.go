package main

import "fmt"

func CountingNumber(n int) float64 {
	var a float64
	b := float64(n)
	for i := 1.0; i <= b; i += 0.5 {
		a += i
	}
	return a
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
	// Case 1
	fmt.Println(CountingNumber(10))
	// Case 2
	fmt.Println(CountingNumber(100))
}

package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	heightFloat := float64(height)
	if gender == "laki-laki" {
		return (heightFloat - 100) - ((heightFloat - 100) * 10 / 100)
	} else {
		return (heightFloat - 100) - ((heightFloat - 100) * 15 / 100)
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("perempuan", 165))
	// Case 1
	fmt.Println(BMICalculator("laki-laki", 170))
	// Case 2
	fmt.Println(BMICalculator("perempuan", 165))
}

package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	score := float32((math + science + english + indonesia) / 4)
	switch {
	case score == 100:
		return "Sempurna"
	case score >= 90:
		return "Sangat Baik"
	case score >= 80:
		return "Baik"
	case score >= 70:
		return "Cukup"
	case score >= 60:
		return "Kurang"
	default:
		return "Sangat kurang"
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
	// Case 1
	fmt.Println(GetPredicate(50, 80, 100, 60))
	// Case 2
	fmt.Println(GetPredicate(100, 100, 100, 100))
}

package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	if x, y := score, absent; x >= 70 && y < 5 {
		return "lulus"
	} else {
		return "tidak lulus"
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(70, 4))
	// Case 1
	fmt.Println(GraduateStudent(100, 4))
	// Case 2
	fmt.Println(GraduateStudent(80, 5))
}

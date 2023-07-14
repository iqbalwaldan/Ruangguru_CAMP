package main

import (
	"fmt"
)

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var reverse string
	for i := len(str) - 1; i >= 0; i-- {
		if i-1 >= 0 && string(str[i-1]) == " " || string(str[i]) == " " {
			reverse += string(str[i])
		} else if i == 0 {
			reverse += string(str[i])
		} else {
			reverse += string(str[i])
			reverse += "_"
		}
	}
	return reverse
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
	// Case 1
	fmt.Println(ReverseString("Hello World"))
	// Case 2
	fmt.Println(ReverseString("I am a student"))
}

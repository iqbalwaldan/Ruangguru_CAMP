package main

import (
	"fmt"
	"strconv"
)

func PhoneNumberChecker(number string, result *string) {
	if number[0:2] == "08" {
		number = "62" + number[1:]
	}
	if number[0:3] == "628" {
		*result = "invalid"
	}
	if len(number) < 11 {
		*result = "invalid"
	} else if len(number) >= 11 {
		num, _ := strconv.Atoi(number[0:5])
		switch {
		case num >= 62811 && num <= 62815:
			*result = "Telkomsel"
		case num >= 62816 && num <= 62819:
			*result = "Indosat"
		case num >= 62821 && num <= 62823:
			*result = "XL"
		case num >= 62827 && num <= 62829:
			*result = "Tri"
		case num >= 62852 && num <= 62853:
			*result = "AS"
		case num >= 62881 && num <= 62888:
			*result = "Smartfren"
		default:
			*result = "invalid"
		}
	}
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "081234567890"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}

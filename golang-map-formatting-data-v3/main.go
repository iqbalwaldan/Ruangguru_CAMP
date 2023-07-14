package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	var result = make(map[string][]string)
	for _, values := range data {
		format := strings.Split(values, "-")
		keys := format[0]
		indexs, _ := strconv.Atoi(format[1])
		possition := format[2]
		value := format[3]
		if _, ok := result[keys]; !ok {
			result[keys] = make([]string, 0)
		}
		if possition == "first" {
			if indexs >= len(result[keys]) {
				result[keys] = append(result[keys], value)
			} else {
				result[keys][indexs] = value + result[keys][indexs]
			}
		} else {
			if indexs >= len(result[keys]) {
				result[keys] = append(result[keys], value)
			} else {
				result[keys][indexs] = result[keys][indexs] + " " + value
			}
		}
	}
	return result
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}

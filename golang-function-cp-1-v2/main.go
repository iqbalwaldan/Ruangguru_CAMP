package main

import (
	"fmt"
	"strconv"
)

func DateFormat(day, month, year int) string {
	var foramtDay string
	if len(strconv.Itoa(day)) == 1 {
		foramtDay = "0" + strconv.Itoa(day)
	} else {
		foramtDay = strconv.Itoa(day)
	}
	var formatMonth string
	switch month {
	case 1:
		formatMonth = "January"
	case 2:
		formatMonth = "February"
	case 3:
		formatMonth = "March"
	case 4:
		formatMonth = "April"
	case 5:
		formatMonth = "May"
	case 6:
		formatMonth = "June"
	case 7:
		formatMonth = "July"
	case 8:
		formatMonth = "August"
	case 9:
		formatMonth = "September"
	case 10:
		formatMonth = "October"
	case 11:
		formatMonth = "November"
	case 12:
		formatMonth = "December"
	}

	foramtYear := strconv.Itoa(year)
	return foramtDay + "-" + formatMonth + "-" + foramtYear // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
	// Case 1
	fmt.Println(DateFormat(1, 1, 2020))
	// Case 2
	fmt.Println(DateFormat(31, 12, 2020))

}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	hour := 0
	minute := 0
	note := "AM"
	var ok bool

	switch time.(type) {
	case string:
		temp := strings.Split(time.(string), ":")
		if len(temp) != 2 || temp[1] == "" {
			return "Invalid input"
		}
		hour, _ = strconv.Atoi(temp[0])
		minute, _ = strconv.Atoi(temp[1])
	case []int:
		temp := time.([]int)
		if len(temp) == 1 {
			return "Invalid input"
		}
		hour = temp[0]
		minute = temp[1]
	case map[string]int:
		temp := time.(map[string]int)
		hour, ok = temp["hour"]
		minute, ok = temp["minute"]
		if !ok {
			return "Invalid input"
		}
	case Time:
		hour = time.(Time).Hour
		minute = time.(Time).Minute
	default:
		return "Invalid input"
	}
	if hour > 12 {
		note = "PM"
		hour -= 12
	} else if hour == 12 {
		note = "PM"
	}
	if hour == 0 {
		return "Invalid input"
	} else {
		return fmt.Sprintf("%02d:%02d %s", hour, minute, note)
	}
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}

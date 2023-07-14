package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	final := ""
	first := true
	for _, v := range data {
		containKey := strings.Contains(v, input)
		if containKey {
			if first {
				final += v
				first = false
			} else {
				final += "," + v
			}
		}
	}
	return final
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}

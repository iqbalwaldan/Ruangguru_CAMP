package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	var countVocal, countConsonant int
	var countBool bool
	for i := 0; i < len(str); i++ {
		char := strings.ToLower(string(str[i]))
		switch char {
		case "a", "i", "u", "e", "o":
			countVocal++
		default:
			if char >= "a" && char <= "z" {
				countConsonant++
			}
		}

	}
	if countConsonant == 0 || countVocal == 0 {
		countBool = true
	}
	return countVocal, countConsonant, countBool
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
	fmt.Println(CountVowelConsonant("kopi"))
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}

// 97 - 122

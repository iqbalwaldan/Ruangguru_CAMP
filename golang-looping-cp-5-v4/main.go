package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	var finalStr, wordStr string
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			wordStr += string(str[i])
		}
		if str[i] == ' ' || i == len(str)-1 {
			var reverse string
			for j := len(wordStr) - 1; j >= 0; j-- {
				reverse += string(wordStr[j])
			}
			if unicode.IsUpper(rune(wordStr[0])) == true {
				reverse = strings.ToUpper(string(reverse[0])) + reverse[1:]
			}
			if unicode.IsLower((rune(wordStr[len(wordStr)-1]))) == true {
				reverse = reverse[:len(reverse)-1] + strings.ToLower(string(reverse[len(reverse)-1]))
			}
			finalStr += reverse + " "
			wordStr = ""
		}
	}
	return finalStr[:len(finalStr)-1]

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
	fmt.Println(ReverseWord("KITA SELALU BERSAMA"))
}

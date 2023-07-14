package main

import "fmt"

func SlurredTalk(words *string) {
	// var word = *words
	wordFinal := ""
	for _, word := range *words {
		switch string(word) {
		case "S", "R", "Z":
			wordFinal += "L"
		case "s", "r", "z":
			wordFinal += "l"
		default:
			wordFinal += string(word)
		}
	}
	*words = wordFinal
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Steven"
	SlurredTalk(&words)
	fmt.Println(words)
}

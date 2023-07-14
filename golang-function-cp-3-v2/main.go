package main

import (
	"fmt"
)

func ChoiceShortestName(name, nameShort string) string {
	if len(name) < len(nameShort) {
		return name
	} else if len(name) == len(nameShort) {
		if name < nameShort {
			return name
		} else {
			return nameShort
		}
	} else {
		return nameShort
	}

	// if len(name) < len(nameShort) {
	// 	return name
	// } else if len(name) > len(nameShort) {
	// 	return nameShort
	// } else {
	// 	if name < nameShort {
	// 		return name
	// 	} else {
	// 		return nameShort
	// 	}

	// }
}

func FindShortestName(names string) string {
	var name string
	nameShort := names
	for _, char := range names {
		if string(char) == " " || string(char) == "," || string(char) == ";" {
			// if len(name) < len(nameShort) {
			// 	nameShort = name
			// } else if len(name) == len(nameShort) {
			// 	if name < nameShort {
			// 		nameShort = name
			// 	}
			// }
			nameShort = ChoiceShortestName(name, nameShort)
			name = ""
		} else {
			name += string(char)
		}
	}
	fmt.Println(name)
	if name != "" {
		// if len(name) < len(nameShort) {
		// 	nameShort = name
		// } else if len(name) == len(nameShort) {
		// 	if name < nameShort {
		// 		nameShort = name
		// 	}
		// }
		nameShort = ChoiceShortestName(name, nameShort)
	}
	return nameShort
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}

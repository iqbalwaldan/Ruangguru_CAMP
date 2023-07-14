package main

import "fmt"

func EmailInfo(email string) string {
	var domain, TLD string
	for i := 0; i < len(email); i++ {

		if string(email[i]) == "@" {
			for j := i + 1; j < len(email); j++ {
				domain += string(email[j])
				if string(email[j+1]) == "." {
					for k := j + 2; k < len(email); k++ {
						TLD += string(email[k])
					}
					break
				}
			}
		}
	}
	return "Domain: " + domain + " dan TLD: " + TLD // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
	// Case 1
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
	// Case 2
	fmt.Println(EmailInfo("ptmencaricintasejati@gmail.co.id"))
}

package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	totalTiket := VIP + regular + student
	VIP *= 30
	regular *= 20
	student *= 10
	totalHargaTiket := float32(VIP + regular + student)
	if totalHargaTiket >= 100 {
		if day%2 == 1 {
			if totalTiket < 5 {
				return totalHargaTiket - (totalHargaTiket * 15 / 100)
			} else {
				return totalHargaTiket - (totalHargaTiket * 25 / 100)
			}
		} else {
			if totalTiket < 5 {
				return totalHargaTiket - (totalHargaTiket * 10 / 100)
			} else {
				return totalHargaTiket - (totalHargaTiket * 20 / 100)
			}
		}
	} else {
		return totalHargaTiket
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
	// Case 1
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
	// Case 2
	fmt.Println(GetTicketPrice(3, 3, 3, 20))
}

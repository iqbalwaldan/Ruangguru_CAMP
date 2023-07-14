package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {
	result := make(map[string]float32)
	for _, d := range data {
		splitData := strings.Split(d, ":")
		name := splitData[0] + "-" + splitData[1]
		price, _ := strconv.Atoi(splitData[2])
		total := float32(price)
		destination := splitData[3]
		// Menambahkan biaya admin
		switch day {
		case "senin", "rabu", "jumat":
			total += total * 10 / 100
		case "selasa", "kamis", "sabtu":
			total += total * 5 / 100
		}
		// Menentukan Pengiriman yang dilakukan
		switch day {
		case "senin":
			switch destination {
			case "JKT", "DPK":
				result[name] = total
			}
		case "selasa":
			switch destination {
			case "JKT", "DPK", "BKS":
				result[name] = total
			}
		case "rabu":
			switch destination {
			case "JKT", "BDG":
				result[name] = total
			}
		case "kamis":
			switch destination {
			case "JKT", "BDG", "BKS":
				result[name] = total
			}
		case "jumat":
			switch destination {
			case "JKT", "BKS":
				result[name] = total
			}
		case "sabtu":
			switch destination {
			case "JKT", "BDG":
				result[name] = total
			}
		}
	}
	return result // TODO: replace this
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}

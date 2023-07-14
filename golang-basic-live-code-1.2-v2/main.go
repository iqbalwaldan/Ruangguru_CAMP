package main

import (
	"fmt"
)

func main() {
	for {
		var panjang float64
		var dari, ke string

		fmt.Println("=== Kalkulator Konversi Satuan Panjang ===")

		fmt.Print("Masukkan panjang: ")
		fmt.Scanln(&panjang)

		fmt.Print("Masukkan satuan dari (m/cm/ft/in): ")
		fmt.Scanln(&dari)

		fmt.Print("Masukkan satuan ke (m/cm/ft/in): ")
		fmt.Scanln(&ke)

		result := ConvertLength(panjang, dari, ke)
		fmt.Printf("%.2f %s = %.2f %s\n", panjang, dari, result, ke)

		var pilihan string
		fmt.Print("Apakah Anda ingin mengkonversi kembali? (y/n): ")
		fmt.Scanln(&pilihan)

		if pilihan == "n" {
			break
		}
	}
}

func ConvertLength(panjang float64, dari, ke string) float64 {
	if panjang <= 0 {
		return 0
	}

	switch dari {
	case "cm":
		switch ke {
		case "cm":
			return panjang
		case "m":
			return panjang / 100
		case "ft":
			return panjang / 30.48
		case "in":
			return panjang / 2.54
		default:
			return panjang
		}
	case "m":
		switch ke {
		case "cm":
			return panjang * 100
		case "m":
			return panjang
		case "ft":
			return panjang * 3.281
		case "in":
			return panjang * 39.37
		default:
			return panjang
		}
	case "ft":
		switch ke {
		case "cm":
			return panjang * 30.48
		case "m":
			return panjang / 3.281
		case "ft":
			return panjang
		case "in":
			return panjang * 12
		default:
			return panjang
		}
	case "in":
		switch ke {
		case "cm":
			return panjang * 2.54
		case "m":
			return panjang / 39.37
		case "ft":
			return panjang / 12
		case "in":
			return panjang
		default:
			return panjang
		}
	default:
		return panjang
	}
}

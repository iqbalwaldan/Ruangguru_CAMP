package main

import (
	"fmt"
)

func main() {
	for {
		var bentuk string
		var sisi, alas, tinggi float64

		fmt.Println("=== Kalkulator Geometri ===")

		fmt.Print("Masukkan bentuk geometri (persegi/persegi-panjang): ")
		fmt.Scanln(&bentuk)

		switch SelectForm(bentuk) {
		case "persegi":
			fmt.Print("Masukkan sisi: ")
			fmt.Scanln(&sisi)

			resultLuas, resultKeliling, err := CalculateSquare(sisi)
			if err != "" {
				fmt.Println(err)
			}

			fmt.Printf("Luas persegi: %.2f\n", resultLuas)
			fmt.Printf("Keliling persegi: %.2f\n", resultKeliling)
		case "persegi-panjang":
			fmt.Print("Masukkan panjang: ")
			fmt.Scanln(&alas)

			fmt.Print("Masukkan lebar: ")
			fmt.Scanln(&tinggi)

			resultLuas, resultKeliling, err := CalculateRectangle(alas, tinggi)
			if err != "" {
				fmt.Println(err)
			}
			fmt.Printf("Luas persegi panjang: %.2f\n", resultLuas)
			fmt.Printf("Keliling persegi panjang: %.2f\n", resultKeliling)
		default:
			fmt.Println("Bentuk geometri tidak valid!")
		}

		var pilihan string
		fmt.Print("Apakah Anda ingin menghitung lagi? (y/n): ")
		fmt.Scanln(&pilihan)

		if pilihan == "n" {
			break
		}
	}
}

func SelectForm(bentuk string) string {
	if bentuk == "persegi" || bentuk == "persegi-panjang" {
		return bentuk
	} else {
		return "Bentuk geometri tidak valid!"
	}
}

func CalculateSquare(sisi float64) (float64, float64, string) {
	if sisi <= 0 {
		return 0, 0, "sisi harus lebih besar dari 0"
	}
	return sisi * sisi, 4 * sisi, "" // TODO: replace this
}

func CalculateRectangle(panjang, lebar float64) (float64, float64, string) {
	if panjang <= 0 || lebar <= 0 {
		return 0, 0, "panjang dan lebar harus lebih besar dari 0"
	}
	return panjang * lebar, 2 * (panjang + lebar), "" // TODO: replace this
}

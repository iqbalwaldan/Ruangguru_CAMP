package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func Coin(amount int) []int {
	result := make([]int, 0)
	for amount != 0 {
		switch {
		case amount >= 1000:
			result = append(result, 1000)
			amount -= 1000
		case amount >= 500:
			result = append(result, 500)
			amount -= 500
		case amount >= 200:
			result = append(result, 200)
			amount -= 200
		case amount >= 100:
			result = append(result, 100)
			amount -= 100
		case amount >= 50:
			result = append(result, 50)
			amount -= 50
		case amount >= 20:
			result = append(result, 20)
			amount -= 20
		case amount >= 10:
			result = append(result, 10)
			amount -= 10
		case amount >= 5:
			result = append(result, 5)
			amount -= 5
		default:
			result = append(result, 1)
			amount -= 1
		}
	}
	return result
}

func MoneyChanges(amount int, products []Product) []int {
	cost := 0
	for _, product := range products {
		cost += product.Price + product.Tax
	}
	return Coin(amount - cost)
}

func main() {
	products := []Product{
		{Name: "Teh", Price: 2000, Tax: 0},
		{Name: "Kopi", Price: 3000, Tax: 500},
		{Name: "Susu", Price: 10000, Tax: 0},
	}
	fmt.Println(MoneyChanges(20000, products))
}

package main

import "fmt"

func ExchangeCoin(amount int) []int {
	coins := []int{}
	coin := amount
	for coin != 0 {
		if coin >= 1000 {
			coin -= 1000
			coins = append(coins, 1000)
		} else if coin >= 500 {
			coin -= 500
			coins = append(coins, 500)
		} else if coin >= 200 {
			coin -= 200
			coins = append(coins, 200)
		} else if coin >= 100 {
			coin -= 100
			coins = append(coins, 100)
		} else if coin >= 50 {
			coin -= 50
			coins = append(coins, 50)
		} else if coin >= 20 {
			coin -= 20
			coins = append(coins, 20)
		} else if coin >= 10 {
			coin -= 10
			coins = append(coins, 10)
		} else if coin >= 5 {
			coin -= 5
			coins = append(coins, 5)
		} else {
			coin -= 1
			coins = append(coins, 1)
		}
	}
	return coins
}

func main() {
	fmt.Println(ExchangeCoin(1752))
	fmt.Println(ExchangeCoin(5000))
	fmt.Println(ExchangeCoin(1234))
	fmt.Println(ExchangeCoin(0))
}

package main

import (
	"fmt"
)

func CountProfit(data [][][2]int) []int {
	provitMap := make(map[int]int)
	for _, cabang := range data {
		// fmt.Println("Cabang", cabang)
		for i, bulan := range cabang {
			// fmt.Println("Bulan", bulan)
			profit := bulan[0] - bulan[1]
			provitMap[i+1] += profit
		}
	}
	// fmt.Println(provitMap)
	numberOfBulan := 0
	for k := range provitMap {
		// fmt.Println("numOfBulan", k)
		if k >= numberOfBulan {
			numberOfBulan = k
		}
	}
	result := make([]int, numberOfBulan)
	for bulan, profit := range provitMap {
		result[bulan-1] = profit
	}
	return result
}

func main() {
	fmt.Println(CountProfit(
		[][][2]int{
			{
				{
					1000, //Pendapatan dari bulan 1 di cabang 1
					500,  //Pengeluaran dari bulan 1 di cabang 1
				}, //Bulan 1 di Cabang 1
				{
					500,
					200,
				},
			}, //Cabang 1
			{
				{
					1200,
					200,
				},
				{
					1000,
					800,
				},
			},
			{
				{
					500,
					100,
				},
				{
					700,
					100,
				},
			},
		},
	))
	fmt.Println(CountProfit([][][2]int{{{1000, 800}, {700, 500}}, {{1000, 800}, {900, 200}}}))
	fmt.Println(CountProfit([][][2]int{{{1000, 500}, {500, 150}, {600, 100}, {800, 750}}}))
	fmt.Println(CountProfit([][][2]int{{{1000, 200}}, {{500, 100}}, {{600, 100}}, {{450, 150}}, {{100, 50}}}))
	fmt.Println(CountProfit([][][2]int{}))
}

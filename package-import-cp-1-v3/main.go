package main

import (
	"fmt"
	"strconv"
	"strings"

	"a21hc3NpZ25tZW50/internal"
)

func AdvanceCalculator(calculate string) float32 {
	data := strings.Split(calculate, " ")
	if len(data) == 0 {
		return 0
	}
	result, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return 0
	}
	calc := internal.NewCalculator(float32(result))
	for i := 1; i < len(data); i++ {
		if i%2 == 1 {
			number, err := strconv.ParseFloat(data[i+1], 32)
			if err != nil {
				return 0
			}
			switch {
			case data[i] == "+":
				calc.Add(float32(number))
			case data[i] == "-":
				calc.Subtract(float32(number))
			case data[i] == "*":
				calc.Multiply(float32(number))
			case data[i] == "/":
				calc.Divide(float32(number))
			}
		}
	}
	return calc.Result()
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}

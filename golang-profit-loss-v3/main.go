package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data := strings.Split(string(content), "\n")

	if len(data) == 1 && data[0] == "" || len(data) == 0 {
		return []string{}, nil
	}
	return data, nil
}

func CalculateProfitLoss(data []string) string {
	prof := 0
	lastDate := ""
	for _, line := range data {
		current := strings.Split(line, ";")
		date := current[0]
		lastDate = date
		trxType := current[1]
		amount, _ := strconv.Atoi(current[2])
		if trxType == "income" {
			prof += amount
		} else if trxType == "expense" {
			prof -= amount
		}
	}
	if prof >= 0 {
		return fmt.Sprintf("%s;profit;%d", lastDate, prof)
	} else {
		return fmt.Sprintf("%s;loss;%d", lastDate, -prof)
	}
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}

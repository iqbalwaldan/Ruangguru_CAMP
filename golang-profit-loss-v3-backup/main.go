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

	if string(content) == "" {
		return []string{}, nil
	}
	strArr := strings.Split(string(content), "\n")
	return strArr, nil
}

func CalculateProfitLoss(data []string) string {
	prof := 0
	lastDate := ""
	for _, line := range data {
		tokens := strings.Split(line, ";")
		date := tokens[0]
		lastDate = date
		trxType := tokens[1]
		amount, _ := strconv.Atoi(tokens[2])
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

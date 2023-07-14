package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func moneyIncomeOrExpense(output *[]string, currentDate string, money int) {
	if money < 0 {
		*output = append(*output, fmt.Sprintf("%s;expense;%d", currentDate, -money))
	} else {
		*output = append(*output, fmt.Sprintf("%s;income;%d", currentDate, money))
	}
}

func trxType(trxType string, trxAmount int, money *int) {
	if trxType == "income" {
		*money += trxAmount
	} else {
		*money -= trxAmount
	}
}

func RecordTransactions(path string, transactions []Transaction) error {
	if len(transactions) == 0 {
		return nil
	}
	sort.Slice(transactions, func(i, j int) bool {
		if transactions[i].Date > transactions[j].Date {
			return false
		} else {
			return true
		}
	})

	currentDate := transactions[0].Date
	money := 0
	output := make([]string, 0)

	for _, trx := range transactions {
		if trx.Date == currentDate {
			trxType(trx.Type, trx.Amount, &money)
		} else {
			moneyIncomeOrExpense(&output, currentDate, money)
			money = 0
			trxType(trx.Type, trx.Amount, &money)
			currentDate = trx.Date
		}
	}
	moneyIncomeOrExpense(&output, currentDate, money)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err2 := file.WriteString(strings.Join(output, "\n"))
	if err != nil {
		return err2
	}
	return nil
}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"01/01/2021", "expense", 30000},
		{"01/01/2021", "income", 20000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}

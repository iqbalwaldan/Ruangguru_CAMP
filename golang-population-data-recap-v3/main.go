package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	result := []map[string]any{}
	for _, d := range data {
		currentResult := make(map[string]interface{})
		mark := strings.Split(d, ";")
		currentResult["name"] = mark[0]
		currentResult["age"], _ = strconv.Atoi(mark[1])
		currentResult["address"] = mark[2]
		if mark[3] != "" {
			currentResult["height"], _ = strconv.ParseFloat(mark[3], 64)
		}
		if mark[4] != "" {
			currentResult["isMarried"], _ = strconv.ParseBool(mark[4])
		}
		result = append(result, currentResult)
	}
	return result // TODO: replace this
}

func main() {
	data := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	fmt.Println(PopulationData(data))
}

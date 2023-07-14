package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	result := [][]string{}
	for i, md := range mapData {
		data := []string{i, md}
		result = append(result, data)
	}
	return result
}

func main() {
	fmt.Println(MapToSlice(map[string]string{
		"2AAmyxdNoH": "1Tyd4YMWS9YtNDl",
		"6Y9TOxmnRW": "I55yYyPyLdNTihp",
		"6wxZPqx1w9": "ONcMRrQkf75W4Rp",
		"FJV7VgatQ1": "a24OaFlDFUkxIlu",
		"RsvJZqcAAg": "sC2uoBmIU5giCT0",
		"UjO4vEuYCy": "XapD94SxiqyC95e",
		"apv840AE9y": "1yFfXJAPMJlw5Zz",
		"qcaY750D1A": "4EXtN4Z2dnmX3sp",
		"tb5RPg3twh": "I9FFuQnxRmtcgZq",
		"wDIHXkA2fI": "65zsfdNYPr0wsTv",
	}))
}

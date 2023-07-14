package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Study struct {
	StudyName   string `json:"study_name"`
	StudyCredit int    `json:"study_credit"`
	Grade       string `json:"grade"`
}

type Report struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return Report{}, err
	}
	report := Report{}
	err = json.Unmarshal(jsonData, &report)
	if err != nil {
		return Report{}, err
	}
	fmt.Println(report)
	return report, nil
}

func GradePoint(report Report) float64 {
	return 0.0 // TODO: replace this
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Study struct {
	StudyName   string `json:"study_name"`
	StudyCredit int    `json:"study_credit"`
	Grade       string `json:"grade"`
}

type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

func ReadJSON(filename string) (Report, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Report{}, err
	}
	defer file.Close()
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return Report{}, err
	}
	report := Report{}
	err = json.Unmarshal(jsonData, &report)
	if err != nil {
		return Report{}, err
	}
	return report, nil
}

func GradePoint(report Report) float64 {
	ip := 0.0
	credit := 0
	if len(report.Studies) == 0 {
		return 0
	}
	for _, study := range report.Studies {
		grade := 0.0
		switch study.Grade {
		case "A":
			grade = 4.0
		case "AB":
			grade = 3.5
		case "B":
			grade = 3.0
		case "BC":
			grade = 2.5
		case "C":
			grade = 2.0
		case "CD":
			grade = 1.5
		case "D":
			grade = 1.0
		case "DE":
			grade = 0.5
		case "E":
			grade = 0.0
		default:
			return 0
		}
		ip += float64(study.StudyCredit) * grade
		credit += study.StudyCredit
		// Solution Video
		// grade := map[string]float64{
		// 	"A":  4.0,
		// 	"AB": 3.5,
		// 	"B":  3.0,
		// 	"BC": 2.5,
		// 	"C":  2.0,
		// 	"CD": 1.5,
		// 	"D":  1.0,
		// 	"DE": 0.5,
		// 	"E":  0.0,
		// }
		// ip += float64(study.StudyCredit) * grade[study.Grade]
		// credit += study.StudyCredit
	}
	return ip / float64(credit)
}

func main() {
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}

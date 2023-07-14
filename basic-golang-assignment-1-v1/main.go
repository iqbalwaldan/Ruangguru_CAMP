package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	var message string
	if id == "" || name == "" {
		message = "ID or Name is undefined!"
	} else if len(id) != 5 {
		message = "ID must be 5 characters long!"
	} else {
		var StudentsData string
		for i := 0; i < len(Students); i++ {
			if string(Students[i]) == " " {
				continue
			} else {
				StudentsData += string(Students[i])
			}
		}
		StudentsData += ","
		var logData string
		for _, data := range StudentsData {
			if string(data) == "," {
				split := strings.Split(logData, "_")
				if id == split[0] && name == split[1] {
					message = "Login berhasil: " + split[1] + " (" + split[2] + ")"
					break
				} else {
					message = "Login gagal: data mahasiswa tidak ditemukan"
				}
				logData = ""
			} else {
				logData += string(data)
			}
		}
	}
	return message // TODO: replace this
}

func Register(id string, name string, major string) string {
	var message string
	if id == "" || name == "" || major == "" {
		message = "ID, Name or Major is undefined!"
	} else if len(id) != 5 {
		message = "ID must be 5 characters long!"
	} else {
		var StudentsData string
		for i := 0; i < len(Students); i++ {
			if string(Students[i]) == " " {
				continue
			} else {
				StudentsData += string(Students[i])
			}
		}
		StudentsData += ","
		var logData string
		for _, data := range StudentsData {
			if string(data) == "," {
				split := strings.Split(logData, "_")
				if id == split[0] {
					message = "Registrasi gagal: id sudah digunakan"
					break
				} else if id != split[0] {
					message = "Registrasi berhasil: " + name + " (" + major + ")"
					Students += ", " + id + "_" + name + "_" + major
					break
				}
				logData = ""
			} else {
				logData += string(data)
			}
		}
	}
	return message // TODO: replace this
}

func GetStudyProgram(code string) string {
	var message string
	if code == "" {
		message = "Code is undefined!"
	} else {
		StudentStudyProgramsData := " " + StudentStudyPrograms + ","
		var logData string
		for _, data := range StudentStudyProgramsData {
			if string(data) == "," {
				split := strings.Split(logData, "_")
				if " "+code == split[0] {
					message = split[1]
					break
				} else {
					message = "Code is undefined!"
				}
				logData = ""
			} else {
				logData += string(data)
			}
		}
	}
	return message // TODO: replace this
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

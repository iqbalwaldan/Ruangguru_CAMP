package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"SI": "Sistem Informasi",
	"MI": "Manajemen Informasi",
}

type studentModifier func(string, *string)

func Login(id string, name string) string {
	var msg string
	if id == "" || name == "" {
		msg = "ID or Name is undefined!"
	} else {
		for _, s := range Students {
			splitStudent := strings.Split(s, "_")
			idStudent := splitStudent[0]
			nameStudent := splitStudent[1]
			if idStudent == id && nameStudent == name {
				msg = "Login berhasil: " + nameStudent
				break
			} else {
				msg = "Login gagal: data mahasiswa tidak ditemukan"

			}
		}
	}
	return msg
}

func Register(id string, name string, major string) string {
	var msg string
	if id == "" || name == "" || major == "" {
		msg = "ID, Name or Major is undefined!"
	} else {
		for _, s := range Students {
			splitStudent := strings.Split(s, "_")
			idStudent := splitStudent[0]
			if idStudent == id {
				msg = "Registrasi gagal: id sudah digunakan"
				break
			} else if id != idStudent {
				msg = "Registrasi berhasil: " + name + " (" + major + ")"
				Students = append(Students, id+"_"+name+"_"+major)
				break
			}
		}
	}
	return msg
}

func GetStudyProgram(code string) string {
	var msg string
	if _, ok := StudentStudyPrograms[code]; !ok {
		msg = "Kode program studi tidak ditemukan"
	} else {
		msg = StudentStudyPrograms[code]
	}
	return msg
}

func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	var msg string
	for i, s := range Students {
		splitStudent := strings.Split(s, "_")
		nameStudent := splitStudent[1]
		if nameStudent == nama {
			UpdateStudyProgram(programStudi, &Students[i])
			msg = "Program studi mahasiswa berhasil diubah."
			break
		} else {
			msg = "Mahasiswa tidak ditemukan."
		}
	}
	return msg
}

func UpdateStudyProgram(programStudi string, students *string) {
	splitData := strings.Split(*students, "_")
	*students = splitData[0] + "_" + splitData[1] + "_" + programStudi
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		for i, student := range Students {
			fmt.Println(i+1, student)
		}

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Change student study program")
		fmt.Println("5. Keluar")

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
			helper.ClearScreen()
			var nama, programStudi string
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan program studi baru: ")
			fmt.Scanln(&programStudi)

			fmt.Println(ModifyStudent(programStudi, nama, UpdateStudyProgram))
			helper.Delay(5)
		case "5":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

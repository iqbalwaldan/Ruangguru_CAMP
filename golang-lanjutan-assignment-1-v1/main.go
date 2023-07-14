package main

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetStudyProgram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			resp, err := json.Marshal(model.ErrorResponse{Error: "Method is not allowed!"})
			if err != nil {
				panic(err)
			}

			w.WriteHeader(405)
			_, err = w.Write(resp)
			if err != nil {
				panic(err)
			}
			return
		}

		content, err := os.ReadFile("data/list-study.txt")
		if err != nil {
			panic(err)
		}

		txt := string(content)
		txtPerbaris := strings.Split(txt, "\r\n")
		hasilJurusan := make([]model.StudyData, 0)

		for _, val := range txtPerbaris {
			jurusan := strings.Split(val, "_")
			hasilJurusan = append(hasilJurusan, model.StudyData{
				Code: jurusan[0],
				Name: jurusan[1],
			})
		}

		w.WriteHeader(200)
		resp, err := json.Marshal(hasilJurusan)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(resp)
		if err != nil {
			panic(err)
		}
		return
	}
}

func AddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			resp, err := json.Marshal(model.ErrorResponse{Error: "Method is not allowed!"})
			if err != nil {
				panic(err)
			}

			w.WriteHeader(405)
			_, err = w.Write(resp)
			if err != nil {
				panic(err)
			}
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		err = r.Body.Close()
		if err != nil {
			panic(err)
		}

		var user model.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			panic(err)
		}

		if user.ID == "" || user.Name == "" || user.StudyCode == "" {
			resp, err := json.Marshal(model.ErrorResponse{Error: "ID, name, or study code is empty"})
			if err != nil {
				panic(err)
			}

			w.WriteHeader(400)
			_, err = w.Write(resp)
			if err != nil {
				panic(err)
			}
			return
		}

		contentUser, err := os.ReadFile("data/users.txt")
		if err != nil {
			panic(err)
		}
		txtUser := string(contentUser)
		txtUserPerbaris := strings.Split(txtUser, "\n")

		contentJurusan, err := os.ReadFile("data/list-study.txt")
		if err != nil {
			panic(err)
		}
		txtJurusan := string(contentJurusan)
		txtJurusanPerbaris := strings.Split(txtJurusan, "\n")

		for _, val := range txtUserPerbaris {
			usr := strings.Split(val, "_")
			if len(usr) != 3 {
				continue
			}

			if usr[0] == user.ID {
				resp, err := json.Marshal(model.ErrorResponse{Error: "user id already exist"})
				if err != nil {
					panic(err)
				}

				w.WriteHeader(400)
				_, err = w.Write(resp)
				if err != nil {
					panic(err)
				}
				return
			}
		}

		jurusanAda := false
		for _, val := range txtJurusanPerbaris {
			jrs := strings.Split(val, "_")
			if user.StudyCode == jrs[0] {
				jurusanAda = true
				break
			}

		}

		if jurusanAda == false {
			resp, err := json.Marshal(model.ErrorResponse{Error: "study code not found"})
			if err != nil {
				panic(err)
			}

			w.WriteHeader(400)
			_, err = w.Write(resp)
			if err != nil {
				panic(err)
			}
			return
		}
		fileFinal := ""
		if txtUser == "" {
			fileFinal = string([]byte(user.ID + "_" + user.Name + "_" + user.StudyCode))

		} else {
			fileFinal = txtUser + "\n" + string([]byte(user.ID+"_"+user.Name+"_"+user.StudyCode))

		}

		err = os.WriteFile("data/users.txt", []byte(fileFinal), 0644)
		if err != nil {
			panic(err)
		}

		resp, err := json.Marshal(model.SuccessResponse{Username: user.ID, Message: "add user success"})
		if err != nil {
			panic(err)
		}

		w.WriteHeader(200)
		_, err = w.Write(resp)
		if err != nil {
			panic(err)
		}
		return
	}
}

func DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			resp, err := json.Marshal(model.ErrorResponse{Error: "Method is not allowed!"})
			if err != nil {
				panic(err)
			}

			w.WriteHeader(405)
			_, err = w.Write(resp)
			if err != nil {
				panic(err)
			}
			return
		}

		id := r.URL.Query().Get("id")

		if id == "" {
			resp, err := json.Marshal(model.ErrorResponse{Error: "user id is empty"})
			if err != nil {
				panic(err)
			}

			w.WriteHeader(400)
			_, err = w.Write(resp)
			if err != nil {
				panic(err)
			}
			return
		}

		contentUser, err := os.ReadFile("data/users.txt")
		if err != nil {
			panic(err)
		}
		txtUser := string(contentUser)
		txtUserPerbaris := strings.Split(txtUser, "\n")

		var idYangDihapus string
		for _, val := range txtUserPerbaris {
			usr := strings.Split(val, "_")
			if len(usr) != 3 {
				resp, err := json.Marshal(model.ErrorResponse{Error: "user id not found"})
				if err != nil {
					panic(err)
				}

				w.WriteHeader(400)
				_, err = w.Write(resp)
				if err != nil {
					panic(err)
				}
				return
			}

			if usr[0] == id {
				idYangDihapus = usr[0]
				break
			}

		}

		if idYangDihapus == "" {
			resp, err := json.Marshal(model.ErrorResponse{Error: "user id not found"})
			if err != nil {
				panic(err)
			}

			w.WriteHeader(400)
			_, err = w.Write(resp)
			if err != nil {
				panic(err)
			}
			return
		}

		var txtBaru string
		for _, val := range txtUserPerbaris {
			usr := strings.Split(val, "_")
			if len(usr) != 3 {
				continue
			}
			if usr[0] == idYangDihapus {
				continue
			}
			txtBaru = val + "\n"
		}
		if txtBaru != "" {
			txtBaru = txtBaru[0 : len(txtBaru)-1]

			err = os.WriteFile("data/users.txt", []byte(txtBaru), 0644)
			if err != nil {
				panic(err)
			}
		}

		resp, err := json.Marshal(model.SuccessResponse{Username: id, Message: "delete success"})
		if err != nil {
			panic(err)
		}

		w.WriteHeader(200)
		_, err = w.Write(resp)
		if err != nil {
			panic(err)
		}
		return
	}
}

func main() {
	http.HandleFunc("/study-program", GetStudyProgram())
	http.HandleFunc("/user/add", AddUser())
	http.HandleFunc("/user/delete", DeleteUser())

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

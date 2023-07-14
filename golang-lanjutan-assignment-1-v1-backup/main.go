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
			res, err := json.Marshal(model.ErrorResponse{Error: "Method is not allowed!"})
			if err != nil {
				panic(err)
			}
			w.WriteHeader(405)
			_, err = w.Write(res)
			if err != nil {
				panic(err)
			}
			return
		}

		data, err := Read("data/list-study.txt", "\r\n")
		if err != nil {
			panic(err)
		}

		stringJson := make([]any, 0)

		for _, study := range data {
			dataSplit := strings.Split(study, "_")
			idStudy := dataSplit[0]
			nameStudy := dataSplit[1]
			stringJson = append(stringJson, model.StudyData{Code: idStudy, Name: nameStudy})
		}
		study, err := json.Marshal(stringJson)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(study)
	}
}

func AddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			res, err := json.Marshal(model.ErrorResponse{Error: "Method is not allowed!"})
			if err != nil {
				panic(err)
			}
			w.WriteHeader(405)
			_, err = w.Write(res)
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

		var u model.User
		err = json.Unmarshal(body, &u)
		if err != nil {
			panic(err)
		}

		if u.ID == "" || u.Name == "" || u.StudyCode == "" {
			res, err := json.Marshal(model.ErrorResponse{Error: "ID, name, or study code is empty"})
			if err != nil {
				panic(err)
			}
			w.WriteHeader(400)
			_, err = w.Write(res)
			if err != nil {
				panic(err)
			}
			return
		}
		dataJurusan, err := Read("data/list-study.txt", "\n")
		if err != nil {
			panic(err)
		}
		// mapJurusan := make(map[string]string)
		// for _, jurusan := range dataJurusan {
		// 	strJurusan := strings.Split(jurusan, "_")
		// 	mapJurusan[strJurusan[0]] = strJurusan[1]
		// }

		dataUser, err := Read("data/users.txt", "\r\n")
		if err != nil {
			panic(err)
		}

		for _, user := range dataUser {
			strUser := strings.Split(user, "_")
			if len(user) != 3 {
				continue
			}
			if strUser[0] == u.ID {
				res, err := json.Marshal(model.ErrorResponse{Error: "user id already exist"})
				if err != nil {
					panic(err)
				}
				w.WriteHeader(400)
				_, err = w.Write(res)
				if err != nil {
					panic(err)
				}
				return
			}
		}

		exist := false
		for _, jrn := range dataJurusan {
			jrsn := strings.Split(jrn, "_")
			if u.StudyCode == jrsn[0] {
				exist = true
				break
			}
		}
		err = os.WriteFile("data/users.txt", []byte(fmt.Sprintf("%s_%s_%s\n", u.ID, u.Name, u.StudyCode)), 0644)
		if err != nil {
			panic(err)
		}

		if !exist {
			res, err := json.Marshal(model.ErrorResponse{Error: "study code not found"})
			if err != nil {
				panic(err)
			}
			w.WriteHeader(400)
			_, err = w.Write(res)
			if err != nil {
				panic(err)
			}
			return
		}

		res, err := json.Marshal(model.SuccessResponse{Username: u.ID, Message: "add user success"})
		if err != nil {
			panic(err)
		}

		w.WriteHeader(200)
		_, err = w.Write(res)
		if err != nil {
			panic(err)
		}
		return
	}
}

func DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			res, err := json.Marshal(model.ErrorResponse{Error: "Method is not allowed!"})
			if err != nil {
				panic(err)
			}
			w.WriteHeader(405)
			_, err = w.Write(res)
			if err != nil {
				panic(err)
			}
			return
		}

		id := r.URL.Query().Get("id")
		if id == "" {
			res, err := json.Marshal(model.ErrorResponse{Error: "user id is empty"})
			if err != nil {
				panic(err)
			}
			w.WriteHeader(400)
			_, err = w.Write(res)
			if err != nil {
				panic(err)
			}
			return
		}

		dataUser, err := Read("data/users.txt", "\n")
		if err != nil {
			panic(err)
		}

		var newFileUser string
		var delateId string
		// code := false

		for _, user := range dataUser {
			strUser := strings.Split(user, "_")
			if len(user) != 3 {
				// code = true
				res, err := json.Marshal(model.ErrorResponse{Error: "user id not found"})
				if err != nil {
					panic(err)
				}
				w.WriteHeader(400)
				_, err = w.Write(res)
				if err != nil {
					panic(err)
				}
				return
			}

			if strUser[0] == id {
				delateId = strUser[0]
				break
			}
		}

		if delateId == "" {
			res, err := json.Marshal(model.ErrorResponse{Error: "user id not found"})
			if err != nil {
				panic(err)
			}
			w.WriteHeader(400)
			_, err = w.Write(res)
			if err != nil {
				panic(err)
			}
			return
		}

		for _, user := range dataUser {
			strUser := strings.Split(user, "_")
			if len(user) != 3 {
				continue
			}
			if strUser[0] == delateId {
				continue
			}
			newFileUser = user + "\n"
		}

		if newFileUser != "" {
			newFileUser = newFileUser[:len(newFileUser)-1]
			err = os.WriteFile("data/users.txt", []byte(newFileUser), 0644)
			if err != nil {
				panic(err)
			}
		}
		// code = false
		// if code {
		// 	res, err := json.Marshal(model.ErrorResponse{Error: "user id not found"})
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	w.WriteHeader(400)
		// 	_, err = w.Write(res)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	return
		// } else {

		res, err := json.Marshal(model.SuccessResponse{Username: id, Message: "delete success"})
		if err != nil {
			panic(err)
		}

		w.WriteHeader(200)
		_, err = w.Write(res)
		if err != nil {
			panic(err)
		}
		return
	}
}

// }

func Read(path string, split string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data := strings.Split(string(content), split)

	if len(data) == 1 && data[0] == "" || len(data) == 0 {
		return []string{}, nil
	}
	return data, nil
}

func main() {
	http.HandleFunc("/study-program", GetStudyProgram())
	http.HandleFunc("/user/add", AddUser())
	http.HandleFunc("/user/delete", DeleteUser())

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

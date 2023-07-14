package handler

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/model"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var UserLogin = make(map[string]model.User)

// DESC: func Auth is a middleware to check user login id, only user that already login can pass this middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("user_login_id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		if _, ok := UserLogin[c.Value]; !ok || c.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user login id not found"})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", c.Value)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DESC: func AuthAdmin is a middleware to check user login role, only admin can pass this middleware
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		coki, err := r.Cookie("user_login_role")
		if err != nil {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user login role not Admin"})
			return
		}

		if coki.Value != "admin" {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user login role not Admin"})
			return
		}

		cont := r.Context()
		cont = context.WithValue(cont, "userRole", coki.Value)

		next.ServeHTTP(w, r.WithContext(cont))

	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	var user model.UserLogin
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	if user.ID == "" || user.Name == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "ID or name is empty"})
		return
	}

	data, exist := UserLogin[user.ID]
	if exist {
		http.SetCookie(w, &http.Cookie{Name: "user_login_id", Value: data.ID})
		http.SetCookie(w, &http.Cookie{Name: "user_login_role", Value: data.Role})

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(model.SuccessResponse{Username: user.ID, Message: "login success"})
	}

	txtUserLine, err := Read("data/users.txt", "\n")
	if err != nil {
		panic(err)
	}

	userExist := false
	var dataUserLogin model.User
	for _, val := range txtUserLine {
		usr := strings.Split(val, "_")
		if len(usr) != 4 {
			continue
		}
		if user.ID == usr[0] && user.Name == usr[1] {
			dataUserLogin = model.User{
				ID:        usr[0],
				Name:      usr[1],
				Role:      usr[2],
				StudyCode: usr[3],
			}
			userExist = true
			break
		}
	}

	if userExist == false {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user not found"})
		return
	}

	cookie := &http.Cookie{
		Name:  "user_login_id",
		Value: dataUserLogin.ID,
	}
	http.SetCookie(w, cookie)
	cookie = &http.Cookie{
		Name:  "user_login_role",
		Value: dataUserLogin.Role,
	}
	http.SetCookie(w, cookie)

	UserLogin[user.ID] = dataUserLogin

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(model.SuccessResponse{Username: user.ID, Message: "login success"})

	return

}

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	if user.ID == "" || user.Name == "" || user.StudyCode == "" || user.Role == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "ID, name, study code or role is empty"})
		return
	}

	role := false
	if user.Role == "admin" {
		role = true
	} else if user.Role == "user" {
		role = true
	}

	if role == false {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "role must be admin or user"})
		return
	}

	txtJurusanLine, err := Read("data/list-study.txt", "\n")
	if err != nil {
		panic(err)
	}

	jurusanExist := false
	for _, val := range txtJurusanLine {
		jrs := strings.Split(val, "_")
		if user.StudyCode == jrs[0] {
			jurusanExist = true
			break
		}

	}

	if jurusanExist == false {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "study code not found"})
		return
	}

	txtUserLine, err := Read("data/users.txt", "\n")
	if err != nil {
		panic(err)
	}

	for _, val := range txtUserLine {
		usr := strings.Split(val, "_")
		if len(usr) != 4 {
			continue
		}

		if usr[0] == user.ID {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user id already exist"})
			return
		}
	}

	fileFinal := ""
	txtUser, err := os.ReadFile("data/users.txt")
	txtUserFinal := fmt.Sprintf(string([]byte(user.ID + "_" + user.Name + "_" + user.StudyCode + "_" + user.Role)))
	if err != nil {
		panic(err)
	}
	if string(txtUser) == "" {
		fileFinal = txtUserFinal
	} else {
		fileFinal = string(txtUser) + "\n" + txtUserFinal
	}

	err = os.WriteFile("data/users.txt", []byte(fileFinal), 0644)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(model.SuccessResponse{Username: user.ID, Message: "register success"})
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	userLoginExist(userID, w)

	delete(UserLogin, userID)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(model.SuccessResponse{Username: userID, Message: "logout success"})
	return
}

func GetStudyProgram(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	userLoginExist(userID, w)

	txtJurusanLine, err := Read("data/list-study.txt", "\n")
	if err != nil {
		panic(err)
	}

	jurusanFinal := make([]model.StudyData, 0)

	for _, val := range txtJurusanLine {
		jurusan := strings.Split(val, "_")
		jurusanFinal = append(jurusanFinal, model.StudyData{Code: jurusan[0], Name: jurusan[1]})
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(jurusanFinal)
	return

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	userLoginExist(userID, w)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	if user.ID == "" || user.Name == "" || user.StudyCode == "" || user.Role == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "ID, name, study code or role is empty"})
		return
	}

	// role := false
	// if user.Role == "admin" {
	// 	role = true
	// } else if user.Role == "user" {
	// 	role = true
	// }

	// if role == false {
	// 	w.WriteHeader(400)
	// 	json.NewEncoder(w).Encode(model.ErrorResponse{Error: "role must be admin or user"})
	// 	return
	// }

	txtUserLine, err := Read("data/users.txt", "\n")
	if err != nil {
		panic(err)
	}

	txtJurusanLine, err := Read("data/list-study.txt", "\n")
	if err != nil {
		panic(err)
	}

	for _, val := range txtUserLine {
		usr := strings.Split(val, "_")
		if len(usr) != 3 {
			continue
		}

		if usr[0] == user.ID {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user id already exist"})
			return
		}
	}

	jurusanExist := false
	for _, val := range txtJurusanLine {
		jrs := strings.Split(val, "_")
		if user.StudyCode == jrs[0] {
			jurusanExist = true
			break
		}

	}

	if jurusanExist == false {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "study code not found"})
		return
	}

	fileFinal := ""
	txtUser, err := os.ReadFile("data/users.txt")
	txtUserFinal := fmt.Sprintf(string([]byte(user.ID + "_" + user.Name + "_" + user.StudyCode + "_" + user.Role)))
	if err != nil {
		panic(err)
	}
	if string(txtUser) == "" {
		fileFinal = txtUserFinal
	} else {
		fileFinal = string(txtUser) + "\n" + txtUserFinal
	}

	err = os.WriteFile("data/users.txt", []byte(fileFinal), 0644)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(model.SuccessResponse{Username: user.ID, Message: "add user success"})
	return

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here

}

func userLoginExist(id string, w http.ResponseWriter) {
	_, userLoginExist := UserLogin[id]
	if !userLoginExist {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user login id not found"})
		return
	}
}

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

// DESC: Gunakan variable ini sebagai goroutine di handler GetWeather
var GetWetherByRegionAPI = client.GetWeatherByRegion

func GetWeather(w http.ResponseWriter, r *http.Request) {
	var listRegion = []string{"jakarta", "bandung", "surabaya", "yogyakarta", "medan", "makassar", "manado", "palembang", "semarang", "bali"}
	fmt.Println(listRegion)

	// DESC: dapatkan data weather dari 10 data di atas menggunakan goroutine
	// TODO: answer here
}

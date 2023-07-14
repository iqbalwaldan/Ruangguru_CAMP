package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", templateHandler)
	http.HandleFunc("/process", submitHandler)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("view.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("view.html"))

		// Method `r.ParseForm()` berguna untuk parsing data form yang di kirim dari view
		// Juga mengembalikan data `error` jika proses parsing gagal
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Untuk mengembalikan data inputan name (`<input name="name" />`).
		var name = r.FormValue("name")
		// Untuk mengembalikan data inputan role (`<input name="role" />`).
		var role = r.FormValue("role")

		// data ditampung dalam variabel `data` yang bertipe `map[string]string`
		var data = map[string]string{"name": name, "role": role}

		// Variabel `data`di sisipkan ke view, melalui statement `tmpl.Execute(w, data)`
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

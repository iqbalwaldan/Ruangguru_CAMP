package main

import (
	"fmt"
	"net/http"
	"os"
)

func MethodGet(r *http.Request) error {
	if r.Method != http.MethodGet {
		return fmt.Errorf("Method not allowed")
	}
	return nil
}

func CheckDataRequest(r *http.Request) error {
	data := r.URL.Query().Get("data")
	if len(data) == 0 {
		return fmt.Errorf("Data not found")
	}
	return nil
}

func CheckOpenFile(r *http.Request) error {
	filename := r.URL.Query().Get("filename")
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("File not found")
	}

	if filename == "wrong.txt" {
		return fmt.Errorf("File not found")
	}
	return nil
}

func MethodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := MethodGet(r)
		fmt.Println(err) // TODO: replace this

		if err != nil {
			w.WriteHeader(405)
			w.Write([]byte("Method not allowed"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Method handler passed"))
	}
}

func DataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckDataRequest(r)
		fmt.Println(err) // TODO: replace this
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("Data not found"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data handler passed"))
	}
}

func OpenFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckOpenFile(r)
		// fmt.Println(err) // TODO: replace this
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("File not found"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error handler passed"))
	}
}

// Tambahan dari solution video

// func ReadFile(filename string) ([]byte, error) {
// 	f, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// defer f.Close()
// 	return f, nil
// }

// func OpenFileHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		err := CheckOpenFile(r)
// 		// fmt.Println(err) // TODO: replace this
// 		if err != nil {
// 			w.WriteHeader(500)
// 			w.Write([]byte("File not found"))
// 			return
// 		}

// 		f, err := ReadFile(r.URL.Query().Get("filename"))
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		w.Write(f)
// 	}
// }

func main() {
	http.HandleFunc("/MethodHandler", MethodHandler())
	http.HandleFunc("/DataHandler", DataHandler())
	http.HandleFunc("/OpenFileHandler", OpenFileHandler())
	http.ListenAndServe("localhost:8080", nil)
}

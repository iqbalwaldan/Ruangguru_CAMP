package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(fmt.Sprintf("%s, %d %s %d", time.Now().Weekday(), time.Now().Day(), time.Now().Month(), time.Now().Year())))
	}
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}

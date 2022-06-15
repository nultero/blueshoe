package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func echoBack() http.HandlerFunc {
	return func(wr http.ResponseWriter, rq *http.Request) {
		numStr := rq.URL.Path[1:]
		num, _ := strconv.Atoi(numStr)
		if num == 6 {

		}
		fmt.Println(num * 2)
	}
}

func main() {
	port := ":8000"
	http.HandleFunc("/", echoBack())
	http.ListenAndServe(port, nil)
}

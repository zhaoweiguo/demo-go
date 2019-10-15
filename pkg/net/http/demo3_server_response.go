package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	// http.Header
	header := w.Header()
	fmt.Printf("header: %v\n", header)

	code, err := w.Write([]byte("hello world\n"))
	fmt.Printf("code:%d, %v\n", code, err)
	//	w.WriteHeader(404)

}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":7777", nil)
}

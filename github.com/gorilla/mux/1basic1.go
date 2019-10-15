package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/abc", Handler)

	http.Handle("/", r)
	addr := ":7800"
	server := &http.Server{
		Addr:           addr,
		Handler:        r,
	}
	server.ListenAndServe()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start...")
}

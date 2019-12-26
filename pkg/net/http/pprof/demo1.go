package main

import (
	"fmt"
	"html"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	simple_handler()

}

func simple_handler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(":8080", nil)
}

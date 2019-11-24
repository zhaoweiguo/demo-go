package main

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

func main() {

	simple_handler()
	mux_handler()

}

func simple_handler() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(":8080", nil)
}

func mux_handler() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})

	// 更多控制
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//func ListenAndServeTLS(add string, certFile string, keyFile string, handle Handle) error
	// 实例和http请求一样
	//不过把ListenAndServe换为ListenAndServeTLS,如:
	s.ListenAndServeTLS("cert.pem", "key.pem")
}

// ============ inner ============

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 3"))
}

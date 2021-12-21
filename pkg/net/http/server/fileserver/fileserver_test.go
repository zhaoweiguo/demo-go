package main

import (
	"log"
	"net/http"
	"testing"
)

func TestDemo1(t *testing.T) {
	fileServer := http.FileServer(http.Dir("/tmp/abc/123/testChapter"))
	log.Println("http://127.0.0.1:8081")
	log.Fatal(http.ListenAndServe(":8081", fileServer))
}

func TestDemo2(t *testing.T) {
	fileServer := http.FileServer(http.Dir(".."))
	//http.Handle("/fileserver2/", fileServer)
	log.Println("http://127.0.0.1:8081")
	log.Fatal(http.ListenAndServe(":8081", fileServer))
}

func demo3() {
	fileServer := http.FileServer(http.Dir("/tmp"))
	http.Handle("/aaa/", http.StripPrefix("/aaa/", fileServer))
	//http.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

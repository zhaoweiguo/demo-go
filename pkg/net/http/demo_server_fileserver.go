package main

import (
	"log"
	"net/http"
)

func main() {
	demo3()
}

func demo1() {
	fileServer := http.FileServer(http.Dir("/tmp"))
	log.Fatal(http.ListenAndServe(":8081", fileServer))
}

func demo2() {
	fileServer := http.FileServer(http.Dir("/tmp"))
	http.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8081", fileServer))
}

func demo3() {
	fileServer := http.FileServer(http.Dir("/tmp"))
	http.Handle("/aaa/", http.StripPrefix("/aaa/", fileServer))
	//http.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	http.HandleFunc("/", handleFunc1)
	fmt.Println("listen port 7777")
	http.ListenAndServe(":7777", nil)
}
func handleFunc1(w http.ResponseWriter, r *http.Request) {
	log.Println(r.ContentLength)
	str, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(str))
	log.Println(r.Header)

	w.Write([]byte("ok"))
}

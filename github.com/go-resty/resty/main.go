package main

import (
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

func main() {
	url := "https://127.0.0.1:8080/index.html"
	timeout := time.Duration(10) * time.Second
	response, err := resty.New().SetTimeout(timeout).R().Get(url)

	log.Println("error response := ", response, "err := ", err)

}

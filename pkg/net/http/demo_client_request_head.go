package main

import (
	"log"
	"net/http"
)

func main() {

	//func (c *Client) Head(url string) (r *Response, err error)
	resp, error := http.Head("https://www.baidu.com")
	log.Println(resp, error)

}

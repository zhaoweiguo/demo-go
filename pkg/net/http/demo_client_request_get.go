package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	uri := "http://www.baidu.com"
	get_simple(uri)
	get_reader(uri)
	get_do(uri)
}

//func (c *Client) Get(url string) (r *Response, err error)
func get_simple(uri string) {
	resp, error := http.Get(uri)
	log.Println(resp, error)
}

func get_reader(uri string) {
	req, _ := http.NewRequest("GET", uri, strings.NewReader("z=post&both=y"))
	fmt.Println(req)

	body := io.Reader(strings.NewReader("z=post&both=y"))
	fmt.Printf("body:%T | %v\n", body, body)

	fmt.Println("==abc==")
	rc, ok := body.(io.ReadCloser)
	fmt.Println(rc, ok)
}

func get_do(uri string) {

	//func (c *Client) Do(req *Request) (resp *Response, err error)
	// 个性化需求, 如: 自定义"user-Agent", 如传递cookie
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("User-Agent", "Go Customer User-Agent")
	req.Header.Add("If-None-Match", "TheFileEtag")
	client := &http.Client{}
	resp, err := client.Do(req)
	log.Println(resp, err)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"unsafe"
)

func main() {

	url := "http://39.97.31.155/weixin/corp/text"
	msg := "hello world"
	//post_req_json(url, msg)
	post_urlencode(url, msg)
	post_simple(url)
}

func post_simple(url string) {
	//func (c *Client) Post(url string, bodyType string, body io.Reader) (r *Response, err error)
	imageDataBuf := io.Reader(strings.NewReader("z=post&both=y"))
	resp, error := http.Post(url, "image/jpeg", imageDataBuf)
	log.Println(resp, error)
}

func post_req_json(url string, msg string) {
	song := make(map[string]interface{})
	song["text"] = msg

	bytesData, err := json.Marshal(song)
	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)

}

//func (c *Client) PostForm(url string, data url.Values) (r *Response, err error)
// 实现标准编码格式为"application/x-www-form-urlencoded"的表单提交
func post_urlencode(uri, msg string) {
	resp, error := http.PostForm(uri, url.Values{"title": {"article title"}, "content": {msg}})
	log.Println(resp, error)

}

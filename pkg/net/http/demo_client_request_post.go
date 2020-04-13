package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"unsafe"
)

func main() {
	post_json()

	post_urlencode()
	post_simple()
}

func post_simple() {
	token := "c625ed2d7c870d766748a8c38afc0d77e2d8d17c023864c469e296695244a438"
	uri := "https://oapi.dingtalk.com/robot/send?access_token=" + token
	contentType := "application/json;charset=utf-8"
	msg := map[string]interface{}{
		"msgtype": "text",
		"text":    map[string]string{"content": "hello world!"},
	}
	body, _ := json.Marshal(msg)
	resp, error := http.Post(uri, contentType, bytes.NewReader(body))
	log.Println(resp, error)
}

func post_json() {
	uri := "https://oapi.dingtalk.com/robot/send?access_token=c625ed2d7c870d766748a8c38afc0d77e2d8d17c023864c469e296695244a438"
	body := []byte("{\"msgtype\": \"text\", \"text\": {\"content\": \"我就是我, 是不一样的烟火\"} }")
	reader := bytes.NewReader(body)
	request, err := http.NewRequest("POST", uri, reader)
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
func post_urlencode() {
	uri := "http://39.97.31.155/weixin/corp/text"
	msg := "hello world"
	resp, error := http.PostForm(uri, url.Values{"title": {"article title"}, "content": {msg}})
	log.Println(resp, error)

}

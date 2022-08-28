package test

import (
	"log"
	"net/http"
	"net/url"
	"testing"
)

//func (c *Client) PostForm(url string, data url.Values) (r *Response, err error)
// 实现标准编码格式为"application/x-www-form-urlencoded"的表单提交
func TestPost_urlencode(t *testing.T) {
	uri := "http://39.97.31.155/weixin/corp/text"
	msg := "hello world"
	resp, err := http.PostForm(uri, url.Values{"title": {"article title"}, "content": {msg}})
	log.Println(resp, err)

}


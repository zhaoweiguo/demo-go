package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestPost(t *testing.T) {
	token := "c625ed2d7c870d766748a8c38afc0d77e2d8d17c023864c469e296695244a438"
	uri := "https://oapi.dingtalk.com/robot/send?access_token=" + token
	contentType := "application/json;charset=utf-8"
	msg := map[string]interface{}{
		"msgtype": "text",
		"text":    map[string]string{"content": "hello world!"},
	}
	body, _ := json.Marshal(msg)
	resp, err := http.Post(uri, contentType, bytes.NewReader(body))

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err: ", err)
	}
	log.Println(string(respBody), err)
}


func TestPost2(t *testing.T) {
	uri := "https://163.com"
	contentType := "application/x-www-form-urlencoded"
	resp, error := http.Post(uri, contentType, nil)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err: ", err)
	}
	log.Println(string(body), error)
}




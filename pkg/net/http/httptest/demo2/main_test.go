package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
在测试机上模拟一个服务器，然后进行调用测试
用于测试服务端代码route是否正确
*/

func mockServer() *httptest.Server {
	//适配器转换
	sendJson := route()
	return httptest.NewServer(http.HandlerFunc(sendJson))
}

func TestSendJSON(t *testing.T) {
	//创建一个模拟的服务器
	server := mockServer()
	defer server.Close()
	//Get请求发往模拟服务器的地址
	resq, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("创建Get失败")
	}
	defer resq.Body.Close()

	log.Println("code:", resq.StatusCode)
	json, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("body:%s\n", json)
}

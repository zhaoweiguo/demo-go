package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zhaoweiguo/demo-go/pkg/net/http/httptest/demo1/common"
)

/*
模拟http请求
主要利用httptest.NewRecorder()创建一个http.ResponseWriter，模拟了真实服务端的响应，
这种响应时通过调用http.DefaultServeMux.ServeHTTP方法触发的
*/

func init() {
	common.Routes()
}

func TestSendJSON(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/sendjson", nil)
	if err != nil {
		t.Fatal("创建Request失败")
	}

	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	log.Println("code:", rw.Code)

	log.Println("body:", rw.Body.String())
}

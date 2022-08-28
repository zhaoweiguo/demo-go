package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"unsafe"
)

func TestJson(t *testing.T) {
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



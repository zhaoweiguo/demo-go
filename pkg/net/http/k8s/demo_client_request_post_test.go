package k8s

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func Test_post_simple(t *testing.T) {
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

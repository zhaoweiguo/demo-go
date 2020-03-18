package main

/*
参考: https://www.jianshu.com/p/b49cc19d26f0
对应server1: net/http/demo_server_multipart
对应server2: github.com/gin-gonic/demo_multipart
*/
import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	uri := "http://127.0.0.1:8844/multipart"

	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	// file1
	fileWriter1, _ := bodyWriter.CreateFormFile("filename", "file.txt")
	file, _ := os.Open("./tmp/file.txt")
	defer file.Close()
	io.Copy(fileWriter1, file)

	// file2
	fileWriter2, _ := bodyWriter.CreateFormFile("filename2", "file2.txt")
	file2, _ := os.Open("./tmp/file2.txt")
	defer file2.Close()
	io.Copy(fileWriter2, file2)

	// other form data
	extraParams := map[string]string{
		"title":       "标题",
		"author":      "zhaoweiguo",
		"description": "http multipart测试项目",
	}
	for key, value := range extraParams {
		_ = bodyWriter.WriteField(key, value)
	}

	contentType := bodyWriter.FormDataContentType()
	log.Println(contentType)
	log.Println(string(bodyBuffer.Bytes()))
	bodyWriter.Close()

	resp, _ := http.Post(uri, contentType, bodyBuffer)
	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)
	log.Println(resp.Status)
	log.Println(string(resp_body))

}

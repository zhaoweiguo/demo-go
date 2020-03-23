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

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	uri := "http://127.0.0.1:8844/multipart"

	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	// file1
	fileWriter1, _ := bodyWriter.CreateFormFile("filename", "file.txt")
	file, err := os.Open("./tmp/file.txt")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()
	io.Copy(fileWriter1, file)

	//// file2
	//fileWriter2, err := bodyWriter.CreateFormFile("filename2", "file2.txt")
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	//file2, err := os.Open("./tmp/file2.txt")
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	//defer file2.Close()
	//io.Copy(fileWriter2, file2)

	// other form data
	err = bodyWriter.WriteField("title", "标题")
	err = bodyWriter.WriteField("to", "admin@zhaoweiguo.com")
	err = bodyWriter.WriteField("content", "http multipart测试项目")
	err = bodyWriter.WriteField("types", "email")

	contentType := bodyWriter.FormDataContentType()
	log.Println(contentType)
	//log.Println(string(bodyBuffer.Bytes()))
	bodyWriter.Close()

	resp, err := http.Post(uri, contentType, bodyBuffer)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(resp.Status)
	log.Println(string(resp_body))

}

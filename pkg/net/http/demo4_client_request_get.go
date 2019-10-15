package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	req, _ := http.NewRequest("GET", "http://www.baidu.com", strings.NewReader("z=post&both=y"))
	fmt.Println(req)

	body := strings.NewReader("z=post&both=y")
	fmt.Printf("body:%T | %v\n", body, body)

	abc(body)
}

func abc(body io.Reader) {
	fmt.Println("==abc==")
	rc, ok := body.(io.ReadCloser)
	fmt.Println(rc, ok)
}

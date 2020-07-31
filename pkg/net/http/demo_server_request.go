package main

import (
	"fmt"
	"net"
	"net/http"
)

func handleFunc1(w http.ResponseWriter, r *http.Request) {
	// string: GET, POST, PUT, etc.
	method := r.Method
	fmt.Printf("method: %s\n", method)
	// *url.URL: the URI being requested (server requests) or URL to access (client requests).
	url := r.URL
	fmt.Printf("url: %v\n", url)
	//string: "HTTP/1.0"
	proto := r.Proto
	fmt.Printf("proto: %s\n", proto)

	// string 请求机器地址格式: [host:port]
	remoteAddr := r.RemoteAddr
	fmt.Printf("remoteAddr: %s\n", remoteAddr)
	clientIP, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		fmt.Errorf("error in splithostport")
		return
	}

	// http.Header:
	// X-Forwarded-Proto这儿只记录最后一条clientIP
	if prior, ok := r.Header["X-Forwarded-For"]; ok {
		fmt.Printf("X-Forwarded-For:%v\n", prior)
	}
	r.Header.Set("X-Forwarded-For", clientIP)
	xff := r.Header.Get("X-Forwarded-For")
	fmt.Printf("new X-Forwarded-For:%v\n", xff)
	xfp := r.Header.Get("X-Forwarded-Proto")
	fmt.Printf("X-Forwarded-Proto:%v\n", xfp)
	fmt.Printf("request TLS:%v\n", r.TLS)

	header := r.Header
	fmt.Printf("header: %v\n", header)

	// io.ReadCloser
	body := r.Body
	fmt.Printf("body: %v\n", body)
	// int64:
	contentLength := r.ContentLength
	fmt.Printf("contentLength: %d\n", contentLength)

	// []string:
	transferEncoding := r.TransferEncoding
	fmt.Printf("transferEncoding:%v\n", transferEncoding)

	// string
	host := r.Host
	fmt.Printf("host:%s\n", host)

	// url.Values
	form := r.Form
	fmt.Printf("form:%v\n", form)

	// *multipart.Form
	multipartForm := r.MultipartForm
	fmt.Printf("multipartForm: %v\n", multipartForm)

	// http.Header
	trailer := r.Trailer
	fmt.Printf("trailer: %v\n", trailer)

	// string 示例: /ab/c?dd=3
	requestUri := r.RequestURI
	fmt.Printf("requestUri: %v", requestUri)

	w.Write([]byte("ok"))

}

func main() {
	http.HandleFunc("/", handleFunc1)
	fmt.Println("listen port 7777")
	http.ListenAndServe(":7777", nil)
}

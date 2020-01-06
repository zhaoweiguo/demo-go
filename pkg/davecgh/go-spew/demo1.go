package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
	// 注意这儿打印是注释的html, 记得浏览器查看源码查看
	fmt.Fprintf(w, "<!--\n"+html.EscapeString(spew.Sdump(w))+"\n-->")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("http://127.0.0.1:8088")
	if err := http.ListenAndServe(":8088", nil); err != nil {
		fmt.Println("err:", err)
	}
}

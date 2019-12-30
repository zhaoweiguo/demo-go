package main

import (
	"expvar"
	"fmt"
	"net/http"
)

/*
$ curl http://127.0.0.1:8080/debug/vars

	"varFloat:": 9.42,
	"varInt": 3,
	"varMap:": {
		"map": 3
	},
	"varString:": "->1->1->1"

*/

var varInt = expvar.NewInt("varInt") // 设定数
var varString = expvar.NewString("varString:")
var varFloat = expvar.NewFloat("varFloat:")
var varMap = expvar.NewMap("varMap:")

func handler(w http.ResponseWriter, r *http.Request) {
	varInt.Add(1) // 每次http请求增加一次
	varString.Set(varString.Value() + "->1")
	varFloat.Add(3.14)
	varMap.Add("map", 1)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

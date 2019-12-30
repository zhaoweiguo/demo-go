package main

import (
	"expvar"
	"fmt"
	"net/http"
	"time"
)

/*
$ curl http://127.0.0.1:8080/debug/vars

	"uptime": "13.051340383s"
	"http": {
		"req_failed": 10,
		"req_succ": 20
	},
*/

var stats = expvar.NewMap("http")
var requests, requestsFailed expvar.Int
var start = time.Now()

func init() {
	stats.Set("req_succ", &requests)         // 请求成功的次数
	stats.Set("req_failed", &requestsFailed) // 请求失败的次数
	expvar.Publish("uptime", expvar.Func(calculateUptime))
}

func calculateUptime() interface{} {
	return time.Since(start).String()
}

func handler2(w http.ResponseWriter, r *http.Request) {
	requests.Add(2)       // 模拟成功请求次数增加
	requestsFailed.Add(1) // 模拟失败请求次数增加
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler2)
	http.ListenAndServe(":8080", nil)
}
